package utils

import (
	"os"
	"runtime"
)

func GetCurrentFunctionName() string {
	counter, _, _, success := runtime.Caller(1)

	if !success {
		println("functionName: runtime.Caller: failed")
		os.Exit(1)
	}

	return runtime.FuncForPC(counter).Name()
}
