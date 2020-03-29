package main

import (
	"log"
	"os"
	"strconv"

	DBFactory "github.com/femonofsky/go-complete-learning/api/db"
	ServerFactory "github.com/femonofsky/go-complete-learning/api/server"
)

var port int

var dbHost string
var dbPost string
var dbUser string
var dbPassword string
var dbDatabase string

func init() {
	rawPort := os.Getenv("PORT")

	if len(rawPort) > 0 {
		var err error
		port, err = strconv.Atoi(rawPort)
		if err != nil {
			panic(err)
		}
	} else {
		port = 8080
	}

	dbHost = os.Getenv("DB_HOST")
	if len(dbHost) < 1 {
		dbHost = "localhost"
	}
	dbPost = os.Getenv("DB_PORT")
	if len(dbPost) < 1 {
		dbPost = "5432"
	}
	dbUser = os.Getenv("DB_USER")
	if len(dbUser) < 1 {
		dbUser = "postgres"
	}
	dbPassword = os.Getenv("DB_PASSWORD")
	if len(dbPassword) < 1 {
		dbPassword = ""
	}
	dbDatabase = os.Getenv("DB_DATABASE")
	if len(dbDatabase) < 1 {
		dbDatabase = "appconnectdb"
	}

}

func main() {

	db, err := DBFactory.Connect(dbHost, dbPost, dbDatabase, dbUser, dbPassword)

	if err != nil {
		panic(err)
	}
	log.Println("Application has ran!")

	server := ServerFactory.NewServer(port, db)
	server.Start()

}
