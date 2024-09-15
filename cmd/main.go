package main

import (
	"log"

	"github.com/llucasreis/go-dbjson-null-poc/infra/postgres"
)

func main() {
	dbConn, err := postgres.OpenConnection()
	if err != nil {
		log.Fatal("opening postgres connection: ", err)
	}
	repo := postgres.NewRepository(dbConn)

	game, err := repo.FindByID(int64(1))
	if err != nil {
		log.Fatal("finding game by id: ", err)
	}
	log.Printf("Game data: %+v \n", game)

	game2, err := repo.FindByID(int64(2))
	if err != nil {
		log.Fatal("finding game by id: ", err)
	}
	log.Printf("Game data: %+v \n", game2)

	game3, err := repo.FindByID(int64(3))
	if err != nil {
		log.Fatal("finding game by id: ", err)
	}
	log.Printf("Game data: %+v \n", game3)
}
