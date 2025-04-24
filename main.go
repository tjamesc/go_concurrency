package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func producer(start, end, step int, inCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := start; i <= end; i += step {
		inCh <- i
		sleepTime := rand.Intn(1500)
		time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	}
}

func consumer(inCh <-chan int, outCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range inCh {
		squared := num * num
		sleepTime := rand.Intn(3000)
		time.Sleep(time.Duration(sleepTime) * time.Millisecond)
		outCh <- squared
	}
}

func finalFilter(outCh <-chan int) {
	lastPrinted := -1
	for squared := range outCh {
		if lastPrinted == -1 || squared > lastPrinted {
			fmt.Println(squared)
			lastPrinted = squared
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	inCh := make(chan int, 5)
	outCh := make(chan int, 5)

	var wgProducers sync.WaitGroup
	wgProducers.Add(2)
	go producer(1, 29, 2, inCh, &wgProducers)
	go producer(2, 30, 2, inCh, &wgProducers)

	go func() {
		wgProducers.Wait()
		close(inCh)
	}()

	var wgConsumers sync.WaitGroup
	wgConsumers.Add(2)
	go consumer(inCh, outCh, &wgConsumers)
	go consumer(inCh, outCh, &wgConsumers)

	go func() {
		wgConsumers.Wait()
		close(outCh)
	}()

	finalFilter(outCh)
}
