package logger

import (
	"log"
	"sync"
)

func PrintResult(results *sync.Map) {
	log.Println("Results:")
	results.Range(func(key interface{}, value interface{}) bool {
		log.Printf("%v %v", key, value)
		return true
	})
}
