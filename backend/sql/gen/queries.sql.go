// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: queries.sql

package sqlc

import (
	"context"
)

const doesUserExist = `-- name: DoesUserExist :one
SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)
`

func (q *Queries) DoesUserExist(ctx context.Context, username string) (bool, error) {
	row := q.db.QueryRow(ctx, doesUserExist, username)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const insertSecret = `-- name: InsertSecret :exec
INSERT INTO secrets (value, key, url, user_id)
VALUES ($1, $2, $3, $4)
`

type InsertSecretParams struct {
	Value  []byte
	Key    string
	Url    string
	UserID int32
}

func (q *Queries) InsertSecret(ctx context.Context, arg InsertSecretParams) error {
	_, err := q.db.Exec(ctx, insertSecret,
		arg.Value,
		arg.Key,
		arg.Url,
		arg.UserID,
	)
	return err
}

const insertUser = `-- name: InsertUser :exec
INSERT INTO users (username)
VALUES ($1)
`

func (q *Queries) InsertUser(ctx context.Context, username string) error {
	_, err := q.db.Exec(ctx, insertUser, username)
	return err
}

const selectSecrets = `-- name: SelectSecrets :many
SELECT id, value, key, url, user_id, created_at, updated_at FROM secrets
WHERE user_id = $1
`

func (q *Queries) SelectSecrets(ctx context.Context, userID int32) ([]Secret, error) {
	rows, err := q.db.Query(ctx, selectSecrets, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Secret
	for rows.Next() {
		var i Secret
		if err := rows.Scan(
			&i.ID,
			&i.Value,
			&i.Key,
			&i.Url,
			&i.UserID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectUserByName = `-- name: SelectUserByName :one
SELECT id, username, created_at, updated_at FROM users
WHERE username = $1
`

func (q *Queries) SelectUserByName(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRow(ctx, selectUserByName, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}