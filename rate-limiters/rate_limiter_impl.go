package rate_limiters

import (
	"errors"
	"log"
	"net/http"

	ctx "context"

	config "github.com/Devrao-2006/Go-Limit/config/redis-config"
	constants "github.com/Devrao-2006/Go-Limit/constants"
	helper "github.com/Devrao-2006/Go-Limit/helper"
	"github.com/redis/go-redis/v9"
)

const allowed = "ALLOWED"
const rejected = "REJECTED"
const not_processed = "NOT PROCESSED"

func FixedRateLimitHandler(w http.ResponseWriter, r *http.Request) (string, error) {
	var db *redis.Client = nil
	if db = config.GetClient(config.REDIS_DB); db == nil {
		log.Printf("DB Not Initilized")
		return not_processed, errors.New("DB Not initialzed")
	}

	ip, err := helper.GetIPfromheader(r.Header)

	if err != nil {
		return not_processed, errors.New("No IP Found")
	}

	key := "fixed_" + ip

	var val int

	cntx := ctx.Background()

	if err1 := db.Get(cntx, key).Scan(&val); err1 != nil {
		if(err1 == redis.Nil) {
			var new_val int = 1
			db.Set(cntx, key, new_val, constants.Max_window_size)
			return allowed, nil
		}
		log.Fatalf("Cannot Get the Key")
		return not_processed, errors.New("Cannnot Get the Key for this IP")
	}

	if val < constants.Max_rqeuests_allowed_in_fixed_window {
		db.Incr(cntx, key)
		return allowed, nil
	}

	return rejected, nil
}
