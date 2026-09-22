package cmd

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	redis "github.com/Devrao-2006/Go-Limit/config/redis-config"
	limit "github.com/Devrao-2006/Go-Limit/rate-limiters"
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
	mux.HandleFunc("/sliding-rate-limit", limit.SlidingWindowLimitHandler)
	mux.HandleFunc("/bucket-rate-limit", limit.BucketLimitHandler)

	srv := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		log.Println("Server is running on port 8000...")
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server listen failed: %v\n", err)
		}
	}()

	<-ctx.Done()
	log.Println("Shutdown signal received, shutting down gracefully...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	redis.CloseRedis()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited cleanly.")
}
