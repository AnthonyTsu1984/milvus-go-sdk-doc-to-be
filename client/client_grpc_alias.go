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

package client

import (
	"context"

	"github.com/anthonytsu1984/gosdktobe/entity"
)

// This method changes an alias of a collection.
//
// # Parameters
//
//   - collectionName
//
//     Specifies the name of a collection.
//
//     A collection name should be a string of 1 to 255 characters, starting with a letter or an underscore (_) and containing only numbers, letters, and underscores (_).
//
//   - alias
//
//     Specifies a new alias for the collection.
//
//     A collection alias is a string of 1 to 255 characters, starting with a letter or an underscore (_) and containing only numbers, letters, and underscores (_).
//
//   - timeout
//
//     Specifies the timeout duration of this operation in seconds.
//
//     This parameter is optional. If left unspecified, no such limit applies.
//
// # Returns
//
// In normal cases, this method returns nil, indicating that this operation succeeds. For other values, see "Possible Errors".
//
// # Possible Errors
func (c *GrpcClient) AlterAlias(ctx context.Context, collectionName string, alias string, timeout float32) error {
	return nil
}

// This method creates an alias for a collection.
//
// # Parameters
//
//   - collectionName
//
//     Specifies the name of a collection.
//
//     A collection name should be a string of 1 to 255 characters, starting with a letter or an underscore (_) and containing only numbers, letters, and underscores (_).
//
//   - alias
//
//     Specifies a new alias for the collection.
//
//     A collection alias is a string of 1 to 255 characters, starting with a letter or an underscore (_) and containing only numbers, letters, and underscores (_).
//
//   - timeout
//
//     Specifies the timeout duration of this operation in seconds.
//
//     This parameter is optional. If left unspecified, no such limit applies.
//
// # Returns
//
// In normal cases, this method returns nil, indicating that this operation succeeds. For other values, see "Possible Errors".
//
// # Possible Errors
func (c *GrpcClient) CreateAlias(ctx context.Context, collectionName string, alias string, timeout float32) error {
	return nil
}

// This method drops an alias from a collection.
//
// # Parameters
//
//   - alias
//
//     Specifies a new alias for the collection.
//
//     A collection alias is a string of 1 to 255 characters, starting with a letter or an underscore (_) and containing only numbers, letters, and underscores (_).
//
//   - timeout
//
//     Specifies the timeout duration of this operation in seconds.
//
//     This parameter is optional. If left unspecified, no such limit applies.
//
// # Returns
//
// In normal cases, this method returns nil, indicating that this operation succeeds. For other values, see "Possible Errors".
//
// # Possible Errors
func (c *GrpcClient) DropAlias(ctx context.Context, alias string, timeout float32) error {
	return nil
}

// This method lists aliases of a collection.
//
// # Parameters
//
//   - alias
//
//     Specifies an alias to drop.
//
//     A collection alias should be a string of 1 to 255 characters, starting with a letter or an underscore (_) and containing only numbers, letters, and underscores (_).
//
//   - timeout
//
//     Specifies the timeout duration of this operation in seconds.
//
//     This parameter is optional. If left unspecified, no such limit applies.
//
// # Returns
//
// This method returns a list of aliases and an error,
//
//   - []string
//
//     Lists the aliases of the specified collection.
//
//   - error
//
//     This is nil in normal cases. For other values, see "Possible Errors".
//
// # Possible Errors
func (c *GrpcClient) ListAliases(ctx context.Context, collectionName string, timeout float32) ([]string, error) {
	a := []string{"a", "b"}
	return a, nil
}

// This method describes an alias in detail.
//
// # Parameters
//
//   - alias
//
//     Specifies an alias to describe.
//
//     A collection alias is a string of 1 to 255 characters, starting with a letter or an underscore (_) and containing only numbers, letters, and underscores (_).
//
// # Returns
//
// In normal cases, this method returns an [entity.AliasInfo] object and an error,
//
//   - [entity.AliasInfo]
//
//     Lists the details of an alias.
//
//   - nil
//
//     This is nil in normal cases. For other values, see "Possible Errors".
//
//   - timeout
//
//     Specifies the timeout duration of this operation in seconds.
//
//     This parameter is optional. If left unspecified, no such limit applies.
//
// # Possible Errors
func (c *GrpcClient) DescribeAlias(ctx context.Context, alias string, timeout float32) (entity.AliasInfo, error) {
	a := entity.AliasInfo{ID: 1, CollectionName: "collection", Alias: "alias"}
	return a, nil
}

// This method shows whether the specified alias exists.
//
// # Parameters
//
//   - alias
//
//     Specifies an alias to check.
//
//     A collection alias is a string of 1 to 255 characters, starting with a letter or an underscore (_) and containing only numbers, letters, and underscores (_).
//
//   - timeout
//
//     Specifies the timeout duration of this operation in seconds.
//
//     This parameter is optional. If left unspecified, no such limit applies.
//
// # Returns
//
// In normal cases, this method returns an AliasInfo object and an error,
//
//   - bool
//
//     Indicates whether the specified alias exists.
//
//   - nil
//
//     This is nil in normal cases. For other values, see "Possible Errors".
//
// # Possible Errors
func (c *GrpcClient) HasAlias(ctx context.Context, alias string, timeout float32) (bool, error) {
	a := true
	return a, nil
}
