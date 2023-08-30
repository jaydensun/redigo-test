local val = redis.call('zrangebyscore', KEYS[1], 0, ARGV[1], 'limit', 0, 1)
if(next(val) ~= nil) then
    redis.call('zremrangebyrank', KEYS[1], 0, #val - 1)
end
return val