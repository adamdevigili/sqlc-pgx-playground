-- name: ListPlayerSkills :many
SELECT * 
    FROM player_skill
    WHERE skill_id = $1
    ORDER BY name;

-- name: AddSkillToPlayer :exec
INSERT INTO player_skill (
  player_id, skill_id, value
) VALUES (
  $1, $2, $3
);

-- name: ChangePlayerSkill :exec
UPDATE player_skill 
	SET skill_id = $1 
	WHERE player_id = $2;

-- name: DeleteAllPlayerSkills :exec
DELETE FROM player_skill;