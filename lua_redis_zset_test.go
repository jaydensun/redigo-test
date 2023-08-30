package main

import (
    "fmt"
    "github.com/gomodule/redigo/redis"
    "io/ioutil"
    "testing"
)

func TestZset(t *testing.T) {
    var data []byte
    var err error
    if data, err = ioutil.ReadFile("zset.lua"); err != nil {
        fmt.Printf("read failed %v", err.Error())
        return
    }
    script := redis.NewScript(1, string(data))
    pool := CliPool()
    defer func(pool *redis.Pool) {
        err := pool.Close()
        if err != nil {
            fmt.Println(err)
        }
    }(pool)

    cli := pool.Get()
    //script.Load(cli)

    args := redis.Args{}.AddFlat([]any{"myzset", 2})
    items, err := redis.ByteSlices(script.Do(cli, args...))
    if err != nil {
        fmt.Println(err)
        return
    }
    if len(items) != 1 {
        err = redis.ErrNil
        fmt.Println(err)
        return
    }
    fmt.Printf("return %v", string(items[0]))
    fmt.Println()
}
