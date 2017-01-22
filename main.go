package main

import (
	_ "ComplimentEngine/routers"
	"github.com/astaxie/beego"
        "gopkg.in/redis.v5"
)

func main() {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": "localhost:6379",
		},
	})
	client := redis.NewClient(&redis.Options{
		Addr:     "45.33.105.123",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	beego.Run()
}

