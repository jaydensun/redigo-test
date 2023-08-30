local key=KEYS[1]
local val=ARGV[1]
return redis.call('SET', key, val)