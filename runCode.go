package main

import (
	"os/exec"
	"io/ioutil"
	"os"
	"github.com/pkg/errors"
	"log"
)

func runCode(id, code string) (result Result, err error) {
	err = setCode(id, code)
	result, err = runDocker(id)
	return result, err
}

func setCode(id, code string) (err error) {
	log.Println(id)
	cmdMkdir := "mkdir " + id
	cmd := exec.Command("bash", "-c", cmdMkdir)
	log.Println(cmd)
	err = cmd.Run()
	if err != nil {
		errors.Wrap(err, "setCode(): cannot make directory")
		log.Println(err)
		return err
	}
	err = ioutil.WriteFile("./" + id + "/main.py", []byte(code), os.ModePerm)
	if err != nil {
		errors.Wrap(err, "setCode(): cannot make code file")
		return err
	}
	return nil
}

