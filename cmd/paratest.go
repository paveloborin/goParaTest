package main

import (
	"sync"
	"github.com/paveloborin/goParaTest/helper"
	"github.com/paveloborin/goParaTest/workers"
	"time"
	"log"
)

func main() {
	var wg sync.WaitGroup
	done := make(chan bool)
	defer close(done)

	config := helper.GetConsoleParams()

	start := time.Now()

	for result := range workersRun(done, config, &wg) {
		result.Print()
	}

	log.Printf("\nTest execution took %s\n", time.Since(start))
}

func workersRun(done <-chan bool, config helper.Config, wg *sync.WaitGroup) <-chan workers.ResultStruct {
	results := make(chan workers.ResultStruct)
	tokens := make(chan int, config.ProcessesNum)

	for i := 0; i < config.ProcessesNum; i++ {
		tokens <- i + 1
		log.Printf("Init token %v", i+1)
	}

	go func() {
		defer close(results)
		defer close(tokens)

		for testName := range helper.DirParse(done, config.TestDir) {
			wg.Add(1)
			go workers.Run(done, wg, testName, config, results, tokens)
		}
		wg.Wait()
	}()

	return results;
}
