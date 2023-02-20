package client

import (
	"context"

	"github.com/anthonytsu1984/gosdktobe/entity"
)

// This method describes the index of a collection.
//
// # Parameters
//
//   - collectionName
//
//     Specifies the name of a collection.
//
//     A collection name should be a string of 1 to 255 characters, starting with a letter or an underscore (_) and containing only numbers, letters, and underscores (_).
//
//   - indexName
//
//     Specifies the name of an index.
//
//     An index name should be a string of 1 to 255 characters, starting with a letter or an underscore (_) and containing only numbers, letters, and underscores (_).
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
//   - [entity.IndexInfo]
//
//     Lists the details of an index.
//
//   - error
//
//     This is nil in normal cases. For other values, see "Possible Errors".
//
// # Possible Errors
func (c *GrpcClient) DescribeIndex(ctx context.Context, collectionName string, indexName string, timeout float32) (entity.IndexInfo, error) {
	a := entity.IndexInfo{ID: 0}
	return a, nil
}

// This method drops an index from a collection.
//
// # Paramters
//
//   - collectionName
//
//     Specifies the name of a collection.
//
//     A collection name should be a string of 1 to 255 characters, starting with a letter or an underscore (_) and containing only numbers, letters, and underscores (_).
//
//   - indexName
//
//     Specifies the name of an index.
//
//     An index name should be a string of 1 to 255 characters, starting with a letter or an underscore (_) and containing only numbers, letters, and underscores (_).
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
func (c *GrpcClient) DropIndex(ctx context.Context, collectionName string, indexName string, timeout float32) error {
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
//   - fieldName
//
//     Specifies the name of the field for further reference.
//
//     A field name should be a string of 1 to 255 characters, starting with a letter or an underscore (_) and containing only numbers, letters, and underscores (_).
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
//     Lists the indexes built on the specified index.
//
//   - error
//
//     This is nil in normal cases. For other values, see "Possible Errors".
//
// # Possible Errors
func (c *GrpcClient) ListIndexes(ctx context.Context, collectionName string, fieldName string, timeout float32) ([]string, error) {
	a := []string{"a", "b"}
	return a, nil
}

// This method checks whether the specified index exists.
//
// # Parameters
//
//   - collectionName
//
//     Specifies the name of a collection.
//
//     A collection name should be a string of 1 to 255 characters, starting with a letter or an underscore (_) and containing only numbers, letters, and underscores (_).
//
//   - indexName
//
//     Specifies the name of an index.
//
//     An index name should be a string of 1 to 255 characters, starting with a letter or an underscore (_) and containing only numbers, letters, and underscores (_).
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
//     Indicates whether the specified index exists.
//
//   - error
//
//     This is nil in normal cases. For other values, see "Possible Errors".
//
// # Possible Errors
func (c *GrpcClient) HasIndex(ctx context.Context, collectionName string, indexName string, timeout float32) (bool, error) {
	return true, nil
}
