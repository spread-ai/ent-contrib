// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package multischema_test

import (
	"context"
	"testing"
	"time"

	gen "entgo.io/contrib/entgql/internal/multischema"
	"entgo.io/contrib/entgql/internal/multischema/ent"
	"entgo.io/contrib/entgql/internal/multischema/ent/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
)

const mysqlDSN = "root:pass@tcp(localhost:3308)/?parseTime=true&multiStatements=true"

func TestMySQL(t *testing.T) {
	ctx := context.Background()

	root, err := sql.Open(dialect.MySQL, mysqlDSN)
	if err != nil {
		t.Skipf("mysql not available: %v", err)
	}
	pingCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	if err := root.DB().PingContext(pingCtx); err != nil {
		root.Close()
		t.Skipf("mysql not reachable on :3308: %v", err)
	}
	// Drop the databases (foreign_key_checks off so the cross-schema join-table
	// FK doesn't block it) and close the root connection. Registered before the
	// schemas are created so it also runs if a later step fails. Note: this must
	// not be a plain `defer root.Close()` — deferred calls run before t.Cleanup,
	// which would close the connection out from under the cleanup.
	t.Cleanup(func() {
		for _, stmt := range []string{
			"SET foreign_key_checks = 0",
			"DROP DATABASE IF EXISTS db1",
			"DROP DATABASE IF EXISTS db2",
			"DROP DATABASE IF EXISTS public",
			"SET foreign_key_checks = 1",
		} {
			// Best-effort cleanup; ignore errors.
			_, err = root.DB().ExecContext(ctx, stmt)
			require.NoError(t, err)
		}
		root.Close()
	})

	// Place the entities in db1 and the M2M join table in db2, so the loader
	// has to qualify the join table with its configured schema.
	migrate.GroupsTable.Schema = "db1"
	migrate.UsersTable.Schema = "db1"
	migrate.GroupUsersTable.Schema = "db2"

	// Create the schemas and tables (schema.Dump emits the CREATE DATABASE
	// statements for db1/db2), then a neutral "public" schema for the client
	// connection's default search path.
	plan, err := schema.Dump(ctx, dialect.MySQL, "8.0.19", migrate.Tables)
	require.NoError(t, err)
	_, err = root.DB().ExecContext(ctx, plan)
	require.NoError(t, err)
	_, err = root.DB().ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS `public`")
	require.NoError(t, err)

	// Connect with "public" as the default schema and map every table to its
	// real schema so all references are explicitly qualified.
	conn, err := sql.Open(dialect.MySQL, "root:pass@tcp(localhost:3308)/public?parseTime=true")
	require.NoError(t, err)
	defer conn.Close()
	ec := ent.NewClient(ent.Driver(conn), ent.AlternateSchema(ent.SchemaConfig{
		Group:      "db1",
		User:       "db1",
		GroupUsers: "db2",
	}))

	// Seed two groups with different membership counts.
	users := make([]*ent.User, 3)
	for i := range users {
		users[i] = ec.User.Create().SetName("user").SaveX(ctx)
	}
	ec.Group.Create().SetName("g1").AddUsers(users[0], users[1], users[2]).SaveX(ctx)
	ec.Group.Create().SetName("g2").AddUsers(users[0]).SaveX(ctx)

	// Query the groups connection, selecting only the totalCount of each
	// group's users connection. Selecting totalCount without edges drives the
	// multi-node M2M loader that builds the schema-qualified join.
	srv := handler.NewDefaultServer(gen.NewSchema(ec))
	gqlClient := client.New(srv)

	var resp struct {
		Groups struct {
			Edges []struct {
				Node struct {
					Name  string
					Users struct {
						TotalCount int
					}
				}
			}
		}
	}
	gqlClient.MustPost(`{
		groups {
			edges {
				node {
					name
					users {
						totalCount
					}
				}
			}
		}
	}`, &resp)

	got := make(map[string]int, len(resp.Groups.Edges))
	for _, e := range resp.Groups.Edges {
		got[e.Node.Name] = e.Node.Users.TotalCount
	}
	require.Equal(t, map[string]int{"g1": 3, "g2": 1}, got)
}
