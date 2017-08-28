package workers

import (
	"sync"
	"github.com/paveloborin/goParaTest/helper"
	"fmt"
	"time"
	"math/rand"
)

type ResultStruct struct {
	Result           string
	TestName         string
	TokenIndex       int
	ErrorDescription string
}

func Run(wg *sync.WaitGroup, testName, phpPath, phpUnitPath, phpUnitConfiguration string, results *sync.Map) {
	defer wg.Done()

	time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)

	tokenIndex := -1

	for tokenIndex < 0 {
		tokenIndex = helper.GetFreeToken()
		time.Sleep(300)
	}

	resultStruct := runTest(testName, tokenIndex, phpPath, phpUnitPath, phpUnitConfiguration)
	time.Sleep(500)
	helper.SetTokenFree(tokenIndex)

	results.Store(testName, resultStruct)
}

func runTest(testName string, tokenIndex int, phpPath, phpUnitPath, phpUnitConfiguration string) ResultStruct {

	//log.Printf("running %v %v", testName, tokenIndex)

	_, err := helper.ExeCmd(fmt.Sprintf("%v %v --configuration %v", phpPath, phpUnitPath, phpUnitConfiguration), testName, tokenIndex)

	//result := fmt.Sprintf("+ %v", tokenIndex)
	result := ResultStruct{Result: "+", TestName: testName, TokenIndex: tokenIndex}
	if err != nil {
		result.Result = "-"
		result.ErrorDescription = fmt.Sprintf("%s", err)
	}

	fmt.Print(result.Result)

	return result
}
