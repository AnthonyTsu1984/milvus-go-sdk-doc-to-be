package client

import (
	"context"

	"github.com/anthonytsu1984/gosdktobe/entity"
)

// This GrpcClient method creates a collection with a pre-defined schema.
//
// # Paramters
//
//   - schema
//
//     Specifies a schema for the collection.
//
//     To prepare a schema, use [entity.Schema] to construct a Schema struct.
//
//   - numShards
//
//     Specifies the number of shards in the collection. Any incoming DML request is routed to a shard based on the hash value of the affected primary key.
//
//     The number of shards should be no greater than 256. The value defaults to 2, indicating that two shards are to be created.
//
//   - opts
//
//     Specifies optional parameters.
//
//     As to this method, you can set the consistency level of the collection to be created.
//
// # Returns
//
// nil, indicating that this operation succeeds.
func (*GrpcClient) CreateCollection(ctx context.Context, schema *entity.Schema, numShards int32, opts ...CreateCollectionOption) error {
	return nil
}
