package constants

import "time"

var IP_HEADERS_PRIORITY_LIST = []string{"x-real-ip", "x-client-ip", "x-forwarded-for", "cf-connecting-ip", "fastly-client-ip", "true-client-ip", "x-cluster-client-ip", "x-forwarded", "forwarded-for", "forwarded", "x-appengine-user-ip"}

var Max_window_size = 5 * time.Second
var Max_rqeuests_allowed_in_fixed_window = 10;