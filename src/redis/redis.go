package redis

import (
	"encoding/json"
	"fmt"
	"go-redis-restApi/src/models"
	"log"

	"github.com/gomodule/redigo/redis"
)

// Redis connection pool instance
var pool = newPool()

// var pool2 *redis.Pool

/*
*	Redis Pool
 */

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}

}

/***********************************
*	Redis CRUD Functions via Redigo	*
*									*
 ***********************************/
func allArtists() {

	var artists Artist

	// Get a Redis pool
	c := pool.Get()
	defer c.Close()

	// keys, err := c.Do("KEYS", "artistName:*")
	// Handle Error

	values, err := redis.Values(c.Do("HGETALL", "artistName:*"))
	if err != nil {
		log.Fatal(err)
	}

	err = redis.ScanStruct(values, &artists)
	if err != nil {
		log.Fatal(err)
	}

	// for _, k := key.([]interface{}) {
	// 	var artist Artist

	// 	reply, err := c.Do("GET", k([]byte))
	// 	// Handle Error

	// 	if err := json.Unmarshal(reply.([])byte, &artist); err != nil {
	// 		panic(err)
	// 	}

	// 	artists = append(artists, artist)

	// }
	return artists
}

func findArtistConcert(name string) *Artist {

	var artist Artist

	// Get a Redis pool
	conn := pool.Get()
	defer conn.Close()

	// reply, err := c.Do("HMGET", "artistName:"+ name)
	// // Handle Error

	values, err := redis.Values(conn.Do("HGETALL", "artistName:"+name))
	if err != nil {
		log.Fatal(err)
	}

	err = redis.ScanStruct(values, &artist)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("GET OK")
	// if err = json.Unmarshal(reply.([]byte), &artist); err ! nil {
	// 	panic(err)
	// }

	return &artist

}

func createConcert(a models.Artist) {
	// Namespace - Artist:ArtistName - ex: Artist:U2

	artist = a.ArtistName

	// Get a Redis pool
	c := pool.Get()
	defer c.Close()

	b, err := json.Marshal(a)

	reply, err = c.Do("HMSET", "artistName:"+a.ArtistName, b)
	fmt.Println("GET", reply)
}
