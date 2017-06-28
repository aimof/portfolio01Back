package main

import "github.com/pkg/errors"

func combine3Errors(err1, err2, err3 error)(combinedErr error) {
	combinedErr = errors.New(err1.Error() + err2.Error() + err3.Error())
	return combinedErr
}
