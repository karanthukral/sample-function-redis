package main

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

func Main(args map[string]interface{}) map[string]interface{} {
	name, ok := args["name"].(string)
	if !ok {
		name = "stranger"
	}
	msg := make(map[string]interface{})
	msg["body"] = "Hello " + name + "!"

	fmt.Println("Parsing redis url....")
	opts, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		panic(err)
	}

	fmt.Println("Creating redis client...")
	client := redis.NewClient(opts)

	fmt.Println("Pinging....")
	pong, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}

	msg["body"] = fmt.Sprintf("%s. Ping result: %s", msg["body"], pong)

	return msg
}
