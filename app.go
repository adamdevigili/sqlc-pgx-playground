package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"

	"github.com/labstack/gommon/log"
	"tutorial.sqlc.dev/app/pkg/models"
)

func run() error {
	ctx := context.Background()

	// db, err := sql.Open("postgres", "user=postgres dbname=postgres sslmode=disable")
	// if err != nil {
	// 	return err
	// }

	// db, err := sql.Open("postgres", "user=postgres dbname=postgres sslmode=disable password=elite360")
	// if err != nil {
	// 	return err
	// }

	connURL := "postgres://postgres:elite360@localhost:5432/postgres"
	db, err := pgx.Connect(context.Background(), connURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close(context.Background())

	// goose.AddMigration()

	q := models.New(db)

	// delete existing players
	q.DeleteAllPlayers(ctx)

	// delete existing players
	q.DeleteAllTeams(ctx)

	q.DeleteAllSkills(ctx)

	q.DeleteAllSports(ctx)

	// delete existing players
	q.DeleteAllPlayerTeams(ctx)

	// insert seed players
	// players := generateSeedPlayers()

	// for _, p := range players {
	// 	_, err := queries.CreatePlayer(ctx, models.CreatePlayerParams{
	// 		ID:        p.ID,
	// 		FirstName: p.FirstName,
	// 		LastName:  p.LastName,
	// 		Name:      p.Name,
	// 	})
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	schemaDir := "./sql/schema"
	schemas, _ := ioutil.ReadDir(schemaDir)
	for _, schema := range schemas {
		path := filepath.Join(schemaDir, schema.Name())
		c, _ := ioutil.ReadFile(path)

		// log.Info(schema)

		// Init bridge tables last
		if strings.Contains(schema.Name(), "_") {
			defer db.Exec(ctx, string(c))
		} else {
			_, err := db.Exec(ctx, string(c))
			if err != nil {
				return err
			}
		}

	}

	players, teams, playerTeams, skills, sports := generateSeedData()

	log.Info("adding skills")
	for _, s := range skills {
		_, err := q.CreateSkill(ctx, *s)
		if err != nil {
			return err
		}
	}

	log.Info("adding sports")
	for _, s := range sports {
		_, err := q.CreateSport(ctx, *s)
		if err != nil {
			return err
		}
	}

	log.Info("adding players")
	log.Infof("%+v", players[0])

	for _, p := range players {
		_, err := q.CreatePlayer(ctx, *p)
		if err != nil {
			return err
		}
	}

	// log.Info("adding playerSkills")
	// for _, s := range playerSkills {
	// 	err := queries.AddSkillToPlayer(ctx, *s)
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	log.Info("adding teams")
	for _, t := range teams {
		_, err := q.CreateTeam(ctx, *t)
		if err != nil {
			return err
		}
	}

	log.Info("adding playerTeams")
	for _, pt := range playerTeams {
		err := q.AddPlayerToTeam(ctx, *pt)
		if err != nil {
			return err
		}
	}

	fetchedPlayer, err := q.GetPlayer(ctx, players[0].ID)
	if err != nil {
		return err
	}
	fP, _ := json.Marshal(fetchedPlayer)
	log.Info(string(fP))

	// log.Println(insertedPlayer)

	// get the author we just inserted
	// fetchedAuthor, err := queries.GetPlayer(ctx, insertedPlayer.ID)
	// if err != nil {
	// 	return err
	// }

	// create a team
	// insertedTeam, err := queries.CreateTeam(ctx, models.CreateTeamParams{
	// 	ID:   uuid.New(),
	// 	Name: "Test Team",
	// })
	// if err != nil {
	// 	return err
	// }
	// log.Println(insertedTeam)

	// // create a playerteam
	// err = queries.AddPlayerToTeam(ctx, models.AddPlayerToTeamParams{
	// 	PlayerID: insertedPlayer.ID,
	// 	TeamID:   insertedTeam.ID,
	// })
	// if err != nil {
	// 	return err
	// }
	// log.Println("association created")

	// prints true
	// log.Println(reflect.DeepEqual(insertedPlayer, fetchedAuthor))
	// q, _ := queries.ListPlayersByTeamID(ctx, teams[0].ID)
	// log.Printf("%+v", q)
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
