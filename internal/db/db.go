package db

import (
	"fmt"
	"log"

	"github.com/yash-sojitra/todo/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresStorage() (*gorm.DB, error) {

	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", config.Envs.DBUser, config.Envs.DBPassword, config.Envs.DBAddress, config.Envs.DBName)
	fmt.Println(connStr)
	
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		fmt.Println(connStr)
		log.Fatal(err)
	}

	return db, nil
}
