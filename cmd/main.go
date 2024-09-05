package main

import (
	"fmt"
	"log"

	"github.com/yash-sojitra/todo/api"
	"github.com/yash-sojitra/todo/internal/config"
	"github.com/yash-sojitra/todo/internal/db"
)

func main() {

	database, err := db.NewPostgresStorage()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connected to database")
	
	migrationHandler := db.NewMigrationHandler(database)
	migrationHandler.Migrate()
	log.Println("migrated all tables")
	

	apiServer := api.NewAPIServer(fmt.Sprintf("%s:%s",config.Envs.PublicHost,config.Envs.Port), database)
	err = apiServer.Run()
	if err != nil {
		log.Fatal(err)
	}
}
