local key = KEYS[1]
local limit = tonumber(ARGV[1])
local timestamp = tonumber(ARGV[2])
local id = ARGV[3]
local ttl = tonumber(ARGV[4])
local count = redis.call("zcard", key)

local allowed = count < limit

if allowed then 
    redis.call("zadd", key, timestamp, id)

    if count == 0 then
        redis.call("expire", key, ttl)
    end
end

return {allowed & 1 or 0, count + (allowed & 1 or 0)}