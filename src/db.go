package main

import (
	"fmt"
	"encoding/json"
	"strconv"

	"github.com/gomodule/redigo/redis"
)


type PoolConn struct {
	pool *redis.Pool
}

var connection *redis.Conn

/*
*	Redis Connection
*/
func redisConnection() redis.Conn{

	pool = &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		}
	}
}

func redisConnect() redis.Conn {
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	return conn
}

/*
*	CRUD Functions
*/

func allArtists() Artists {

	
}

func (c *PoolConn) allArtists() {
	
}