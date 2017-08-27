package helper

import (
	"strings"
	"strconv"
	"os/exec"
)

func ExeCmd(cmd, testName string, tokenNum int) ([]byte, error) {
	parts := strings.Fields(cmd + " " + testName + " tokenNum=" + strconv.Itoa(tokenNum))
	head := parts[0]
	parts = parts[1:len(parts)]

	return exec.Command(head, parts...).Output()

}
