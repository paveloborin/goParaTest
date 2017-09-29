package helper

import (
	"sync"
)

func FanIn(done <-chan bool, channels ... <-chan int) <-chan int {
	var wg sync.WaitGroup
	multiplexedStream := make(chan int)

	multiplex := func(c <-chan int) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case multiplexedStream <- i:
			}
		}
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}

	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()

	return multiplexedStream
}
