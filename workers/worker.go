package workers

import (
	"sync"
	"log"
	"github.com/paveloborin/goParaTest/helper"
	"fmt"
)

func Run(wg *sync.WaitGroup, testName, phpPath, phpUnitPath string, results *sync.Map) {
	defer wg.Done()

	tokenIndex := helper.GetFreeToken()
	resultString := runTest(testName, tokenIndex, phpPath, phpUnitPath)
	helper.SetTokenFree(tokenIndex)

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
