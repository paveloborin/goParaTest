package main

import (
	"sync"
	"log"

	"github.com/paveloborin/goParaTest/helper"
	"github.com/paveloborin/goParaTest/workers"
)

var results *sync.Map

func main() {
	var wg sync.WaitGroup

	results = new(sync.Map)
	numParallelProcess, dir, phpPath, phpUnitPath := helper.GetConsoleParams()

	helper.InitTokens(numParallelProcess)
	testNames := helper.DirParse(dir)

	wg.Add(len(testNames))
	for _, testName := range testNames {
		go workers.Worker(&wg, testName, phpPath, phpUnitPath, results)
	}

	wg.Wait()

	log.Println("Results:")
	results.Range(func(key interface{}, value interface{}) bool {
		log.Printf("%v %v", key, value)
		return true
	})
}
