package main

import (
	"flag"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"the-wedding-game-api/routes"
)

func main() {
	migrate()
	router := routes.GetRouter()

	port := flag.String("p", "8080", "port to run the server on")
	flag.Parse()

	err := router.Run(fmt.Sprintf("localhost:%s", *port))
	if err != nil {
		log.Println("Error starting server: ", err)
		return
	}

	log.Println("Server started on port ", *port)
}
