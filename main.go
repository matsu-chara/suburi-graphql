package main

import (
	"log"

	"github.com/matsu-chara/suburi-graphql/server"
)

func main() {
	_, err := server.NewGraphQLServer("")
	if err != nil {
		log.Fatal(err)
	}
	return
}
