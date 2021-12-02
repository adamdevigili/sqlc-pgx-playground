package main

import (
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/labstack/gommon/log"
	"github.com/pioz/faker"
	"tutorial.sqlc.dev/app/pkg/models"
)

var playersToGenerate = 30

func generateSeedData() (
	[]*models.CreatePlayerParams,
	[]*models.CreateTeamParams,
	[]*models.AddPlayerToTeamParams,
	[]*models.CreateSkillParams,
	[]*models.CreateSportParams,
) {

	generateSeedSports()
	seedPlayers := generateSeedPlayers()
	// seedPlayerSkills := generateSeedPlayerSkills(seedPlayers)
	seedTeams := generateSeedTeams(seedPlayers, seedSports[0])
	seedPlayerTeams := generateSeedPlayerTeams(seedPlayers, seedTeams)
	calculateTeamPowerScores(seedPlayers, seedTeams, seedPlayerTeams)

	log.Info(seedPlayers[0])

	return playerToCreateParams(seedPlayers),
		teamToCreateParams(seedTeams),
		playerTeamToCreateParams(seedPlayerTeams),
		skillToCreateParams(seedSkills),
		sportToCreateParams(seedSports)
	// playerSkillToCreateParams(seedPlayerSkills)
}

// A skill for a player
// type playerSkills struct {
// 	ID    uuid.UUID `json:"id"`
// 	Name  string    `json:"name"`
// 	Value int16     `json:"value"`
// }

type playerSkills map[string]int16

// A map of calculated power scores sport name -> power score
type pScoreJSON map[string]float64

func generateSeedPlayers() []*models.Player {
	faker.SetSeed(623)
	players := make([]*models.Player, playersToGenerate)

	for i := range players {
		fn, ln := faker.FirstName(), faker.LastName()
		p := &models.Player{
			FirstName: fn,
			LastName:  ln,
			Name:      fn + " " + ln,
			ID:        uuid.New(),
		}

		sJ := playerSkills{}
		for _, s := range seedSkills {
			// sJ = append(sJ, playerSkills{
			// 	ID:    s.ID,
			// 	Name:  s.Name,
			// 	Value: int16(faker.IntInRange(1, 10)),
			// })
			sJ[s.Name] = int16(faker.IntInRange(1, 10))

		}

		// log.Infof("%+v", sJ)

		// sJBytes, _ := json.Marshal(sJ)
		p.Skills.Set(sJ)

		pS := pScoreJSON{}

		for _, s := range seedSports {
			pS[s.Name] = calcPlayerPowerScoreForSport(p, s)
		}

		p.PowerScores.Set(pS)

		players[i] = p
	}

	// log.Info("first player", players[0])

	return players
}

func getPlayerPowerScore(p pgtype.JSONB, s string) float64 {
	return p.Get().(map[string]interface{})[s].(float64)
}

func calcPlayerPowerScoreForSport(player *models.Player, sport *models.Sport) float64 {
	var powerScore float64

	weights := sport.SkillWeights.Get().(map[string]interface{})
	skills := player.Skills.Get().(map[string]interface{})

	// "handling", 0.5
	for skill, weight := range weights {
		s, ok := skills[skill]
		if ok {
			powerScore += s.(float64) * weight.(float64)
		}
	}

	return powerScore
}

