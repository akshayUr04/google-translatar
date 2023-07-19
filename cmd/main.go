package main

import (
	"fmt"
	"log"

	"github.com/akshayUr04/google-translator/pkg/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	fmt.Println("--------------")
	routes.Routing()
}
