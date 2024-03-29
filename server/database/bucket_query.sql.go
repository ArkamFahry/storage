// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: bucket_query.sql

package database

import (
	"context"
)

const bucketCount = `-- name: BucketCount :one
select count(1) as count
from storage.buckets
`

func (q *Queries) BucketCount(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, bucketCount)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const bucketCreate = `-- name: BucketCreate :one
insert into storage.buckets
    (name, allowed_mime_types, max_allowed_object_size, public)
values ($1,
        $2,
        $3,
        $4)
returning id
`

type BucketCreateParams struct {
	Name                 string
	AllowedMimeTypes     []string
	MaxAllowedObjectSize *int64
	Public               bool
}

func (q *Queries) BucketCreate(ctx context.Context, arg *BucketCreateParams) (string, error) {
	row := q.db.QueryRow(ctx, bucketCreate,
		arg.Name,
		arg.AllowedMimeTypes,
		arg.MaxAllowedObjectSize,
		arg.Public,
	)
	var id string
	err := row.Scan(&id)
	return id, err
}

const bucketDelete = `-- name: BucketDelete :exec
delete
from storage.buckets
where id = $1
`

func (q *Queries) BucketDelete(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, bucketDelete, id)
	return err
}

const bucketDisable = `-- name: BucketDisable :exec
update storage.buckets
set disabled = true
where id = $1
`

func (q *Queries) BucketDisable(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, bucketDisable, id)
	return err
}

const bucketEnable = `-- name: BucketEnable :exec
update storage.buckets
set disabled = false
where id = $1
`

func (q *Queries) BucketEnable(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, bucketEnable, id)
	return err
}

const bucketGetById = `-- name: BucketGetById :one
select id,
       version,
       name,
       allowed_mime_types,
       max_allowed_object_size,
       public,
       disabled,
       locked,
       lock_reason,
       locked_at,
       created_at,
       updated_at
from storage.buckets
where id = $1
limit 1
`

