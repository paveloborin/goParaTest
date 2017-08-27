package helper

import (
	"strings"
	"strconv"
	"os/exec"
)

func ExeCmd(cmd, testName string, tokenNum int) ([]byte, error) {
	//strToken:=strconv.Itoa(tokenNum)
	//log.Println("command is ",cmd,testName,strToken )
	parts := strings.Fields(cmd + " " + testName + " tokenNum=" + strconv.Itoa(tokenNum))
	head := parts[0]
	parts = parts[1:len(parts)]

	return exec.Command(head, parts...).Output()

}