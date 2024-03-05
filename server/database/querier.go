// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import (
	"context"
)

type Querier interface {
	BucketCount(ctx context.Context) (int64, error)
	BucketCreate(ctx context.Context, arg *BucketCreateParams) (string, error)
	BucketDelete(ctx context.Context, id string) error
	BucketDisable(ctx context.Context, id string) error
	BucketEnable(ctx context.Context, id string) error
	BucketGetById(ctx context.Context, id string) (*StorageBucket, error)
	BucketGetByName(ctx context.Context, name string) (*StorageBucket, error)
	BucketGetObjectCountById(ctx context.Context, id string) (*BucketGetObjectCountByIdRow, error)
	BucketGetSizeById(ctx context.Context, id string) (*BucketGetSizeByIdRow, error)
	BucketListAll(ctx context.Context) ([]*StorageBucket, error)
	BucketListPaginated(ctx context.Context, arg *BucketListPaginatedParams) ([]*StorageBucket, error)
	BucketLock(ctx context.Context, arg *BucketLockParams) error
	BucketSearch(ctx context.Context, name string) ([]*StorageBucket, error)
	BucketUnlock(ctx context.Context, id string) error
	BucketUpdate(ctx context.Context, arg *BucketUpdateParams) error
	EventCreate(ctx context.Context, arg *EventCreateParams) (string, error)
	ObjectCreate(ctx context.Context, arg *ObjectCreateParams) (string, error)
	ObjectDelete(ctx context.Context, id string) error
	ObjectGetByBucketIdAndId(ctx context.Context, arg *ObjectGetByBucketIdAndIdParams) (*StorageObject, error)
	ObjectGetById(ctx context.Context, id string) (*StorageObject, error)
	ObjectGetByIdWithBucketName(ctx context.Context, id string) (*ObjectGetByIdWithBucketNameRow, error)
	ObjectGetByName(ctx context.Context, name string) (*StorageObject, error)
	ObjectSearchByBucketIdAndObjectPath(ctx context.Context, arg *ObjectSearchByBucketIdAndObjectPathParams) ([]*StorageObject, error)
	ObjectUpdate(ctx context.Context, arg *ObjectUpdateParams) error
	ObjectUpdateLastAccessedAt(ctx context.Context, id string) error
	ObjectUpdateUploadStatus(ctx context.Context, arg *ObjectUpdateUploadStatusParams) error
	ObjectsListBucketIdPaged(ctx context.Context, arg *ObjectsListBucketIdPagedParams) ([]*ObjectsListBucketIdPagedRow, error)
}

var _ Querier = (*Queries)(nil)
