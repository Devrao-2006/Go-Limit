# Go-Limit

Redis-backed rate limiter service in Go. Implements three limiting algorithms, each behind its own HTTP endpoint, with the core logic in Lua scripts executed atomically via `EVAL`.

## Algorithms

| Endpoint | Algorithm | Script | Key |
|---|---|---|---|
| `POST /fixed-rate-limit` | Fixed window counter | `rate-limiters/fixed_limit.lua` | `fixed_<ip>` |
| `POST /sliding-rate-limit` | Sliding window log (sorted set) | `rate-limiters/sliding_limit.lua` | `sliding_<ip>` |
| `POST /bucket-rate-limit` | Token bucket | `rate-limiters/bucket_limit.lua` | `bucket_<ip>` |

Client identity is resolved from the first populated header in `constants.IP_HEADERS_PRIORITY_LIST` (`x-real-ip`, `x-forwarded-for`, etc. — see `helper/helper.go`). A request with none of these headers set is rejected with `400`.

Limits are configured in `constants/constants.go`:

- `Max_requests_allowed_in_a_window` — request cap per window (fixed/sliding) or bucket capacity (bucket)
- `Max_window_size` — window length for fixed/sliding
- `Refill_interval` — seconds per token refill (bucket)
- `Max_bucket_ttl` — Redis key TTL for the bucket, derived as capacity × refill interval

## Project layout

```
cmd/main.go                        entrypoint, route registration
config/config.go                   .env loading
config/redis-config/redis.go       Redis client pool, keyed by DB index
constants/constants.go             limiter tuning knobs, IP header priority
helper/helper.go                   client IP resolution
rate-limiters/rate_limiter_impl.go HTTP handlers, script loading (Init)
rate-limiters/*.lua                limiter algorithms
```

## Running locally

Requires a running Redis instance and a `.env` file in the working directory:

```
REDIS_URL=localhost:6379
```

```bash
go run ./cmd
```

The server listens on port `8000`. Each endpoint responds `200 Request allowed` or `429 Rate limit exceeded`; `400`/`500` on bad input or backend failure.

```bash
curl -H "x-real-ip: 1.2.3.4" http://localhost:8000/fixed-rate-limit
```
