package main

import (
	"log"
	"os"
	"strconv"

	ServerFactory "github.com/femonofsky/go-complete-learning/api/server"
)

var port int

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
}

func main() {
	log.Println("Application has ran!")
	server := ServerFactory.NewServer(port)
	server.Start()

}
