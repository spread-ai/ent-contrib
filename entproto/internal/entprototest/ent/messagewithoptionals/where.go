// Code generated by ent, DO NOT EDIT.

package messagewithoptionals

import (
	"time"

	"entgo.io/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldLTE(FieldID, id))
}

// StrOptional applies equality check predicate on the "str_optional" field. It's identical to StrOptionalEQ.
func StrOptional(v string) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldEQ(FieldStrOptional, v))
}

// IntOptional applies equality check predicate on the "int_optional" field. It's identical to IntOptionalEQ.
func IntOptional(v int8) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldEQ(FieldIntOptional, v))
}

// UintOptional applies equality check predicate on the "uint_optional" field. It's identical to UintOptionalEQ.
func UintOptional(v uint8) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldEQ(FieldUintOptional, v))
}

// FloatOptional applies equality check predicate on the "float_optional" field. It's identical to FloatOptionalEQ.
func FloatOptional(v float32) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldEQ(FieldFloatOptional, v))
}

// BoolOptional applies equality check predicate on the "bool_optional" field. It's identical to BoolOptionalEQ.
func BoolOptional(v bool) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldEQ(FieldBoolOptional, v))
}

// BytesOptional applies equality check predicate on the "bytes_optional" field. It's identical to BytesOptionalEQ.
func BytesOptional(v []byte) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldEQ(FieldBytesOptional, v))
}

// UUIDOptional applies equality check predicate on the "uuid_optional" field. It's identical to UUIDOptionalEQ.
func UUIDOptional(v uuid.UUID) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldEQ(FieldUUIDOptional, v))
}

// TimeOptional applies equality check predicate on the "time_optional" field. It's identical to TimeOptionalEQ.
func TimeOptional(v time.Time) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldEQ(FieldTimeOptional, v))
}

// StrOptionalEQ applies the EQ predicate on the "str_optional" field.
func StrOptionalEQ(v string) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldEQ(FieldStrOptional, v))
}

// StrOptionalNEQ applies the NEQ predicate on the "str_optional" field.
func StrOptionalNEQ(v string) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldNEQ(FieldStrOptional, v))
}

// StrOptionalIn applies the In predicate on the "str_optional" field.
func StrOptionalIn(vs ...string) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldIn(FieldStrOptional, vs...))
}

// StrOptionalNotIn applies the NotIn predicate on the "str_optional" field.
func StrOptionalNotIn(vs ...string) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldNotIn(FieldStrOptional, vs...))
}

// StrOptionalGT applies the GT predicate on the "str_optional" field.
func StrOptionalGT(v string) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldGT(FieldStrOptional, v))
}

// StrOptionalGTE applies the GTE predicate on the "str_optional" field.
func StrOptionalGTE(v string) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldGTE(FieldStrOptional, v))
}

// StrOptionalLT applies the LT predicate on the "str_optional" field.
func StrOptionalLT(v string) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldLT(FieldStrOptional, v))
}

// StrOptionalLTE applies the LTE predicate on the "str_optional" field.
func StrOptionalLTE(v string) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldLTE(FieldStrOptional, v))
}

// StrOptionalContains applies the Contains predicate on the "str_optional" field.
func StrOptionalContains(v string) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldContains(FieldStrOptional, v))
}

// StrOptionalHasPrefix applies the HasPrefix predicate on the "str_optional" field.
func StrOptionalHasPrefix(v string) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldHasPrefix(FieldStrOptional, v))
}

// StrOptionalHasSuffix applies the HasSuffix predicate on the "str_optional" field.
func StrOptionalHasSuffix(v string) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldHasSuffix(FieldStrOptional, v))
}

// StrOptionalIsNil applies the IsNil predicate on the "str_optional" field.
func StrOptionalIsNil() predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldIsNull(FieldStrOptional))
}

// StrOptionalNotNil applies the NotNil predicate on the "str_optional" field.
func StrOptionalNotNil() predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldNotNull(FieldStrOptional))
}

