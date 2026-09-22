package constants

import "time"

var IP_HEADERS_PRIORITY_LIST = []string{"x-real-ip", "x-client-ip", "x-forwarded-for", "cf-connecting-ip", "fastly-client-ip", "true-client-ip", "x-cluster-client-ip", "x-forwarded", "forwarded-for", "forwarded", "x-appengine-user-ip"}

var Max_window_size = 5 * time.Second
var Max_requests_allowed_in_a_window = 10;
var Refill_interval = 1
var Max_bucket_ttl = Max_requests_allowed_in_a_window * Refill_interval