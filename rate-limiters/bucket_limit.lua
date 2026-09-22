local key = KEYS[1]
local limit = tonumber(ARGV[1])
local time_per_token = tonumber(ARGV[2]) 
local now = tonumber(ARGV[3])
local max_refill_time_seconds = tonumber(ARGV[4]) 

local bucket = redis.call('HMGET', key, 'count', 'last_refill')
local count = tonumber(bucket[1])
local last_refill = tonumber(bucket[2])

if count == nil then
    count = limit
    last_refill = now
end

local time_passed = now - last_refill
local tokens_to_add = math.floor(time_passed / time_per_token)

if tokens_to_add > 0 then
    count = math.min(limit, count + tokens_to_add)
    last_refill = last_refill + (tokens_to_add * time_per_token)
end

local allowed = count >= 1
if allowed then
    count = count - 1
end

redis.call('HMSET', key, 'count', count, 'last_refill', last_refill)

redis.call('EXPIRE', key, max_refill_time_seconds)

return {allowed and 1 or 0, count}