// StrOptionalEqualFold applies the EqualFold predicate on the "str_optional" field.
func StrOptionalEqualFold(v string) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldEqualFold(FieldStrOptional, v))
}

// StrOptionalContainsFold applies the ContainsFold predicate on the "str_optional" field.
func StrOptionalContainsFold(v string) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldContainsFold(FieldStrOptional, v))
}

// IntOptionalEQ applies the EQ predicate on the "int_optional" field.
func IntOptionalEQ(v int8) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldEQ(FieldIntOptional, v))
}

// IntOptionalNEQ applies the NEQ predicate on the "int_optional" field.
func IntOptionalNEQ(v int8) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldNEQ(FieldIntOptional, v))
}

// IntOptionalIn applies the In predicate on the "int_optional" field.
func IntOptionalIn(vs ...int8) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldIn(FieldIntOptional, vs...))
}

// IntOptionalNotIn applies the NotIn predicate on the "int_optional" field.
func IntOptionalNotIn(vs ...int8) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldNotIn(FieldIntOptional, vs...))
}

// IntOptionalGT applies the GT predicate on the "int_optional" field.
func IntOptionalGT(v int8) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldGT(FieldIntOptional, v))
}

// IntOptionalGTE applies the GTE predicate on the "int_optional" field.
func IntOptionalGTE(v int8) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldGTE(FieldIntOptional, v))
}

// IntOptionalLT applies the LT predicate on the "int_optional" field.
func IntOptionalLT(v int8) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldLT(FieldIntOptional, v))
}

// IntOptionalLTE applies the LTE predicate on the "int_optional" field.
func IntOptionalLTE(v int8) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldLTE(FieldIntOptional, v))
}

// IntOptionalIsNil applies the IsNil predicate on the "int_optional" field.
func IntOptionalIsNil() predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldIsNull(FieldIntOptional))
}

// IntOptionalNotNil applies the NotNil predicate on the "int_optional" field.
func IntOptionalNotNil() predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldNotNull(FieldIntOptional))
}

// UintOptionalEQ applies the EQ predicate on the "uint_optional" field.
func UintOptionalEQ(v uint8) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldEQ(FieldUintOptional, v))
}

// UintOptionalNEQ applies the NEQ predicate on the "uint_optional" field.
func UintOptionalNEQ(v uint8) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldNEQ(FieldUintOptional, v))
}

// UintOptionalIn applies the In predicate on the "uint_optional" field.
func UintOptionalIn(vs ...uint8) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldIn(FieldUintOptional, vs...))
}

// UintOptionalNotIn applies the NotIn predicate on the "uint_optional" field.
func UintOptionalNotIn(vs ...uint8) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldNotIn(FieldUintOptional, vs...))
}

// UintOptionalGT applies the GT predicate on the "uint_optional" field.
func UintOptionalGT(v uint8) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldGT(FieldUintOptional, v))
}

// UintOptionalGTE applies the GTE predicate on the "uint_optional" field.
func UintOptionalGTE(v uint8) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldGTE(FieldUintOptional, v))
}

// UintOptionalLT applies the LT predicate on the "uint_optional" field.
func UintOptionalLT(v uint8) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldLT(FieldUintOptional, v))
}

// UintOptionalLTE applies the LTE predicate on the "uint_optional" field.
func UintOptionalLTE(v uint8) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldLTE(FieldUintOptional, v))
}

// UintOptionalIsNil applies the IsNil predicate on the "uint_optional" field.
func UintOptionalIsNil() predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldIsNull(FieldUintOptional))
}

// UintOptionalNotNil applies the NotNil predicate on the "uint_optional" field.
func UintOptionalNotNil() predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldNotNull(FieldUintOptional))
}

// FloatOptionalEQ applies the EQ predicate on the "float_optional" field.
func FloatOptionalEQ(v float32) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldEQ(FieldFloatOptional, v))
}

