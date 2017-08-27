package helper

import (
	"flag"
)

func GetConsoleParams() (int, string, string, string) {
	parProcessesNum := flag.Int("processes", 2, "number of parallel processes")
	dir := flag.String("dir", "./", "dir")
	phpPath:= flag.String("phpPath", "php", "path to php")
	phpUnitPath:= flag.String("phpUnitPath", "/Users/paveloborin/PhpstormProjects/symfony/phpunit-6.3.0.phar", "path to phpUnit")
	flag.Parse()

	return *parProcessesNum, *dir, *phpPath, *phpUnitPath
}
