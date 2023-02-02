package entity

import (
	common "github.com/milvus-io/milvus-proto/go-api/commonpb"
)

type ConsistencyLevel common.ConsistencyLevel

func (cl ConsistencyLevel) CommonConsistencyLevel() common.ConsistencyLevel {
	return common.ConsistencyLevel(cl)
}

type Schema struct {
	CollectionName string
	Description    string
	AutoID         bool
	Fields         []*Field
}

type Field struct {
	ID          int64
	Name        string
	PrimaryKey  bool
	AutoID      bool
	Description string
	DataType    FieldType
	TypeParams  map[string]string
	IndexParams map[string]string
}

// FieldType field data type alias type
// used in go:generate trick, DO NOT modify names & string
type FieldType int32

// Name returns field type name
func (t FieldType) Name() string {
	switch t {
	case FieldTypeBool:
		return "Bool"
	case FieldTypeInt8:
		return "Int8"
	case FieldTypeInt16:
		return "Int16"
	case FieldTypeInt32:
		return "Int32"
	case FieldTypeInt64:
		return "Int64"
	case FieldTypeFloat:
		return "Float"
	case FieldTypeDouble:
		return "Double"
	case FieldTypeString:
		return "String"
	case FieldTypeVarChar:
		return "VarChar"
	case FieldTypeBinaryVector:
		return "BinaryVector"
	case FieldTypeFloatVector:
		return "FloatVector"
	default:
		return "undefined"
	}
}

// String returns field type
func (t FieldType) String() string {
	switch t {
	case FieldTypeBool:
		return "bool"
	case FieldTypeInt8:
		return "int8"
	case FieldTypeInt16:
		return "int16"
	case FieldTypeInt32:
		return "int32"
	case FieldTypeInt64:
		return "int64"
	case FieldTypeFloat:
		return "float32"
	case FieldTypeDouble:
		return "float64"
	case FieldTypeString:
		return "string"
	case FieldTypeVarChar:
		return "string"
	case FieldTypeBinaryVector:
		return "[]byte"
	case FieldTypeFloatVector:
		return "[]float32"
	default:
		return "undefined"
	}
}

// PbFieldType represents FieldType corresponding schema pb type
func (t FieldType) PbFieldType() (string, string) {
	switch t {
	case FieldTypeBool:
		return "Bool", "bool"
	case FieldTypeInt8:
		fallthrough
	case FieldTypeInt16:
		fallthrough
	case FieldTypeInt32:
		return "Int", "int32"
	case FieldTypeInt64:
		return "Long", "int64"
	case FieldTypeFloat:
		return "Float", "float32"
	case FieldTypeDouble:
		return "Double", "float64"
	case FieldTypeString:
		return "String", "string"
	case FieldTypeVarChar:
		return "VarChar", "string"
	case FieldTypeBinaryVector:
		return "[]byte", ""
	case FieldTypeFloatVector:
		return "[]float32", ""
	default:
		return "undefined", ""

	}

}

// Match schema definition
const (
	// FieldTypeNone zero value place holder
	FieldTypeNone FieldType = 0 // zero value place holder
	// FieldTypeBool field type boolean
	FieldTypeBool FieldType = 1
	// FieldTypeInt8 field type int8
	FieldTypeInt8 FieldType = 2
	// FieldTypeInt16 field type int16
	FieldTypeInt16 FieldType = 3
	// FieldTypeInt32 field type int32
	FieldTypeInt32 FieldType = 4
	// FieldTypeInt64 field type int64
	FieldTypeInt64 FieldType = 5
	// FieldTypeFloat field type float
	FieldTypeFloat FieldType = 10
	// FieldTypeDouble field type double
	FieldTypeDouble FieldType = 11
	// FieldTypeString field type string
	FieldTypeString FieldType = 20
	// FieldTypeVarChar field type varchar
	FieldTypeVarChar FieldType = 21 // variable-length strings with a specified maximum length
	// FieldTypeBinaryVector field type binary vector
	FieldTypeBinaryVector FieldType = 100
	// FieldTypeFloatVector field type float vector
	FieldTypeFloatVector FieldType = 101
)
