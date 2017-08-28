package logger

import (
	"sync"
	"github.com/paveloborin/goParaTest/workers"
	"fmt"
)

func PrintResult(results *sync.Map) {
	fmt.Println("Results:")
	results.Range(func(key interface{}, value interface{}) bool {
		resultStruct := value.(workers.ResultStruct)
		fmt.Print(resultStruct.Result)

		if (resultStruct.Result == "-") {
			fmt.Printf("\nTest: %v, TokenIndex: %v, Error: %v\n", resultStruct.TestName, resultStruct.TokenIndex, resultStruct.ErrorDescription)
		}

		return true
	})
	fmt.Println("\nEnd")
}
