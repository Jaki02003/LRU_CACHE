package main

import (
	"fmt"
	"github.com/Jaki02003/LRU_CACHE/cache"
)

func main() {
	newInstance := cache.NewLRUCache(3)
	newInstance.Put(10, "a")
	newInstance.Put("Jak", 10)
	newInstance.Put("sakib", 75)
	newInstance.Put("02", "Mahdi")
	fmt.Println(newInstance.Get(10))
	fmt.Println(newInstance.Get("Jak"))
	fmt.Println(newInstance.Get("sakib"))
	fmt.Println(newInstance.Get("02"))
}
