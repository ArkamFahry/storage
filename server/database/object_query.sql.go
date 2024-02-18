// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: object_query.sql

package database

import (
	"context"
	"time"
)

const objectCreate = `-- name: ObjectCreate :one
insert into storage.objects
    (bucket_id, name, mime_type, size, metadata, upload_status)
values ($1,
        $2,
        $3,
        $4,
        $5,
        $6)
returning id
`

type ObjectCreateParams struct {
	BucketID     string
	Name         string
	ContentType  *string
	Size         int64
	Metadata     []byte
	UploadStatus string
}

func (q *Queries) ObjectCreate(ctx context.Context, arg *ObjectCreateParams) (string, error) {
	row := q.db.QueryRow(ctx, objectCreate,
		arg.BucketID,
		arg.Name,
		arg.ContentType,
		arg.Size,
		arg.Metadata,
		arg.UploadStatus,
	)
	var id string
	err := row.Scan(&id)
	return id, err
}

const objectDelete = `-- name: ObjectDelete :exec
delete
from storage.objects
where id = $1
`

func (q *Queries) ObjectDelete(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, objectDelete, id)
	return err
}

const objectGetByBucketIdAndName = `-- name: ObjectGetByBucketIdAndName :one
select id,
       version,
       bucket_id,
       name,
       mime_type,
       size,
       metadata,
       upload_status,
       last_accessed_at,
       created_at,
       updated_at
from storage.objects
where bucket_id = $1
  and name = $2
limit 1
`

type ObjectGetByBucketIdAndNameParams struct {
	BucketID string
	Name     string
}

type ObjectGetByBucketIdAndNameRow struct {
	ID             string
	Version        int32
	BucketID       string
	Name           string
	MimeType       string
	Size           int64
	Metadata       []byte
	UploadStatus   string
	LastAccessedAt *time.Time
	CreatedAt      time.Time
	UpdatedAt      *time.Time
}

