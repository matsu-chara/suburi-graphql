//go:generate gqlgen -schema ./schema.graphql
package server

import (
	"sync"

	"github.com/go-redis/redis"
)

type graphQLServer struct {
	redisClint      *redis.Client
	messageChannels map[string]chan Message
	userChannels    map[string]chan string
	mutex           sync.Mutex
}

func NewGraphQLServer(redisURL string) (*graphQLServer, error) {
	return nil, nil
}
