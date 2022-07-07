package main

import (

	"log"
	"time"
	"sync"
)

var wg = sync.WaitGroup{}

func chanRoutine() {
    log.Print("hello 1")
	wg.Add(1)
    go func() {
        time.Sleep(1 * time.Second)
        log.Print("hello 3")
		wg.Done()
    }()
	// wg.Wait()
    log.Print("hello 2")
	wg.Wait()
}


func main() {

	chanRoutine()
}