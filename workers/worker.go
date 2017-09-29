package workers

import (
	"sync"
	"fmt"
	"github.com/paveloborin/goParaTest/helper"
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

func Run(done <-chan bool, wg *sync.WaitGroup, testName string, config helper.Config, results chan<- ResultStruct, tokens chan int) {
	defer wg.Done()

	select {
	case <-done:
		return
	case tokenIndex := <-tokens:
		results <- runTest(testName, tokenIndex, config)
		tokens <- tokenIndex

	}

}

func runTest(testName string, tokenIndex int, config helper.Config) ResultStruct {

	res, err := helper.ExeCmd(fmt.Sprintf("%v %v --configuration %v", config.PhpPath, config.PhpUnitPath, config.PhpUnitConfiguration), testName, tokenIndex)

	result := ResultStruct{Result: "+", TestName: testName, TokenIndex: tokenIndex}
	//time.Sleep(10 * time.Millisecond)
	if err != nil {
		result.Result = "-"
		result.ErrorDescription = fmt.Sprintf("%s %s", err, res)
	}

	return result
}
