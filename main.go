package main

import (
	"log"

	"github.com/matsu-chara/suburi-graphql/server"
)

func main() {
	_, err := server.NewGraphQLServer(cfg.RedisURL)
	if err != nil {
		log.Fatal(err)
	}
	return
}
