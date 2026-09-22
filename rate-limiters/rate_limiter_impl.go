package rate_limiters

import (
	"log"
	"net/http"
	"os"

	ctx "context"

	config "github.com/Devrao-2006/Go-Limit/config/redis-config"
	constants "github.com/Devrao-2006/Go-Limit/constants"
	helper "github.com/Devrao-2006/Go-Limit/helper"
)

var Scripts struct {
	Fixed_window string
}

func Init() error {
	fixed_window_script, err := os.ReadFile("./fixed_limit.lua")
	if err != nil {
		return err
	}

	Scripts.Fixed_window = string(fixed_window_script)
	
	return nil
}

func FixedRateLimitHandler(w http.ResponseWriter, r *http.Request) {
	db := config.GetClient(config.REDIS_DB)
	if db == nil {
		log.Printf("DB Not Initialized")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	ip, err := helper.GetIPfromheader(r.Header)
	if err != nil {
		http.Error(w, "Bad Request: No IP Found", http.StatusBadRequest)
		return
	}

	key := "fixed_" + ip
	cntx := ctx.Background()

	limitVal := constants.Max_requests_allowed_in_fixed_window
	windowSize := int(constants.Max_window_size.Seconds())

	result, err2 := db.Eval(cntx, Scripts.Fixed_window, []string{key}, limitVal, windowSize).Int64()
	if err2 != nil {
		log.Printf("Rate limit script error for IP %s: %v", ip, err2)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if result > int64(limitVal) {
		http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Request allowed"))
}