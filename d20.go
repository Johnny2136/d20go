package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var onlyOnce sync.Once

// prepare the dice
var d20 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13,
	14, 15, 16, 17, 18, 19, 20}

func rollD20() int {

	onlyOnce.Do(func() {
		rand.Seed(time.Now().UnixNano()) // only run once
	})

	return d20[rand.Intn(len(d20))]
}

func main() {
	d20a := rollD20()

	fmt.Println("D20  = ", d20a)
	if d20a == 1 {
		fmt.Println("Critical fail!!!")
	}
	if d20a == 20 {
		fmt.Println("Critical success!!!")
	}
}
