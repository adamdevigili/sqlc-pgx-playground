-- name: GetTeam :one
SELECT * FROM team
WHERE id = $1 LIMIT 1;

-- name: ListTeams :many
SELECT * FROM team
ORDER BY name;

-- name: ListTeamsForPlayer :many
SELECT t.* 
	FROM player p
	JOIN team t ON t.id = ANY(p.teams)
	WHERE p.id = $1;

-- name: CreateTeam :one
INSERT INTO team (
  id, name, sport_name, power_score
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: DeleteTeam :exec
DELETE FROM team
WHERE id = $1;

-- name: DeleteAllTeams :exec
DELETE FROM team;