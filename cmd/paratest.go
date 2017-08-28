package main

import (
	"sync"
	"github.com/paveloborin/goParaTest/helper"
	"github.com/paveloborin/goParaTest/workers"
	"github.com/paveloborin/goParaTest/logger"
	"time"
	"log"
)

var results *sync.Map

func main() {
	var wg sync.WaitGroup

	results = new(sync.Map)
	numParallelProcess, dir, phpPath, phpUnitPath, phpUnitConfiguration := helper.GetConsoleParams()

	helper.InitTokens(numParallelProcess)
	testNames := helper.DirParse(dir)

	start := time.Now()
	wg.Add(len(testNames))
	for _, testName := range testNames {
		go workers.Run(&wg, testName, phpPath, phpUnitPath, phpUnitConfiguration, results)
	}

	wg.Wait()

	elapsed := time.Since(start)
	log.Printf("\nTest execution took %s\n", elapsed)

	logger.PrintResult(results)
}
