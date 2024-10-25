# In Memory Cache DB
In Memory Cache DB module on Go. 

# Installation
Run command:
```
go get github.com/Rothelhurt/In-Memory-Cache-DB
```

# Usage exemples
```
package main

import (
	"fmt"
	"cache"
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
```
