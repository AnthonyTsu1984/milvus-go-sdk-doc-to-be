package client

import (
	"google.golang.org/grpc"

	server "github.com/milvus-io/milvus-proto/go-api/milvuspb"
)

type GrpcClient struct {
	Conn    *grpc.ClientConn
	Service server.MilvusServiceClient
}
