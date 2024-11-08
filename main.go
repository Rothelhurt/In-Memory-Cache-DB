package main

import (
	"fmt"
	"golang-ninja/hw1/cache"
	"time"
)

func main() {
	cache := cache.New()

	cache.Set("userId", 42, 5)
	userId := cache.Get("userId")

	fmt.Println(userId)

	time.Sleep(time.Second * 2)
	UserId2 := cache.Get("userId")
	fmt.Println(UserId2)

	time.Sleep(time.Second * 6)
	UserId3 := cache.Get("userId")
	fmt.Println(UserId3)

	cache.ShowDB()

	// userId, err := cache.Get("userId")
	// fmt.Println(cache)
	// userId := cache.Get("userId")

	// fmt.Println(userId)
	// nothing := cache.Get("Something")
	// fmt.Println(nothing)

	// fmt.Println(cache.Delete("userId"))
	// fmt.Println(cache.Delete("Something"))

	// fmt.Println(cache.Get("userId"))
}
