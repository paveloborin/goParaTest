package helper

import (
	"time"
)

var tokens map[int]bool

func InitTokens(numParallelProcess int) {
	tokens = map[int]bool{}
	for i := 0; i < numParallelProcess; i++ {
		tokens[i] = true
	}
}

func setTokenBusy(tokenIndex int) {
	tokens[tokenIndex] = false
}

func SetTokenFree(tokenIndex int) {
	tokens[tokenIndex] = true
}

func GetFreeToken() int {

	for true {
		for i, tokenAvailable := range tokens {
			if tokenAvailable {
				setTokenBusy(i)
				return i
			}
		}
		time.Sleep(200)

	}

	return 0
}
