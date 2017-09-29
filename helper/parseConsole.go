package helper

import (
	"flag"
)


func GetConsoleParams() Config {
	config := LoadConfiguration()

	parProcessesNum := flag.Int("processes", config.ProcessesNum, "number of parallel processes")
	dir := flag.String("dir", config.TestDir, "dir")
	phpPath := flag.String("phpPath", config.PhpPath, "path to php")
	phpUnitPath := flag.String("phpUnitPath", config.PhpUnitPath, "path to phpUnit")
	phpUnitConfiguration := flag.String("phpUnitConfiguration", config.PhpUnitConfiguration, "path to phpUnit config")
	flag.Parse()

	return Config{ProcessesNum: *parProcessesNum, TestDir: *dir, PhpPath: *phpPath, PhpUnitPath: *phpUnitPath, PhpUnitConfiguration: *phpUnitConfiguration}
}
