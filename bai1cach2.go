package main

import (

	"log"
	"time"
)

func chanRoutine() {

	ch := make(chan string)
    log.Print("hello 1")
    go func() {
        time.Sleep(1 * time.Second)
        // log.Print("hello 3")
		ch <- "hello 3"
		close(ch)
    }()
	// log.Print(<- ch)
    log.Print("hello 2")
	log.Print(<- ch)
}


func main() {

	chanRoutine()
}