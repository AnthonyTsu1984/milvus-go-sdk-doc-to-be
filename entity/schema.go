// Copyright (C) 2019-2023 Zilliz. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License
// is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
// or implied. See the License for the specific language governing permissions and limitations under the License.

// This package offers necessary type definitions relevant to the Milvus client.
package entity

import (
	common "github.com/milvus-io/milvus-proto/go-api/commonpb"
)

// This type references the ConsistencyLevel type defined in a proprietary repository.
type ConsistencyLevel common.ConsistencyLevel

// This method sets the consistency level in a ConsistencyLevel type.
func (cl ConsistencyLevel) CommonConsistencyLevel() common.ConsistencyLevel {
	return common.ConsistencyLevel(cl)
}

// This type defines the structure of a collection schema.
//
// # Fields
//
//   - CollectionName
//
//     A collection name should be a string of 1 to 255 characters, starting with a letter or an underscore (_) and containing only numbers, letters, and underscores (_).
//
//   - Description
//
//     A description should be a string of 1 to 65535 characters, starting with a letter or an underscore (_) and containing only numbers, letters, and underscores (_).
//
//   - AutoID ??
//
//     Whether the primary key is automatically generated. If set to true, you do not need to set the primary key in fields.
//
//   - Fields
//
//     A list of [Field] struct values.
type Schema struct {
	CollectionName string
	Description    string
	AutoID         bool
	Fields         []*Field
}

// This type defines the structure of a schema field.
//
// # Fields
//
//   - ID
//
//     The unique ID of a field. ??
//
//   - Name
//
//     A field name should be a string of 1 to 255 characters, starting with a letter or an underscore (_) and containing only numbers, letters, and underscores (_).
//
//   - PriamryKey
//
//     A boolean value indicates whether a field is a primary key. There is only one primary field among all the fields in a schema.
//
//   - AutoID
//
//     A boolean value indicates whether the primary key is automatically generated.
//
//   - Description
//
//     A field description should be a string of 1 to 65535 characters, starting with a letter or an underscore (_) and containing only numbers, letters, and underscores (_).
//
//   - DataType
//
//     A [FieldType] struct value.
//
//   - TypeParams
//
//     A list of key-value pairs that define a set of parameters specific to the specified data type.
//
//   - IndexParams
//
//     A list of key-value pairs that define a set of parameters specific to the index-building of a field.
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

// This type defines the possible field types with a set of int32 numbers.
type FieldType int32

// A [FieldType] method that returns a string as the name of a specified data type.
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

// A [FieldType] method that returns a string as the description of a field type.
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

// A [FieldType] method that returns ??
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

// The following constant definitions are values used to represent each [FieldType].
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
