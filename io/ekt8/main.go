package main

import (
	"./db"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().Unix())
		if rand.Int()%2 == 1 {
			go db.Query()
		} else {
			go db.Insert()
		}
	}
	time.Sleep(10 * time.Second)
}
