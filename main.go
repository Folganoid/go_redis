package main

import (
	"github.com/go-redis/redis"
	"strconv"
	"fmt"
)

var client = redis.NewClient(&redis.Options{
Addr:     "localhost:6379",
Password: "", // no password set
DB:       0,  // use default DB
})

func setVals() {

	for i := 1 ; i < 1000 ; i++ {
		fmt.Print(".")
		cnt := strconv.Itoa(i)
		err := client.Set("key" + cnt, "value" + cnt, 0).Err()
		if err != nil {
			panic(err)
		}
	}
}

func getVals() {
	for i := 1 ; i < 1000 ; i++ {
		cnt := strconv.Itoa(i)
		val, err := client.Get("key"+cnt).Result()
		if err != nil {
			panic(err)
		}
		fmt.Println("key"+cnt+": ", val)
	}
}


func main() {

	setVals()
	getVals()

}