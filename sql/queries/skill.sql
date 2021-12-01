-- name: GetSkill :one
SELECT * FROM skill
WHERE id = $1 LIMIT 1;

-- name: ListSkills :many
SELECT * FROM skill
ORDER BY name;

-- name: CreateSkill :one
INSERT INTO skill (
  id, name, description
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteSkill :exec
DELETE FROM skill
WHERE id = $1;

-- name: DeleteAllSkills :exec
DELETE FROM skill;