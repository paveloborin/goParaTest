package helper

import (
	"sync"
	"log"
)

var tokens *sync.Map

func InitTokens(numParallelProcess int) {
	tokens = new(sync.Map)
	for i := 0; i < numParallelProcess; i++ {
		log.Printf("Init token %v", i)
		tokens.Store(i, true)
	}
}

func setTokenBusy(tokenIndex int) {
	tokens.Store(tokenIndex, false)
}

func SetTokenFree(tokenIndex int) {
	tokens.Store(tokenIndex, true)
}

func GetFreeToken() int {

	index := -1

	tokens.Range(func(key interface{}, value interface{}) bool {
		if (value.(bool) == true) {
			setTokenBusy(key.(int))
			index = key.(int)

			return false
		}
		return true
	})

	return index
}
