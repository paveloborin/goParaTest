package helper

import (
	"strings"
	"os"
	"fmt"
	"path/filepath"
)

var filesAr []string

func DirParse(dirName string) []string {
	filesAr = []string{}
	err := filepath.Walk(dirName, visit)

	if (err != nil) {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}

	return filesAr
}

func visit(path string, f os.FileInfo, err error) error {
	if (strings.Contains(path, "Functional") && strings.Contains(f.Name(), "Test.php")) {
		filesAr = append(filesAr, path)
	}

	return nil
}
