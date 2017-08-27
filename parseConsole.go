package main

import (
	"flag"
)

func getConsoleParams() (int, string) {
	parProcessesNum := flag.Int("processes", 1, "number of parallel processes")
	dir := flag.String("dir", "./", "dir")
	flag.Parse()

	return *parProcessesNum, *dir
}