func (q *Queries) ObjectGetByBucketIdAndName(ctx context.Context, arg *ObjectGetByBucketIdAndNameParams) (*ObjectGetByBucketIdAndNameRow, error) {
	row := q.db.QueryRow(ctx, objectGetByBucketIdAndName, arg.BucketID, arg.Name)
	var i ObjectGetByBucketIdAndNameRow
	err := row.Scan(
		&i.ID,
		&i.Version,
		&i.BucketID,
		&i.Name,
		&i.MimeType,
		&i.Size,
		&i.Metadata,
		&i.UploadStatus,
		&i.LastAccessedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const objectGetById = `-- name: ObjectGetById :one
select id,
       version,
       bucket_id,
       name,
       mime_type,
       size,
       metadata,
       upload_status,
       last_accessed_at,
       created_at,
       updated_at
from storage.objects
where id = $1
limit 1
`

type ObjectGetByIdRow struct {
	ID             string
	Version        int32
	BucketID       string
	Name           string
	MimeType       string
	Size           int64
	Metadata       []byte
	UploadStatus   string
	LastAccessedAt *time.Time
	CreatedAt      time.Time
	UpdatedAt      *time.Time
}

func (q *Queries) ObjectGetById(ctx context.Context, id string) (*ObjectGetByIdRow, error) {
	row := q.db.QueryRow(ctx, objectGetById, id)
	var i ObjectGetByIdRow
	err := row.Scan(
		&i.ID,
		&i.Version,
		&i.BucketID,
		&i.Name,
		&i.MimeType,
		&i.Size,
		&i.Metadata,
		&i.UploadStatus,
		&i.LastAccessedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const objectGetByIdWithBucketName = `-- name: ObjectGetByIdWithBucketName :one
select object.id,
       object.version,
       object.bucket_id,
       bucket.name as bucket_name,
       object.name,
       object.mime_type,
       object.size,
       object.metadata,
       object.upload_status,
       object.last_accessed_at,
       object.created_at,
       object.updated_at
from storage.objects as object
         inner join storage.buckets as bucket on object.bucket_id = bucket.id
where object.id = $1
limit 1
`

type ObjectGetByIdWithBucketNameRow struct {
	ID             string
	Version        int32
	BucketID       string
	BucketName     string
	Name           string
	MimeType       string
	Size           int64
	Metadata       []byte
	UploadStatus   string
	LastAccessedAt *time.Time
	CreatedAt      time.Time
	UpdatedAt      *time.Time
}

func (q *Queries) ObjectGetByIdWithBucketName(ctx context.Context, id string) (*ObjectGetByIdWithBucketNameRow, error) {
	row := q.db.QueryRow(ctx, objectGetByIdWithBucketName, id)
	var i ObjectGetByIdWithBucketNameRow
	err := row.Scan(
		&i.ID,
		&i.Version,
		&i.BucketID,
		&i.BucketName,
		&i.Name,
		&i.MimeType,
		&i.Size,
		&i.Metadata,
		&i.UploadStatus,
		&i.LastAccessedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const objectGetByName = `-- name: ObjectGetByName :one
select id,
       version,
       bucket_id,
       name,
       mime_type,
       size,
       metadata,
       upload_status,
       last_accessed_at,
       created_at,
       updated_at
from storage.objects
where name = $1
limit 1
`

type ObjectGetByNameRow struct {
	ID             string
	Version        int32
	BucketID       string
	Name           string
	MimeType       string
	Size           int64
	Metadata       []byte
	UploadStatus   string
	LastAccessedAt *time.Time
	CreatedAt      time.Time
	UpdatedAt      *time.Time
}

func (q *Queries) ObjectGetByName(ctx context.Context, name string) (*ObjectGetByNameRow, error) {
	row := q.db.QueryRow(ctx, objectGetByName, name)
	var i ObjectGetByNameRow
	err := row.Scan(
		&i.ID,
		&i.Version,
		&i.BucketID,
		&i.Name,
		&i.MimeType,
		&i.Size,
		&i.Metadata,
		&i.UploadStatus,
		&i.LastAccessedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const objectSearchByBucketIdAndObjectPath = `-- name: ObjectSearchByBucketIdAndObjectPath :many
select object.id,
       object.version,
       object.bucket_id,
       object.name,
       object.mime_type,
       object.size,
       object.metadata,
       object.upload_status,
       object.last_accessed_at,
       object.created_at,
       object.updated_at
from storage.objects as object
where object.bucket_id = $1
  and object.name ilike '%' || $2::text || '%'
limit $4 offset $3
`

type ObjectSearchByBucketIdAndObjectPathParams struct {
	BucketID   string
	ObjectPath string
	Offset     int32
	Limit      int32
}

type ObjectSearchByBucketIdAndObjectPathRow struct {
	ID             string
	Version        int32
	BucketID       string
	Name           string
	MimeType       string
	Size           int64
	Metadata       []byte
	UploadStatus   string
	LastAccessedAt *time.Time
	CreatedAt      time.Time
	UpdatedAt      *time.Time
}

func (q *Queries) ObjectSearchByBucketIdAndObjectPath(ctx context.Context, arg *ObjectSearchByBucketIdAndObjectPathParams) ([]*ObjectSearchByBucketIdAndObjectPathRow, error) {
	rows, err := q.db.Query(ctx, objectSearchByBucketIdAndObjectPath,
		arg.BucketID,
		arg.ObjectPath,
		arg.Offset,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*ObjectSearchByBucketIdAndObjectPathRow
	for rows.Next() {
		var i ObjectSearchByBucketIdAndObjectPathRow
		if err := rows.Scan(
			&i.ID,
			&i.Version,
			&i.BucketID,
			&i.Name,
			&i.MimeType,
			&i.Size,
			&i.Metadata,
			&i.UploadStatus,
			&i.LastAccessedAt,
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

const objectUpdate = `-- name: ObjectUpdate :exec
update storage.objects
set size      = coalesce($1, size),
    mime_type = coalesce($2, mime_type),
    metadata  = coalesce($3, metadata)
where id = $4
`

type ObjectUpdateParams struct {
	Size     *int64
	MimeType *string
	Metadata []byte
	ID       string
}

func (q *Queries) ObjectUpdate(ctx context.Context, arg *ObjectUpdateParams) error {
	_, err := q.db.Exec(ctx, objectUpdate,
		arg.Size,
		arg.MimeType,
		arg.Metadata,
		arg.ID,
	)
	return err
}

const objectUpdateLastAccessedAt = `-- name: ObjectUpdateLastAccessedAt :exec
update storage.objects
set last_accessed_at = now()
where id = $1
`

func (q *Queries) ObjectUpdateLastAccessedAt(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, objectUpdateLastAccessedAt, id)
	return err
}

const objectUpdateUploadStatus = `-- name: ObjectUpdateUploadStatus :exec
update storage.objects
set upload_status = $1
where id = $2
`

type ObjectUpdateUploadStatusParams struct {
	UploadStatus string
	ID           string
}

func (q *Queries) ObjectUpdateUploadStatus(ctx context.Context, arg *ObjectUpdateUploadStatusParams) error {
	_, err := q.db.Exec(ctx, objectUpdateUploadStatus, arg.UploadStatus, arg.ID)
	return err
}

const objectsListBucketIdPaged = `-- name: ObjectsListBucketIdPaged :many
select id,
       bucket_id,
       name,
       mime_type,
       size,
       metadata,
       upload_status,
       last_accessed_at,
       created_at,
       updated_at
from storage.objects
where bucket_id = $1
limit $3 offset $2
`

type ObjectsListBucketIdPagedParams struct {
	BucketID string
	Offset   int32
	Limit    int32
}

type ObjectsListBucketIdPagedRow struct {
	ID             string
	BucketID       string
	Name           string
	MimeType       string
	Size           int64
	Metadata       []byte
	UploadStatus   string
	LastAccessedAt *time.Time
	CreatedAt      time.Time
	UpdatedAt      *time.Time
}

func (q *Queries) ObjectsListBucketIdPaged(ctx context.Context, arg *ObjectsListBucketIdPagedParams) ([]*ObjectsListBucketIdPagedRow, error) {
	rows, err := q.db.Query(ctx, objectsListBucketIdPaged, arg.BucketID, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*ObjectsListBucketIdPagedRow
	for rows.Next() {
		var i ObjectsListBucketIdPagedRow
		if err := rows.Scan(
			&i.ID,
			&i.BucketID,
			&i.Name,
			&i.MimeType,
			&i.Size,
			&i.Metadata,
			&i.UploadStatus,
			&i.LastAccessedAt,
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
