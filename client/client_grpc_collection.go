package client

import (
	"context"

	"github.com/anthonytsu1984/gosdktobe/entity"
)

// This method creates a collection with a pre-defined schema.
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
// In normal cases, this method returns nil, indicating that this operation succeeds. For other values, see "Possible Errors".
//
// # Possible Errors
func (c *GrpcClient) CreateCollection(ctx context.Context, schema *entity.Schema, numShards int32, opts ...CreateCollectionOption) error {
	return nil
}

// This method checks whether the specified collection exists.
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
// This method returns a boolean value (bool) and an error,
//
//   - bool
//
//     Indicates whether the specified collection exists.
//
//   - error
//
//     This is nil in normal cases. For other values, see "Possible Errors".
//
// # Possible Errors
func (c *GrpcClient) HasCollection(ctx context.Context, collectionName string) (bool, error) {
	return nil
}

// This method lists the names of all collections in the connected Milvus instance.
//
// # Parameters
//
// None.
//
// # Returns
//
// This method returns a list of strings ([]string) and an error,
//
//   - []string
//
//     Indicates the list of collection names in the connected Milvus instance.
//
//   - error
//
//     This is nil in normal cases. For other values, see "Possible Errors".
//
// # Possible Errors
func (c *GrpcClient) ListCollections(ctx context.Context) ([]string, error) {
	return nil
}

// This method describes the specified collection in detail.
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
// This method returns a collection model object ([model.Collection]) and an error,
//
//   - [models.Collection]
//
//     Indicates a collection model object that contains the details of the specified collection.
//
//   - error
//
//     This is nil in normal cases. For other values, see "Possible Errors".
//
// # Possible Errors
func (c *GrpcClient) DescribeCollection(ctx context.Context, collectionName string) (models.Collection, error) {
	return nil
}
