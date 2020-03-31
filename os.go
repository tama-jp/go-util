package go_util

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func GetExecPath() string {
	var retString = ""
	exe, err := os.Executable()
	if err != nil {
		fmt.Println("GetExecPath", err)
	} else {
		retString = filepath.Dir(exe)
	}
	return retString
}

func GetCurrentPath() string {
	var retString = ""
	current, err := os.Getwd()
	if err != nil {
		fmt.Println("GetCurrentPath", err)
	} else {
		//retString = filepath.Dir(current)
		retString = current
	}
	return retString
}

func GetCallInfoPath(skip int) (file string, line int, functionName string, ) {

	var funcName = ""
	pc, file, line, ok := runtime.Caller(skip)

	if ok == false {
		file = ""
		line = 0
	} else {
		file = strings.Replace(file, GetCurrentPath()+"/", "./", 1) // <0で無制限

		funcNamePath := runtime.FuncForPC(pc).Name()

		slice := strings.Split(funcNamePath, "/")

		funcName = slice[len(slice)-1]
	}

	return file, line, funcName
}

func GetTime() string {
	t := time.Now()
	const layout = "2006-01-02 15:04:05"
	return t.Format(layout)
}

func GetTimeYYYYMMDD() string {
	t := time.Now()
	const layout = "20060102"
	return t.Format(layout)
}