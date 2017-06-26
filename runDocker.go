package main

import (
	"os/exec"
	"github.com/pkg/errors"
	"io/ioutil"
	"path/filepath"
)

func runDocker(id string) (result Result, err error) {
	absPath, err := filepath.Abs(".")
	if err != nil {
		errors.Wrap(err, "runDocker(): cannot find abs path")
	}
	cmd := exec.Command("sh", "-c", "docker create -i --net none --cpuset-cpus 0 --memory 512m --memory-swap 512m --ulimit nproc=10:10 --ulimit fsize=10000000 --name " + id + "-v " + absPath + id + ":/home/ python:3.6.1-alpine (python3 /home/main.py 1> stdOut.txt 2>stdErr.txt)")
	err = cmd.Run()
	if err != nil {
		err = errors.Wrap(err, "runDocker(): cannot create docker container")
		return initializeResult(), err
	}
	cmd = exec.Command("sh", "-c", "docker start " + id)
	err = cmd.Run()
	bStdOut, err := ioutil.ReadFile(`./` + id + `/stdOut.txt`)
	if err != nil {
		errors.Wrap(err, "runDocker(): cannot read stdOut")
		return initializeResult(), err
	}
	bStdErr, err := ioutil.ReadFile(`./` + id + `/stdErr.txt`)
	if err != nil {
		errors.Wrap(err, "runDocker(): cannot read stdErr")
		return initializeResult(), err
	}
	result = Result {
		StdOut: string(bStdOut),
		StdErr: string(bStdErr),
	}
	return result, nil
}
