package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

type RedisCache struct {
	conn *redis.Client
}

func NewRedisCache() RedisCache {
	log.Println("Connection: Redis...")

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		panic("unable to connect to redis: " + err.Error())
	}

	log.Println("Connected: Redis")

	return RedisCache{
		conn: rdb,
	}
}

func (r RedisCache) InsertBloomFilter(data string) {
	ctx := context.Background()

	_, err := r.conn.Do(ctx, "BF.ADD", "bf_key", data).Bool()
	if err != nil {
		log.Println("error: unable to store data in bloom filter: " + err.Error())
		return
	}
}

func (r RedisCache) CheckBloom(data string) bool {
	ctx := context.Background()

	exists, err := r.conn.Do(ctx, "BF.EXISTS", "bf_key", data).Bool()
	if err != nil {
		panic(err)
	}

	if exists {
		return true
	} else {
		return false
	}
}
