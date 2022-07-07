package main

import (

	"fmt"
	"sync"
	"math/rand"
)

var X map[string]string = map[string]string{}

func randomString() string {

	var ans string = ""
	for i := 0; i < 20; i++ {

		ans += string(rand.Intn(26) + 65)
	}
	return ans
}

var m = sync.RWMutex{}
var wg = sync.WaitGroup{}

func addX() {

	for i := 0; i < 1000; i++ {

		X[randomString()] = randomString()
	}
	m.Unlock()
	wg.Done()
}

func main() {

	wg.Add(3)
	m.Lock()
	go addX()
	m.Lock()
	go addX()
	m.Lock()
	go addX()
	wg.Wait()
	fmt.Println(len(X))
}
