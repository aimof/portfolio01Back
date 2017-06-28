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
	errCleanDir := cleanDir(id)
	errCleanContainer := cleanContainer(id)
	err = combine3Errors(err, errCleanDir, errCleanContainer)
	fr := setResponse(err, result)
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

func setResponse(result Result, err error) (fResp FormatResp){
	if err != nil {
		fResp.Status = 1
		fResp.ErrMessage = err.Error()
		fResp.Result = result
		return fResp
	}
	fResp.Status = 0
	fResp.ErrMessage = ""
	fResp.Result = result
	return fResp
}



