package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

func input(sc *bufio.Scanner, lines chan<- string, wg *sync.WaitGroup, wg2 *sync.WaitGroup) {

	for sc.Scan() {

		lines <- sc.Text()
	}
	wg.Done()
	wg2.Done()
}

func main() {

	file, err := os.Open("file.txt")

	if err != nil {

		fmt.Println("File does not exist!!!")
		return
	}

	defer file.Close()

	var wg = sync.WaitGroup{}
	var wg2 = sync.WaitGroup{}

	numberOfWorker := 2
	lines := make(chan string, 10)
	sc := bufio.NewScanner(file)
	var index = 0
	wg.Add(3)
	go func() {
		
		for line := range lines {

			fmt.Println("Dong " + strconv.Itoa(index) + " co gia tri la :" + line)
			index += 1
		}
		wg.Done()
	}()
	wg2.Add(2)
	for i := 0; i < numberOfWorker; i++ {

		go input(sc, lines, &wg, &wg2)
	}
	wg2.Wait()
	close(lines)
	wg.Wait()
}