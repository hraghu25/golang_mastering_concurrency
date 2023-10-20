package main

import (
	"fmt"
	"sync"
)

type State struct {
	sync.Mutex
	count int
}

func main() {
	state := State{}
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			state.Lock()
			state.count++
			defer state.Unlock()
			wg.Done()
			fmt.Println("state.count = ", state.count)
		}(i)
	}

	wg.Wait()
}
