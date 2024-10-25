package main

import (
	"fmt"
	"golang-ninja/hw1/cache"
)

func main() {
	cache := cache.New()

	cache.Set("userId", 42)
	fmt.Println(cache)
	userId := cache.Get("userId")

	fmt.Println(userId)
	nothing := cache.Get("Something")
	fmt.Println(nothing)

	fmt.Println(cache.Delete("userId"))
	fmt.Println(cache.Delete("Something"))

	fmt.Println(cache.Get("userId"))
}
