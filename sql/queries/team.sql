-- name: GetTeam :one
SELECT * FROM team
WHERE id = $1 LIMIT 1;

-- name: ListTeams :many
SELECT * FROM team
ORDER BY name;

-- name: ListTeamsByPlayerID :many
SELECT *
  FROM player_team pt
  JOIN team ON team.id = pt.team_id
  WHERE pt.player_id = $1;

-- name: CreateTeam :one
INSERT INTO team (
  id, name
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteTeam :exec
DELETE FROM team
WHERE id = $1;

-- name: DeleteAllTeams :exec
DELETE FROM team;