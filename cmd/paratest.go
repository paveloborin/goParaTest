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
	helper.InitTokens(config.ProcessesNum)

	start := time.Now()

	for result := range workersRun(done, config, &wg) {
		result.Print()
	}

	log.Printf("\nTest execution took %s\n", time.Since(start))
}

func workersRun(done <-chan bool, config helper.Config, wg *sync.WaitGroup) <-chan workers.ResultStruct {
	results := make(chan workers.ResultStruct)

	go func() {
		defer close(results)
		for testName := range helper.DirParse(done, config.TestDir) {
			wg.Add(1)
			go workers.Run(wg, testName, config.PhpPath, config.PhpUnitPath, config.PhpUnitConfiguration, results, done)
		}
		wg.Wait()
	}()

	return results;
}
