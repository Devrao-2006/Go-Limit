local current = tonumber(redis.call('GET', KEYS[1]) or "0")
local limit = tonumber(ARGV[1])
local window = tonumber(ARGV[2])

if current > limit then
    return current
end

current = redis.call('INCR', KEYS[1])
if current == 1 then
    redis.call('EXPIRE', KEYS[1], window)
end
return current