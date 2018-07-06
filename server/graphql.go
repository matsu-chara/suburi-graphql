//go:generate gqlgen -schema ./schema.graphql
package server

import (
	"sync"
	"time"

	"github.com/go-redis/redis"
	"github.com/tinrab/retry"
)

type graphQLServer struct {
	redisClient     *redis.Client
	messageChannels map[string]chan Message
	userChannels    map[string]chan string
	mutex           sync.Mutex
}

func NewGraphQLServer(redisURL string) (*graphQLServer, error) {
	client := redis.NewClient(&redis.Options{
		Addr: redisURL,
	})

	retry.ForeverSleep(2*time.Second, func(_ int) error {
		_, err := client.Ping().Result()
		return err
	})
	return &graphQLServer{
		redisClient:     client,
		messageChannels: map[string]chan Message{},
		userChannels:    map[string]chan string{},
		mutex:           sync.Mutex{},
	}, nil
}
