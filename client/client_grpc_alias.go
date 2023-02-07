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

import "context"

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
// # Returns
//
// In normal cases, this method returns nil, indicating that this operation succeeds. For other values, see "Possible Errors".
func (c *GrpcClient) AlterAlias(ctx context.Context, collectionName string, alias string) error {
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
// # Returns
//
// In normal cases, this method returns nil, indicating that this operation succeeds. For other values, see "Possible Errors".
func (c *GrpcClient) CreateAlias(ctx context.Context, collectionName string, alias string) error {
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
// # Returns
//
// In normal cases, this method returns nil, indicating that this operation succeeds. For other values, see "Possible Errors".
func (c *GrpcClient) DropAlias(ctx context.Context, alias string) error {
	return nil
}

// This method lists aliases of a collection.
//
// # Parameters
//
//   - collectionName
//
//     Specifies the name of a collection.
//
//     A collection name should be a string of 1 to 255 characters, starting with a letter or an underscore (_) and containing only numbers, letters, and underscores (_).
//
// # Returns
//
// In normal cases, this method returns nil, indicating that this operation succeeds. For other values, see "Possible Errors".
func (c *GrpcClient) ListAliases(ctx context.Context, collectionName string) error {
	return nil
}
