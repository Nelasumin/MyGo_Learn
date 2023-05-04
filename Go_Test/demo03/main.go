package main

import "math/rand"

/**
* 这是一个基准测试
*
**/

var ServerIndex [10]int

func InitServerIndex() {
	for i := 0; i < 10; i++ {
		ServerIndex[i] = i + 100
	}
}
func Select() int {
	return ServerIndex[rand.Intn(10)]
}
