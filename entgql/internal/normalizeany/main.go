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

// Command normalizeany rewrites the interface{} spelling to any in the given
// generated files. gqlgen renders unbound scalar Go types through go/types,
// whose alias printing is not deterministic across runs — without this
// normalization the generated-files CI gate flip-flops between the two
// spellings.
package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	for _, path := range os.Args[1:] {
		b, err := os.ReadFile(path)
		if err != nil {
			log.Fatalf("normalizeany: %v", err)
		}
		s := strings.ReplaceAll(string(b), "interface{}", "any")
		if s == string(b) {
			continue
		}
		if err := os.WriteFile(path, []byte(s), 0644); err != nil {
			log.Fatalf("normalizeany: %v", err)
		}
	}
}
