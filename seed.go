package main

import (
	"github.com/google/uuid"
	"github.com/pioz/faker"
	"tutorial.sqlc.dev/app/pkg/models"
)

var playersToGenerate = 30

func generateSeedData() ([]*models.CreatePlayerParams, []*models.CreateTeamParams, []*models.AddPlayerToTeamParams, []*models.CreateSkillParams) {

	seedPlayers := generateSeedPlayers()
	seedTeams := generateSeedTeams(seedPlayers)
	seedPlayerTeams := generateSeedPlayerTeams(seedPlayers, seedTeams)

	return playerToCreateParams(seedPlayers), teamToCreateParams(seedTeams), playerTeamToCreateParams(seedPlayerTeams), skillToCreateParams(seedSkills)
}

func generateSeedPlayers() []*models.Player {
	faker.SetSeed(623)
	players := make([]*models.Player, playersToGenerate)

	for i := range players {
		fn, ln := faker.FirstName(), faker.LastName()
		p := &models.Player{
			FirstName: fn,
			LastName:  ln,
			Name:      fn + ln,
			ID:        uuid.New(),
		}

		players[i] = p
	}

	return players
}

func generateSeedTeams(seedPlayers []*models.Player) []*models.Team {
	teamSize := 5
	numTeams := len(seedPlayers) / teamSize

	teams := make([]*models.Team, numTeams)
	//

	// log.Infof("%+v", teams, teamSize, numTeams, len(seedPlayers))

	for i := 0; i < len(teams); i++ {
		teams[i] = &models.Team{
			Name: faker.ColorName(),
			ID:   uuid.New(),
		}
	}

	return teams
}

func generateSeedPlayerTeams(seedPlayers []*models.Player, seedTeams []*models.Team) []*models.PlayerTeam {
	playerTeams := make([]*models.PlayerTeam, len(seedPlayers))

	numTeams := len(seedTeams)
	currTeam := 0

	for i := 0; i < len(seedPlayers); i++ {
		playerTeams[i] = &models.PlayerTeam{
			PlayerID: seedPlayers[i].ID,
			TeamID:   seedTeams[currTeam].ID,
		}
		currTeam += 1
		currTeam %= numTeams
	}

	return playerTeams
}

func playerToCreateParams(players []*models.Player) []*models.CreatePlayerParams {
	r := make([]*models.CreatePlayerParams, len(players))
	for i, p := range players {
		r[i] = &models.CreatePlayerParams{
			ID:        p.ID,
			Name:      p.Name,
			FirstName: p.FirstName,
			LastName:  p.LastName,
		}
	}

	return r
}

func teamToCreateParams(teams []*models.Team) []*models.CreateTeamParams {
	r := make([]*models.CreateTeamParams, len(teams))
	for i, t := range teams {
		r[i] = &models.CreateTeamParams{
			ID:   t.ID,
			Name: t.Name,
		}
	}

	return r
}

func playerTeamToCreateParams(playerteams []*models.PlayerTeam) []*models.AddPlayerToTeamParams {
	r := make([]*models.AddPlayerToTeamParams, len(playerteams))
	for i, t := range playerteams {
		r[i] = &models.AddPlayerToTeamParams{
			PlayerID: t.PlayerID,
			TeamID:   t.TeamID,
		}
	}

	return r
}

func skillToCreateParams(skills []*models.Skill) []*models.CreateSkillParams {
	r := make([]*models.CreateSkillParams, len(skills))
	for i, s := range skills {
		r[i] = &models.CreateSkillParams{
			ID:          s.ID,
			Name:        s.Name,
			Description: s.Description,
		}
	}

	return r
}

var seedSkills = []*models.Skill{
	{
		Name:        "Handling",
		ID:          uuid.New(),
		Description: "Control, as in basketball or soccer, by skillful dribbling and accurate passing",
	},
	{
		Name:        "Power",
		ID:          uuid.New(),
		Description: "Measure of explosive athleticism",
	},
	{
		Name:        "Speed",
		ID:          uuid.New(),
		Description: "How fast a player can run",
	},
	{
		Name:        "Height",
		ID:          uuid.New(),
		Description: "How tall a player is",
	},
	{
		Name:        "Stamina",
		ID:          uuid.New(),
		Description: "How long a player can remain competitive",
	},
}
