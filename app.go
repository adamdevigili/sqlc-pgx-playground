package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/jackc/pgx/v4"
	"tutorial.sqlc.dev/app/pkg/models"
)

func run() error {
	ctx := context.Background()

	// db, err := sql.Open("postgres", "user=postgres dbname=postgres sslmode=disable")
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

	queries := models.New(db)

	// delete existing players
	// err = queries.DeleteAllPlayers(ctx)
	// if err != nil {
	// 	return err
	// }

	// // delete existing players
	// err = queries.DeleteAllTeams(ctx)
	// if err != nil {
	// 	return err
	// }

	// // delete existing players
	// err = queries.DeleteAllPlayerTeams(ctx)
	// if err != nil {
	// 	return err
	// }

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
		// handle file there
		fmt.Println(schema.Name())

		path := filepath.Join(schemaDir, schema.Name())
		c, _ := ioutil.ReadFile(path)
		_, err := db.Exec(ctx, string(c))
		if err != nil {
			return err
		}
	}

	players, teams, playerTeams, skills := generateSeedData()

	for _, s := range skills {
		_, err := queries.CreateSkill(ctx, *s)
		if err != nil {
			return err
		}
	}

	for _, p := range players {
		_, err := queries.CreatePlayer(ctx, *p)
		if err != nil {
			return err
		}
	}

	for _, t := range teams {
		_, err := queries.CreateTeam(ctx, *t)
		if err != nil {
			return err
		}
	}

	for _, pt := range playerTeams {
		err := queries.AddPlayerToTeam(ctx, *pt)
		if err != nil {
			return err
		}
	}

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
	q, _ := queries.ListPlayersByTeamID(ctx, teams[0].ID)
	log.Printf("%+v", q)
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
