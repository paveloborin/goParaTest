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

func (m ResultStruct) Print() {
	fmt.Print(m.Result)

	if (m.Result == "-") {
		fmt.Printf("\nTest: %v, TokenIndex: %v, Error: %v\n", m.TestName, m.TokenIndex, m.ErrorDescription)
	}
}

func Run(wg *sync.WaitGroup, testName string, phpPath, phpUnitPath, phpUnitConfiguration string, results chan<- ResultStruct, done <-chan bool) {
	defer wg.Done()

	select {
	case <-done:
		return
	default:
		time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)

		//TODO канал?
		tokenIndex := -1
		for tokenIndex < 0 {
			tokenIndex = helper.GetFreeToken()
			time.Sleep(300)
		}

		resultStruct := runTest(testName, tokenIndex, phpPath, phpUnitPath, phpUnitConfiguration)
		time.Sleep(500)
		helper.SetTokenFree(tokenIndex)

		results <- resultStruct

	}

}

func runTest(testName string, tokenIndex int, phpPath, phpUnitPath, phpUnitConfiguration string) ResultStruct {

	//res, err := helper.ExeCmd(fmt.Sprintf("%v %v --configuration %v", phpPath, phpUnitPath, phpUnitConfiguration), testName, tokenIndex)

	result := ResultStruct{Result: "+", TestName: testName, TokenIndex: tokenIndex}
	//if err != nil {
	//	result.Result = "-"
	//	result.ErrorDescription = fmt.Sprintf("%s %s", err, res)
	//}

	fmt.Print(result.Result)

	return result
}
