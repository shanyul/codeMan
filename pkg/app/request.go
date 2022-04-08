package app

import (
	"github.com/astaxie/beego/validation"
)

// MarkErrors logs error logs
func MarkErrors(errors []*validation.Error) string {
	var errString string
	for _, err := range errors {
		if errString == "" {
			errString = err.Message
		} else if err.Message != "" {
			errString = errString + ";" + err.Message
		}

	}
	return errString
}
