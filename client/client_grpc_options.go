package client

import (
	"github.com/anthonytsu1984/gosdktobe/entity"
	server "github.com/milvus-io/milvus-proto/go-api/milvuspb"
)

// This function type defines a set of functions that take a CreateCollectionRequest struct type as the only parameter.
type CreateCollectionOption func(*server.CreateCollectionRequest)

// This function sets the consistency level in a CreateCollectionRequest struct type.
//
// # Parameters
//
//   - cl
//
//     Specifies a consistency level in a CreateCollectionRequest struct type. You can find possible options here.
//
// # Return
//
// A function that takes the modified CreateCollectionRequest struct type as the only parameter.
func WithConsistencyLevel(cl entity.ConsistencyLevel) CreateCollectionOption {
	return func(req *server.CreateCollectionRequest) {
		req.ConsistencyLevel = cl.CommonConsistencyLevel()
	}
}
