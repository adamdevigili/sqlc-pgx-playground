// Code generated by sqlc. DO NOT EDIT.
// source: skill.sql

package models

import (
	"context"

	"github.com/google/uuid"
)

const CreateSkill = `-- name: CreateSkill :one
INSERT INTO skill (
  id, name, description
) VALUES (
  $1, $2, $3
)
RETURNING id, name, created_at, updated_at, description
`

type CreateSkillParams struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

func (q *Queries) CreateSkill(ctx context.Context, arg CreateSkillParams) (Skill, error) {
	row := q.db.QueryRow(ctx, CreateSkill, arg.ID, arg.Name, arg.Description)
	var i Skill
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Description,
	)
	return i, err
}

const DeleteAllSkills = `-- name: DeleteAllSkills :exec
DELETE FROM skill
`

func (q *Queries) DeleteAllSkills(ctx context.Context) error {
	_, err := q.db.Exec(ctx, DeleteAllSkills)
	return err
}

const DeleteSkill = `-- name: DeleteSkill :exec
DELETE FROM skill
WHERE id = $1
`

func (q *Queries) DeleteSkill(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, DeleteSkill, id)
	return err
}

const GetSkill = `-- name: GetSkill :one
SELECT id, name, created_at, updated_at, description FROM skill
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetSkill(ctx context.Context, id uuid.UUID) (Skill, error) {
	row := q.db.QueryRow(ctx, GetSkill, id)
	var i Skill
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Description,
	)
	return i, err
}

const ListSkills = `-- name: ListSkills :many
SELECT id, name, created_at, updated_at, description FROM skill
ORDER BY name
`

func (q *Queries) ListSkills(ctx context.Context) ([]Skill, error) {
	rows, err := q.db.Query(ctx, ListSkills)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Skill
	for rows.Next() {
		var i Skill
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Description,
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
