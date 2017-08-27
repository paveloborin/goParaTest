package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"
)

var tokens map[int]bool
var results map[string]string

func main() {
	var wg sync.WaitGroup

	numParallelProcess, dir := getConsoleParams()
	testNames := dirParse(dir)

	initTokens(numParallelProcess)
	results = map[string]string{}

	wg.Add(len(testNames))
	for _, testName := range testNames {
		go worker(&wg, testName)
	}

	wg.Wait()

	for k, v := range results {
		fmt.Println(k, v)
	}

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

func worker(wg *sync.WaitGroup, testName string) {
	defer wg.Done()
	tokenIndex := getFreeToken()
	resultString := runTest(testName, tokenIndex)
	setTokenFree(tokenIndex)

	results[testName] = resultString
}

func runTest(testName string, tokenIndex int) string {
	str := strconv.Itoa(tokenIndex)
	fmt.Println("running" + " " + testName + " " + str)

	time.Sleep(2 * time.Second)

	var wg sync.WaitGroup
	wg.Add(1)

	exe_cmd("php /Users/paveloborin/PhpstormProjects/symfony/phpunit-6.3.0.phar "+testName,
		&wg)
	wg.Wait()

	return "result" + " " + testName + " " + str
}

func exe_cmd(cmd string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("command is ", cmd)
	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]

	out, err := exec.Command(head, parts...).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("%s", out)
}
