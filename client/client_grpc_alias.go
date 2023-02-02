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

// Changes an alias of a collection.
//
// # Parameters
//   - colName Specifies the name of a collection.
//   - alias Specifies a new alias for the collection.
func (c *GrpcClient) AlterAlias(ctx context.Context, colName string, alias string) error

func (c *GrpcClient) CreateAlias(ctx context.Context, colName string, alias string) error

func (c *GrpcClient) DropAlias(ctx context.Context, alias string) error

func (c *GrpcClient) ListAlias(ctx context.Context, alias string) error
