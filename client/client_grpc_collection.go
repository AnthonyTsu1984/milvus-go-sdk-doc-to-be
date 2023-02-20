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
func (c *GrpcClient) CreateCollection(ctx context.Context, schema *entity.CollectionSchema, numShards int32, timeout float32, opts ...CreateCollectionOption) error {
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
//   - timeout
//
//     Specifies the timeout duration of this operation in seconds.
//
//     This parameter is optional. If left unspecified, no such limit applies.
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
func (c *GrpcClient) HasCollection(ctx context.Context, collectionName string, timeout float32) (bool, error) {
	return true, nil
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
//   - timeout
//
//     Specifies the timeout duration of this operation in seconds.
//
//     This parameter is optional. If left unspecified, no such limit applies.
//
// # Returns
//
// This method returns a collection model object ([entity.CollectionInfo]) and an error,
//
//   - [entity.CollectionInfo]
//
//     Lists the details of a collection.
//
//   - error
//
//     This is nil in normal cases. For other values, see "Possible Errors".
//
// # Possible Errors
func (c *GrpcClient) DescribeCollection(ctx context.Context, collectionName string, timeout float32) (entity.CollectionInfo, error) {
	a := entity.CollectionInfo{ID: 1, CollectionName: "collection"}
	return a, nil
}

// This method drops a collection with all the entities it contains.
//
// # Paramters
//
//   - collectionName
//
//     Specifies the name of a collection.
//
//     A collection name should be a string of 1 to 255 characters, starting with a letter or an underscore (_) and containing only numbers, letters, and underscores (_).
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
func (c *GrpcClient) DropCollection(ctx context.Context, collectionName string, timeout float32) error {
	return nil
}

// This method releases the loaded collection from memory. All data in the released collection remains intact after this operation. You can load the collection to memory again using LoadCollection().
//
// # Parameters
//
//   - collectionName
//
//     Specifies the name of a collection.
//
//     A collection name should be a string of 1 to 255 characters, starting with a letter or an underscore (_) and containing only numbers, letters, and underscores (_).
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
func (c *GrpcClient) ReleaseCollection(ctx context.Context, collectionName string, timeout float32) error {
	return nil
}

// This method lists all collection names in the database.
//
// # Parameters
//
//   - collectionName
//
//     Specifies the name of a collection.
//
//     A collection name should be a string of 1 to 255 characters, starting with a letter or an underscore (_) and containing only numbers, letters, and underscores (_).
//
//   - timeout
//
//     Specifies the timeout duration of this operation in seconds.
//
//     This parameter is optional. If left unspecified, no such limit applies.
//
// # Returns
//
// This method returns a list of collection names and an error,
//
//   - []string
//
//     Lists the collection names in the database.
//
//   - error
//
//     This is nil in normal cases. For other values, see "Possible Errors".
//
// # Possible Errors
func (c *GrpcClient) ListCollections(ctx context.Context, timeout float32) ([]string, error) {
	a := []string{"a", "b"}
	return a, nil
}
