package main

import (
	"io/ioutil"
	"strings"
)

func dirParse(dirName string) []string {
	filesAr := []string{}
	files, _ := ioutil.ReadDir(dirName)
	for _, f := range files {
		if(!f.IsDir() && strings.Contains(f.Name(), "Test.php")){
			filesAr = append(filesAr, dirName+"/"+f.Name())
		}

	}
	return filesAr
}
