package main

import (
	"fmt"
	"sync"
)

type Single struct {
}

var instance *Single
var once sync.Once

func GetInstance() *Single {
	once.Do(func() {
		instance = &Single{}
	})

	return instance
}

func main() {
	single := GetInstance()
	single2 := GetInstance()
	fmt.Println(single == single2)
}
