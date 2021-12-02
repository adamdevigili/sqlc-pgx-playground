-- name: GetSport :one
SELECT * FROM sport
WHERE id = $1 LIMIT 1;

-- name: ListSports :many
SELECT * FROM sport
ORDER BY name;

-- name: CreateSport :one
INSERT INTO sport (
  id, name, description, skill_weights, max_active_players_per_team, max_players_per_team
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: DeleteSport :exec
DELETE FROM sport
WHERE id = $1;

-- name: DeleteAllSports :exec
DELETE FROM sport;