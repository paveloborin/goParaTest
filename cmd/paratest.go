package main

import (
	"fmt"
	"sync"
	"time"
	"github.com/paveloborin/goParaTest/helper"
	"log"
)

var (
	tokens  map[int]bool
	results *sync.Map
)

func main() {
	var wg sync.WaitGroup

	results = new(sync.Map)
	numParallelProcess, dir, phpPath, phpUnitPath := helper.GetConsoleParams()

	initTokens(numParallelProcess)
	testNames := helper.DirParse(dir)

	wg.Add(len(testNames))
	for _, testName := range testNames {
		go worker(&wg, testName, phpPath, phpUnitPath)
	}

	wg.Wait()

	log.Println("Results:")
	results.Range(func(key interface{}, value interface{}) bool {
		log.Printf("%v %v", key, value)
		return true
	})

}

func initTokens(numParallelProcess int) {
	tokens = map[int]bool{}
	for i := 0; i < numParallelProcess; i++ {
		tokens[i] = true
	}
}

func getFreeToken() int {

	for true {
		for i, tokenAvailable := range tokens {
			if tokenAvailable {
				setTokenBusy(i)
				return i
			}
		}
		time.Sleep(100)

	}

	return 0
}

func setTokenBusy(tokenIndex int) {
	tokens[tokenIndex] = false
}

func setTokenFree(tokenIndex int) {
	tokens[tokenIndex] = true
}

func worker(wg *sync.WaitGroup, testName, phpPath, phpUnitPath string) {
	defer wg.Done()
	tokenIndex := getFreeToken()
	resultString := runTest(testName, tokenIndex, phpPath, phpUnitPath)
	setTokenFree(tokenIndex)

	results.Store(testName, resultString)
}

func runTest(testName string, tokenIndex int, phpPath, phpUnitPath string) string {

	log.Printf("running %v %v", testName, tokenIndex)

	_, err := helper.ExeCmd(fmt.Sprintf("%v %v", phpPath, phpUnitPath), testName, tokenIndex)

	result := fmt.Sprintf("+ %v", tokenIndex)
	if err != nil {
		result = fmt.Sprintf("- %v Error: %s", tokenIndex, err)
	}

	return result
}
