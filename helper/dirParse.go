package helper

import (
	"strings"
	"os"
	"fmt"
	"path/filepath"
)

var filesAr []string

func DirParse(done <-chan bool, dirName string) <-chan string {
	filesAr = []string{}
	err := filepath.Walk(dirName, visit)

	if (err != nil) {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}

	return testNameGenerator(done, filesAr...)
}

func visit(path string, f os.FileInfo, err error) error {
	if (strings.Contains(path, "Functional") && strings.Contains(f.Name(), "Test.php")) {
		filesAr = append(filesAr, path)
	}

	return nil
}

func testNameGenerator(done <-chan bool, strings ...string) <-chan string {
	stringStream := make(chan string)

	go func() {
		defer close(stringStream)
		for _, testName := range strings {
			select {
			case <-done:
				return
			case stringStream <- testName:
			}
		}
	}()
	return stringStream

}