func generateSeedTeams(seedPlayers []*models.Player, seedSport *models.Sport) []*models.Team {
	teamSize := seedSport.MaxPlayersPerTeam
	numTeams := len(seedPlayers) / int(teamSize)

	log.Infof("creating %d teams for sport %s", numTeams, seedSport.Name)

	teams := make([]*models.Team, numTeams)
	//

	// log.Infof("%+v", teams, teamSize, numTeams, len(seedPlayers))

	for i := 0; i < len(teams); i++ {
		teams[i] = &models.Team{
			Name:      faker.ColorName(),
			ID:        uuid.New(),
			SportName: seedSport.Name,
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

func calculateTeamPowerScores(seedPlayers []*models.Player, seedTeams []*models.Team, seedPlayerTeams []*models.PlayerTeam) {
	for _, pt := range seedPlayerTeams {
		for _, t := range seedTeams {
			if pt.TeamID == t.ID {
				for _, p := range seedPlayers {
					if pt.PlayerID == p.ID {
						t.PowerScore += float32(getPlayerPowerScore(p.PowerScores, t.SportName))
					}
				}
			}
		}
	}
}

func generateSeedPlayerSkills(seedPlayers []*models.Player) []*models.PlayerSkill {
	playerSkills := make([]*models.PlayerSkill, len(seedPlayers)*len(seedSkills))

	log.Info("length of playerskills", len(playerSkills))

	for i := 0; i < len(seedPlayers); i++ {
		for _, s := range seedSkills {
			playerSkills = append(playerSkills, &models.PlayerSkill{
				PlayerID: seedPlayers[i].ID,
				SkillID:  s.ID,
				Value:    int16(faker.IntInRange(1, 10)),
			})
		}
	}

	return playerSkills
}

func generateSeedSports() {
	for _, s := range seedSports {
		s.SkillWeights.Set(seedSkillWeights[s.Name])
	}
}

func playerToCreateParams(players []*models.Player) []*models.CreatePlayerParams {
	r := make([]*models.CreatePlayerParams, len(players))
	for i, p := range players {
		r[i] = &models.CreatePlayerParams{
			ID:          p.ID,
			Name:        p.Name,
			FirstName:   p.FirstName,
			LastName:    p.LastName,
			Skills:      p.Skills,
			PowerScores: p.PowerScores,
		}
	}

	return r
}

func teamToCreateParams(teams []*models.Team) []*models.CreateTeamParams {
	r := make([]*models.CreateTeamParams, len(teams))
	for i, t := range teams {
		r[i] = &models.CreateTeamParams{
			ID:         t.ID,
			Name:       t.Name,
			SportName:  t.SportName,
			PowerScore: t.PowerScore,
		}
	}

	return r
}

func playerTeamToCreateParams(playerTeams []*models.PlayerTeam) []*models.AddPlayerToTeamParams {
	r := make([]*models.AddPlayerToTeamParams, len(playerTeams))
	for i, t := range playerTeams {
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

func sportToCreateParams(sports []*models.Sport) []*models.CreateSportParams {
	r := make([]*models.CreateSportParams, len(sports))
	for i, s := range sports {
		r[i] = &models.CreateSportParams{
			ID:                      s.ID,
			Name:                    s.Name,
			Description:             s.Description,
			MaxActivePlayersPerTeam: s.MaxActivePlayersPerTeam,
			MaxPlayersPerTeam:       s.MaxPlayersPerTeam,
			SkillWeights:            s.SkillWeights,
		}
	}

	return r
}

// func playerSkillToCreateParams(playerSkills []*models.PlayerSkill) []*models.AddSkillToPlayerParams {
// 	r := make([]*models.AddSkillToPlayerParams, len(playerSkills))
// 	log.Info(len(r))
// 	for i, s := range playerSkills {
// 		r[i] = &models.AddSkillToPlayerParams{
// 			PlayerID: s.PlayerID,
// 			SkillID:  s.SkillID,
// 			Value:    s.Value,
// 		}
// 	}

// 	return r
// }

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

type skillWeights map[string]interface{}

var (
	seedSkillWeights = map[string]skillWeights{
		ultimateFrisbee.Name: {
			"Handling": 0.9,
			"Speed":    0.8,
			"Stamina":  0.8,
			"Height":   0.4,
		},
		football.Name: {
			"Strength": 0.7,
			"Speed":    0.8,
			"Stamina":  0.8,
			"Agility":  0.5,
		},

		basketball.Name: {
			"Shooting": 0.9,
			"Speed":    0.6,
			"Stamina":  0.8,
			"Height":   0.8,
			"Passing":  0.5,
		},
	}

	ultimateFrisbee = models.Sport{
		Name: "Ultimate Frisbee",
		ID:   uuid.New(),

		// SkillWeights: pgtype.JSONB{
		// 	Bytes: map[string]float64{
		// 		"handling": 0.9,
		// 		"speed":    0.8,
		// 		"stamina":  0.8,
		// 		"height":   0.4,
		// 	},
		// },
		MaxPlayersPerTeam:       15,
		MaxActivePlayersPerTeam: 7,
	}

	football = models.Sport{
		Name: "Football",
		ID:   uuid.New(),
		// SkillWeights: models.SkillWeightMap{
		// 	"strength": 0.7,
		// 	"speed":    0.8,
		// 	"stamina":  0.8,
		// 	"agility":  0.5,
		// },
		MaxPlayersPerTeam:       50,
		MaxActivePlayersPerTeam: 11,
	}

	basketball = models.Sport{
		Name: "Basketball",
		ID:   uuid.New(),

		// SkillWeights: models.SkillWeightMap{
		// "shooting": 0.9,
		// "speed":    0.6,
		// "stamina":  0.8,
		// "height":   0.8,
		// "passing":  0.5,
		// },
		MaxPlayersPerTeam:       12,
		MaxActivePlayersPerTeam: 5,
	}

	seedSports = []*models.Sport{
		&ultimateFrisbee,
		&basketball,
		&football,
	}
)
