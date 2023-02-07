package client

import (
	"context"
)

// This method creates a named partition in the specified collection.
//
// # Parameters
//
//   - collectionName
//
//     Specifies the name of a collection.
//
//     A collection name should be a string of 1 to 255 characters, starting with a letter or an underscore (_) and containing only numbers, letters, and underscores (_).
//
//   - partitionName
//
//     Specifies the name of a partition.
//
//     A partition name should be a string of 1 to 255 characters, starting with a letter or an underscore (_) and containing only numbers, letters, and underscores (_).
//
// # Returns
//
// In normal cases, this method returns nil, indicating that this operation succeeds. For other values, see "Possible Errors".
//
// # Possible Errors
func (c *GrpcClient) CreatePartition(ctx context.Context, collectionName string, partitionName string) error {
	return nil
}
