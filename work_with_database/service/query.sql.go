// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: query.sql

package service

import (
	"context"
)

const getPassportData = `-- name: getPassportData :one
SELECT id, series, number FROM passports p
WHERE p.id = $1
`

func (q *Queries) getPassportData(ctx context.Context, id string) (Passport, error) {
	row := q.queryRow(ctx, q.getPassportDataStmt, getPassportData, id)
	var i Passport
	err := row.Scan(&i.ID, &i.Series, &i.Number)
	return i, err
}

const insertPassportData = `-- name: insertPassportData :exec
INSERT INTO passports (id, series, number)
VALUES ($1, $2, $3)
`

type insertPassportDataParams struct {
	ID     string
	Series string
	Number string
}

func (q *Queries) insertPassportData(ctx context.Context, arg insertPassportDataParams) error {
	_, err := q.exec(ctx, q.insertPassportDataStmt, insertPassportData, arg.ID, arg.Series, arg.Number)
	return err
}
