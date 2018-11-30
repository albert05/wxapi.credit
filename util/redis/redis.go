package redis

import (
	"github.com/go-redis/redis"
	"time"
)

type Redis struct {
	client *redis.Client
	LastErr error
}

var RD Redis

func Init(host, pass string, db int) {
	RD.client = redis.NewClient(&redis.Options{
		Addr:     host, //"localhost:6379"
		Password: pass,
		DB:       db,
	})

	_, err := RD.client.Ping().Result()
	if err != nil {
		panic(err)
	}
}

func (r *Redis) Get(key string) string {
	v, err := r.client.Get(key).Result()
	r.LastErr = err
	return v
}

func (r *Redis) Set(key string, value interface{}) bool {
	_, err := r.client.Set(key, value, 0).Result()
	r.LastErr = err
	return err == nil
}

func (r *Redis) SetNx(key string, value interface{}, expire int) bool {
	_, err := r.client.SetNX(key, value, time.Duration(expire) * time.Second).Result()
	r.LastErr = err
	return err == nil
}

func (r *Redis) Del(key string) bool {
	_, err := r.client.Del(key).Result()
	r.LastErr = err
	return err == nil
}

func (r *Redis) Pop(key string) string {
	v, err := r.client.LPop(key).Result()
	r.LastErr = err
	return v
}

func (r *Redis) Push(key string, value interface{}) bool {
	_, err := r.client.RPush(key, value).Result()
	r.LastErr = err
	return err == nil
}

func (r *Redis) Close() bool {
	r.LastErr = nil
	err := r.client.Close()
	return err == nil
}
