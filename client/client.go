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

// This package offers a Milvus client implemented in Go.
//
// To install this package, run the following:
//
//	go get -u github.com/anthonytsu1984/gosdktobe
package client

import (
	"google.golang.org/grpc"

	server "github.com/milvus-io/milvus-proto/go-api/milvuspb"
)

// This struct type offers a gRPC client to communicate with Milvus.
type GrpcClient struct {
	Conn    *grpc.ClientConn
	Service server.MilvusServiceClient
}
