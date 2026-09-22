package cmd

import (
	"log"
	"net/http"
	controller "github.com/Devrao-2006/Go-Limit/controller"
	redis "github.com/Devrao-2006/Go-Limit/config/redis-config"
)

func main() {

	if err := redis.InitClients(); err != nil {
		log.Fatalf("Redis Clients were Not initialized: %v", err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/fixed-rate-limit", controller.FixedRateLimitHandler)

	err := http.ListenAndServe("8000", mux)

	if err != nil {
		log.Fatalf("Server Could not start")
	}
}
