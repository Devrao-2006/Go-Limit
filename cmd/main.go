package cmd

import (
	"log"
	"net/http"
	limit "github.com/Devrao-2006/Go-Limit/rate-limiters"
	redis "github.com/Devrao-2006/Go-Limit/config/redis-config"
)

func main() {

	if err := redis.InitClients(); err != nil {
		log.Fatalf("Redis Clients were Not initialized: %v", err)
	}

	if err := limit.Init(); err != nil {
		log.Fatalf("Scripts Were Not Initialzied: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/fixed-rate-limit", limit.FixedRateLimitHandler)

	err := http.ListenAndServe("8000", mux)

	if err != nil {
		log.Fatalf("Server Could not start")
	}
}
