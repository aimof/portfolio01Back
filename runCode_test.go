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
	isExist := err == nil
	if !isExist {
		t.Error("dir is not exist")
	}
}
*/

func TestSetCode(t *testing.T) {
	id := "012345678"
	code := `print("Hello, World!")`
	err := setCode(id, code)
	if err != nil {
		
	}
}
