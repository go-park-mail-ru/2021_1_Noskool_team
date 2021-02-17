package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

func main() {
	var err error
	// соединение
	con, err := redis.DialURL("redis://user:@localhost:6379/0")
	if err != nil {
		fmt.Println(err)
	}
	result, err := redis.String(con.Do("SET", "ok", "bad", "EX", 1))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	data, err := redis.String(con.Do("GET", "ok"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
	time.Sleep(1 * time.Second)
	data, err = redis.String(con.Do("GET", "ok"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)

	defer con.Close()
}