// FloatOptionalNEQ applies the NEQ predicate on the "float_optional" field.
func FloatOptionalNEQ(v float32) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldNEQ(FieldFloatOptional, v))
}

// FloatOptionalIn applies the In predicate on the "float_optional" field.
func FloatOptionalIn(vs ...float32) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldIn(FieldFloatOptional, vs...))
}

// FloatOptionalNotIn applies the NotIn predicate on the "float_optional" field.
func FloatOptionalNotIn(vs ...float32) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldNotIn(FieldFloatOptional, vs...))
}

// FloatOptionalGT applies the GT predicate on the "float_optional" field.
func FloatOptionalGT(v float32) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldGT(FieldFloatOptional, v))
}

// FloatOptionalGTE applies the GTE predicate on the "float_optional" field.
func FloatOptionalGTE(v float32) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldGTE(FieldFloatOptional, v))
}

// FloatOptionalLT applies the LT predicate on the "float_optional" field.
func FloatOptionalLT(v float32) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldLT(FieldFloatOptional, v))
}

// FloatOptionalLTE applies the LTE predicate on the "float_optional" field.
func FloatOptionalLTE(v float32) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldLTE(FieldFloatOptional, v))
}

// FloatOptionalIsNil applies the IsNil predicate on the "float_optional" field.
func FloatOptionalIsNil() predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldIsNull(FieldFloatOptional))
}

// FloatOptionalNotNil applies the NotNil predicate on the "float_optional" field.
func FloatOptionalNotNil() predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldNotNull(FieldFloatOptional))
}

// BoolOptionalEQ applies the EQ predicate on the "bool_optional" field.
func BoolOptionalEQ(v bool) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldEQ(FieldBoolOptional, v))
}

// BoolOptionalNEQ applies the NEQ predicate on the "bool_optional" field.
func BoolOptionalNEQ(v bool) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldNEQ(FieldBoolOptional, v))
}

// BoolOptionalIsNil applies the IsNil predicate on the "bool_optional" field.
func BoolOptionalIsNil() predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldIsNull(FieldBoolOptional))
}

// BoolOptionalNotNil applies the NotNil predicate on the "bool_optional" field.
func BoolOptionalNotNil() predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldNotNull(FieldBoolOptional))
}

// BytesOptionalEQ applies the EQ predicate on the "bytes_optional" field.
func BytesOptionalEQ(v []byte) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldEQ(FieldBytesOptional, v))
}

// BytesOptionalNEQ applies the NEQ predicate on the "bytes_optional" field.
func BytesOptionalNEQ(v []byte) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldNEQ(FieldBytesOptional, v))
}

// BytesOptionalIn applies the In predicate on the "bytes_optional" field.
func BytesOptionalIn(vs ...[]byte) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldIn(FieldBytesOptional, vs...))
}

// BytesOptionalNotIn applies the NotIn predicate on the "bytes_optional" field.
func BytesOptionalNotIn(vs ...[]byte) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldNotIn(FieldBytesOptional, vs...))
}

// BytesOptionalGT applies the GT predicate on the "bytes_optional" field.
func BytesOptionalGT(v []byte) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldGT(FieldBytesOptional, v))
}

// BytesOptionalGTE applies the GTE predicate on the "bytes_optional" field.
func BytesOptionalGTE(v []byte) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldGTE(FieldBytesOptional, v))
}

// BytesOptionalLT applies the LT predicate on the "bytes_optional" field.
func BytesOptionalLT(v []byte) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldLT(FieldBytesOptional, v))
}

// BytesOptionalLTE applies the LTE predicate on the "bytes_optional" field.
func BytesOptionalLTE(v []byte) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldLTE(FieldBytesOptional, v))
}

// BytesOptionalIsNil applies the IsNil predicate on the "bytes_optional" field.
func BytesOptionalIsNil() predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldIsNull(FieldBytesOptional))
}

// BytesOptionalNotNil applies the NotNil predicate on the "bytes_optional" field.
func BytesOptionalNotNil() predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldNotNull(FieldBytesOptional))
}

