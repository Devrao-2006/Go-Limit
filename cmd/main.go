package cmd

import (
	"log"
	"net/http"
	"../controller"
)

func main() {

	

	mux := http.NewServeMux()

	mux.HandleFunc("/fixed-rate-limit", controller.FixedRateLimitHandler)

	err := http.ListenAndServe("8000", mux)

	if err != nil {
		log.Fatalf("Server Could not start")
	}
}
