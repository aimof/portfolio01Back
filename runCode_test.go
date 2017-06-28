package main

import (
	"testing"
	"os"
)
/*
// pass this test
func TestMakeDir(t *testing.T) {
	dirName := "01234567"
	err := makeDir(dirName)
	if err != nil {
		t.Error("error occuer")
	}
	_, err = os.Stat(dirName)
	dirIsExist := err == nil
	if !dirIsExist {
		t.Error("dir is not exist")
	}
}
*/

/*
// Pass
func TestSetCode(t *testing.T) {
	id := "012345678"
	code := `print("Hello, World!")`
	err := setCode(id, code)
	if err != nil {
		t.Errorf("error occur: %v", err)
	}
	_, err = os.Stat(id + "/main.py")
	fileIsExist := err == nil
	if !fileIsExist {
		t.Error("file is not exist")
	}
}
*/