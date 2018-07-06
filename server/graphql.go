//go:generate gqlgen -schema ./schema.graphql
package server

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"github.com/tinrab/retry"
	"github.com/vektah/gqlgen/handler"
)

type graphQLServer struct {
	redisClient     *redis.Client
	messageChannels map[string]chan Message
	userChannels    map[string]chan string
	mutex           sync.Mutex
}

// NewGraphQLServer .
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

// Serve .
func (s *graphQLServer) Serve(route string, port int) error {
	mux := http.NewServeMux()
	mux.Handle(
		route,
		handler.GraphQL(MakeExecutableSchema(s),
			handler.WebsocketUpgrader(websocket.Upgrader{
				CheckOrigin: func(r *http.Request) bool {
					return true
				},
			}),
		),
	)
	mux.Handle("/playground", handler.Playground("GraphQL", route))

	handler := cors.AllowAll().Handler(mux)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), handler)
}