func (q *Queries) BucketGetById(ctx context.Context, id string) (*StorageBucket, error) {
	row := q.db.QueryRow(ctx, bucketGetById, id)
	var i StorageBucket
	err := row.Scan(
		&i.ID,
		&i.Version,
		&i.Name,
		&i.AllowedMimeTypes,
		&i.MaxAllowedObjectSize,
		&i.Public,
		&i.Disabled,
		&i.Locked,
		&i.LockReason,
		&i.LockedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const bucketGetByName = `-- name: BucketGetByName :one
select id,
       version,
       name,
       allowed_mime_types,
       max_allowed_object_size,
       public,
       disabled,
       locked,
       lock_reason,
       locked_at,
       created_at,
       updated_at
from storage.buckets
where name = $1
limit 1
`

func (q *Queries) BucketGetByName(ctx context.Context, name string) (*StorageBucket, error) {
	row := q.db.QueryRow(ctx, bucketGetByName, name)
	var i StorageBucket
	err := row.Scan(
		&i.ID,
		&i.Version,
		&i.Name,
		&i.AllowedMimeTypes,
		&i.MaxAllowedObjectSize,
		&i.Public,
		&i.Disabled,
		&i.Locked,
		&i.LockReason,
		&i.LockedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const bucketGetObjectCountById = `-- name: BucketGetObjectCountById :one
select bucket_id as id, count(1) as count
from storage.objects
where bucket_id = $1
group by bucket_id
`

type BucketGetObjectCountByIdRow struct {
	ID    string
	Count int64
}

func (q *Queries) BucketGetObjectCountById(ctx context.Context, id string) (*BucketGetObjectCountByIdRow, error) {
	row := q.db.QueryRow(ctx, bucketGetObjectCountById, id)
	var i BucketGetObjectCountByIdRow
	err := row.Scan(&i.ID, &i.Count)
	return &i, err
}

const bucketGetSizeById = `-- name: BucketGetSizeById :one
select object.bucket_id as id, bucket.name as name, SUM(object.size) as size
from storage.objects as object
         join storage.buckets as bucket on object.bucket_id = bucket.id
where object.bucket_id = $1
group by object.bucket_id, bucket.name
`

type BucketGetSizeByIdRow struct {
	ID   string
	Name string
	Size int64
}

func (q *Queries) BucketGetSizeById(ctx context.Context, id string) (*BucketGetSizeByIdRow, error) {
	row := q.db.QueryRow(ctx, bucketGetSizeById, id)
	var i BucketGetSizeByIdRow
	err := row.Scan(&i.ID, &i.Name, &i.Size)
	return &i, err
}

const bucketListAll = `-- name: BucketListAll :many
select id,
       version,
       name,
       allowed_mime_types,
       max_allowed_object_size,
       public,
       disabled,
       locked,
       lock_reason,
       locked_at,
       created_at,
       updated_at
from storage.buckets
`

func (q *Queries) BucketListAll(ctx context.Context) ([]*StorageBucket, error) {
	rows, err := q.db.Query(ctx, bucketListAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*StorageBucket
	for rows.Next() {
		var i StorageBucket
		if err := rows.Scan(
			&i.ID,
			&i.Version,
			&i.Name,
			&i.AllowedMimeTypes,
			&i.MaxAllowedObjectSize,
			&i.Public,
			&i.Disabled,
			&i.Locked,
			&i.LockReason,
			&i.LockedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const bucketListPaginated = `-- name: BucketListPaginated :many
select id,
       version,
       name,
       allowed_mime_types,
       max_allowed_object_size,
       public,
       disabled,
       locked,
       lock_reason,
       locked_at,
       created_at,
       updated_at
from storage.buckets
where id >= $1
limit $2
`

type BucketListPaginatedParams struct {
	Cursor string
	Limit  int32
}

func (q *Queries) BucketListPaginated(ctx context.Context, arg *BucketListPaginatedParams) ([]*StorageBucket, error) {
	rows, err := q.db.Query(ctx, bucketListPaginated, arg.Cursor, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*StorageBucket
	for rows.Next() {
		var i StorageBucket
		if err := rows.Scan(
			&i.ID,
			&i.Version,
			&i.Name,
			&i.AllowedMimeTypes,
			&i.MaxAllowedObjectSize,
			&i.Public,
			&i.Disabled,
			&i.Locked,
			&i.LockReason,
			&i.LockedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const bucketLock = `-- name: BucketLock :exec
update storage.buckets
set locked      = true,
    lock_reason = $1::text,
    locked_at   = now()
where id = $2
`

type BucketLockParams struct {
	LockReason string
	ID         string
}

func (q *Queries) BucketLock(ctx context.Context, arg *BucketLockParams) error {
	_, err := q.db.Exec(ctx, bucketLock, arg.LockReason, arg.ID)
	return err
}

const bucketSearch = `-- name: BucketSearch :many
select id,
       version,
       name,
       allowed_mime_types,
       max_allowed_object_size,
       public,
       disabled,
       locked,
       lock_reason,
       locked_at,
       created_at,
       updated_at
from storage.buckets
where name ilike '%' || $1::text || '%'
`

func (q *Queries) BucketSearch(ctx context.Context, name string) ([]*StorageBucket, error) {
	rows, err := q.db.Query(ctx, bucketSearch, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*StorageBucket
	for rows.Next() {
		var i StorageBucket
		if err := rows.Scan(
			&i.ID,
			&i.Version,
			&i.Name,
			&i.AllowedMimeTypes,
			&i.MaxAllowedObjectSize,
			&i.Public,
			&i.Disabled,
			&i.Locked,
			&i.LockReason,
			&i.LockedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const bucketUnlock = `-- name: BucketUnlock :exec
update storage.buckets
set locked      = false,
    lock_reason = null,
    locked_at   = null
where id = $1
`

func (q *Queries) BucketUnlock(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, bucketUnlock, id)
	return err
}

const bucketUpdate = `-- name: BucketUpdate :exec
update storage.buckets
set max_allowed_object_size = coalesce($1, max_allowed_object_size),
    public                  = coalesce($2, public),
    allowed_mime_types      = coalesce($3, allowed_mime_types)
where id = $4
`

type BucketUpdateParams struct {
	MaxAllowedObjectSize *int64
	Public               *bool
	AllowedMimeTypes     []string
	ID                   string
}

func (q *Queries) BucketUpdate(ctx context.Context, arg *BucketUpdateParams) error {
	_, err := q.db.Exec(ctx, bucketUpdate,
		arg.MaxAllowedObjectSize,
		arg.Public,
		arg.AllowedMimeTypes,
		arg.ID,
	)
	return err
}
