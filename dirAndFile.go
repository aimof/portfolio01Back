package main

import (
	"io/ioutil"
	"os"
	"github.com/pkg/errors"
	"log"
	"os/exec"
)

func runCode(id, code string) (result Result, err error) {
	err = makeDir(id)
	err = setCode(id, code)
	result, err = runDocker(id)
	return result, err
}

func makeDir(id string) (err error) {
	log.Println(id)
	err = os.Mkdir(id, 0777)
	if err != nil {
		errors.Wrap(err, "setCode(): cannot make directory")
		log.Println(err)
		return err
	}
	return nil
}

func setCode(id, code string) (err error) {
	err = ioutil.WriteFile("./" + id + "/main.py", []byte(code), os.ModePerm)
	if err != nil {
		errors.Wrap(err, "setCode(): cannot make code file")
		return err
	}
	return nil
}

func cleanDir(id string) (err error) {
	cmd := exec.Command("sh", "-c", "docker rm " + id)
	err = cmd.Run()
	if err != nil {
		errors.Wrap(err, "clean(): cannot remove docker image")
		return err
	}
	return nil
}
