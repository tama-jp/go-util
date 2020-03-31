package go_util

import (
	"testing" // テストで使える関数・構造体が用意されているパッケージをimport
)

func TestGetExecPath(t *testing.T) {
	path := GetExecPath()
	if len(path) == 0 {
		t.Errorf("No Data")
	} else {
		t.Log("TestGetExecPath", path)
	}
}

func TestGetCurrentPath(t *testing.T) {
	path := GetCurrentPath()
	if len(path) == 0 {
		t.Errorf("No Data")
	} else {
		t.Log("TestGetCurrentPath", path)
	}
}

func TestGetGetCallInfoPath(t *testing.T) {
	path, line, funcName := GetCallInfoPath(1)
	if len(path) == 0 {
		t.Errorf("No Data")
	} else {
		t.Log("TestGetCurrentPath:", path, "funcName:", funcName, "line:", line)

	}
}

func TestGetGetTime(t *testing.T) {
	timeString := GetTime()
	if len(timeString) == 0 {
		t.Errorf("No Data")
	} else {
		t.Log("timeString:", timeString)

	}
}
