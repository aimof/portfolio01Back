package main

import (
//	"github.com/labstack/echo"
//	"net/http"
	"time"
	"strconv"
	"os/exec"
	"github.com/pkg/errors"
	"fmt"
)

func main() {
	code := `print("Hello, World!")`
	now := strconv.Itoa(int(time.Now().UnixNano()))
	id := now[len(now)-8:len(now)-1]
	result, err := runCode(id, code)
	errClean := clean(id)
	var fResp FormatResp
	if errClean != nil {
		fResp.Status = 2
		fResp.ErrMessage = errClean.Error()
		fResp.Result = result
	} else if err != nil {
		fResp.Status = 1
		fResp.ErrMessage = err.Error()
		fResp.Result = result
	}
	fResp.Status = 0
	fResp.ErrMessage = ""
	fResp.Result = result
	fr := &fResp
	fmt.Println(fr)
}

/*
func codeRunner(c echo.Context) error {
	code := c.FormValue("code")
	now := strconv.Itoa(int(time.Now().UnixNano()))
	id := now[len(now)-8:len(now)-1]
	result, err := runCode(id, code)
	errClean := clean(id)
	var fResp *FormatResp
	if errClean != nil {
		fResp.Status = 2
		fResp.ErrMessage = errClean.Error()
		fResp.Result = result
	}
	if err != nil {
		fResp.Status = 1
		fResp.ErrMessage = err.Error()
		fResp.Result = result
	}
	fResp.Status = 0
	fResp.ErrMessage = ""
	fResp.Result = result
	return c.JSON(http.StatusOK, fResp)
}
*/
func clean(id string) (err error) {
	cmd := exec.Command("sh", "-c", "docker rm " + id)
	err = cmd.Run()
	if err != nil {
		errors.Wrap(err, "clean(): cannot remove docker image")
	}
	cmd = exec.Command("sh", "-c", "rm -Rf ./" + id)
	err2 := cmd.Run()
	if err2 != nil {
		err = errors.Wrap(err, "clean(): cannnot remove dir:" + err2.Error())
		return err
	}
	return nil
}