// UUIDOptionalEQ applies the EQ predicate on the "uuid_optional" field.
func UUIDOptionalEQ(v uuid.UUID) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldEQ(FieldUUIDOptional, v))
}

// UUIDOptionalNEQ applies the NEQ predicate on the "uuid_optional" field.
func UUIDOptionalNEQ(v uuid.UUID) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldNEQ(FieldUUIDOptional, v))
}

// UUIDOptionalIn applies the In predicate on the "uuid_optional" field.
func UUIDOptionalIn(vs ...uuid.UUID) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldIn(FieldUUIDOptional, vs...))
}

// UUIDOptionalNotIn applies the NotIn predicate on the "uuid_optional" field.
func UUIDOptionalNotIn(vs ...uuid.UUID) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldNotIn(FieldUUIDOptional, vs...))
}

// UUIDOptionalGT applies the GT predicate on the "uuid_optional" field.
func UUIDOptionalGT(v uuid.UUID) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldGT(FieldUUIDOptional, v))
}

// UUIDOptionalGTE applies the GTE predicate on the "uuid_optional" field.
func UUIDOptionalGTE(v uuid.UUID) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldGTE(FieldUUIDOptional, v))
}

// UUIDOptionalLT applies the LT predicate on the "uuid_optional" field.
func UUIDOptionalLT(v uuid.UUID) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldLT(FieldUUIDOptional, v))
}

// UUIDOptionalLTE applies the LTE predicate on the "uuid_optional" field.
func UUIDOptionalLTE(v uuid.UUID) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldLTE(FieldUUIDOptional, v))
}

// UUIDOptionalIsNil applies the IsNil predicate on the "uuid_optional" field.
func UUIDOptionalIsNil() predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldIsNull(FieldUUIDOptional))
}

// UUIDOptionalNotNil applies the NotNil predicate on the "uuid_optional" field.
func UUIDOptionalNotNil() predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldNotNull(FieldUUIDOptional))
}

// TimeOptionalEQ applies the EQ predicate on the "time_optional" field.
func TimeOptionalEQ(v time.Time) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldEQ(FieldTimeOptional, v))
}

// TimeOptionalNEQ applies the NEQ predicate on the "time_optional" field.
func TimeOptionalNEQ(v time.Time) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldNEQ(FieldTimeOptional, v))
}

// TimeOptionalIn applies the In predicate on the "time_optional" field.
func TimeOptionalIn(vs ...time.Time) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldIn(FieldTimeOptional, vs...))
}

// TimeOptionalNotIn applies the NotIn predicate on the "time_optional" field.
func TimeOptionalNotIn(vs ...time.Time) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldNotIn(FieldTimeOptional, vs...))
}

// TimeOptionalGT applies the GT predicate on the "time_optional" field.
func TimeOptionalGT(v time.Time) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldGT(FieldTimeOptional, v))
}

// TimeOptionalGTE applies the GTE predicate on the "time_optional" field.
func TimeOptionalGTE(v time.Time) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldGTE(FieldTimeOptional, v))
}

// TimeOptionalLT applies the LT predicate on the "time_optional" field.
func TimeOptionalLT(v time.Time) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldLT(FieldTimeOptional, v))
}

// TimeOptionalLTE applies the LTE predicate on the "time_optional" field.
func TimeOptionalLTE(v time.Time) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldLTE(FieldTimeOptional, v))
}

// TimeOptionalIsNil applies the IsNil predicate on the "time_optional" field.
func TimeOptionalIsNil() predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldIsNull(FieldTimeOptional))
}

// TimeOptionalNotNil applies the NotNil predicate on the "time_optional" field.
func TimeOptionalNotNil() predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.FieldNotNull(FieldTimeOptional))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.MessageWithOptionals) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.MessageWithOptionals) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.MessageWithOptionals) predicate.MessageWithOptionals {
	return predicate.MessageWithOptionals(sql.NotPredicates(p))
}
