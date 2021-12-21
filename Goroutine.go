package main

import (
	"fmt"
	"sync"
)

var counter int = 0

func main() {
	var wg sync.WaitGroup
	var mutex sync.RWMutex

	var i int
	for i = 0; i < 10; i++ {
		wg.Add(1)
		go asyncIncrementJob(&wg, &mutex)
	}

	wg.Wait()

}

func asyncIncrementJob(wg *sync.WaitGroup, mu *sync.RWMutex) {
	mu.Lock()
	counter = counter + 1
	fmt.Println("Counter incremented to ", counter)
	mu.Unlock()
	wg.Done()
}
