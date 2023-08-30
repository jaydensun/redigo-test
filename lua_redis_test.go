package main

import (
    "fmt"
    "io/ioutil"
    "testing"
    "time"

    "github.com/gomodule/redigo/redis"
)

func CliPool() *redis.Pool {
    return &redis.Pool{
        MaxIdle:     10,
        MaxActive:   100,
        IdleTimeout: time.Duration(20) * time.Second,
        Dial: func() (redis.Conn, error) {
            c, err := redis.Dial("tcp", "xx.xx.xx.xx:16379", redis.DialPassword("xxxx"))
            if err != nil {
                return nil, err
            }
            return c, nil
        },
    }
}
func TestString(t *testing.T) {
    var data []byte
    var err error
    if data, err = ioutil.ReadFile("test.lua"); err != nil {
        fmt.Printf("read failed %v", err.Error())
        return
    }
    script := redis.NewScript(1, string(data))
    pool := CliPool()
    defer pool.Close()

    cli := pool.Get()
    script.Load(cli)

    arg := []string{"key", "value"}
    args := redis.Args{}.AddFlat(arg)
    var reply interface{}
    if reply, err = script.Do(cli, args...); err != nil {
        fmt.Printf("redis exec failed %v", err.Error())
        return
    }
    fmt.Printf("return %v", reply)
}
