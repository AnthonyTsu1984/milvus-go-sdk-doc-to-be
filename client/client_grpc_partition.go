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
func (c *GrpcClient) CreatePartition(ctx context.Context, collectionName string, partitionName string, timeout float32) error {
	return nil
}

// This method drops a named partition and the data it contains from the specified collection.
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
func (c *GrpcClient) DropPartition(ctx context.Context, collectionName string, partitionName string, timeout float32) error {
	return nil
}

// This method checks whether the specified partition exists.
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
//     An partition name should be a string of 1 to 255 characters, starting with a letter or an underscore (_) and containing only numbers, letters, and underscores (_).
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
//     Indicates whether the specified partition exists.
//
//   - error
//
//     This is nil in normal cases. For other values, see "Possible Errors".
//
// # Possible Errors
func (c *GrpcClient) HasPartition(ctx context.Context, collectionName string, partitionName string, timeout float32) (bool, error) {
	return true, nil
}

// This method lists the names of all partitions in the specified collection.
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
//     Lists the indexes built on the specified index.
//
//   - error
//
//     This is nil in normal cases. For other values, see "Possible Errors".
//
// # Possible Errors
func (c *GrpcClient) ListPartitions(ctx context.Context, collectionName string, timeout float32) ([]string, error) {
	a := []string{"a", "b"}
	return a, nil
}

// This method releases the specified partitions from memory. All data in the released partition remain intact. You can load them into memory again using LoadPartition().
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
//     An partition name should be a string of 1 to 255 characters, starting with a letter or an underscore (_) and containing only numbers, letters, and underscores (_).
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
func (c *GrpcClient) ReleasePartition(ctx context.Context, collectionName string, partitionName string, timeout float32) error {
	return nil
}
