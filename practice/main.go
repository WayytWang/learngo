package main

import (
	"fmt"
)

type Cache struct {
	PDate int
}

func (cache Cache) SetPDate() {
	cache.PDate = 5
}


func main() {
	cache1 := &Cache{PDate:3}
	cache2 := &Cache{PDate:3}
	fmt.Println(cache1 == cache2)

	cache1.SetPDate()
	fmt.Println(cache1.PDate)
}