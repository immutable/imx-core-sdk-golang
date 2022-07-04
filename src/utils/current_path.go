package utils

import (
	"errors"
	"path/filepath"
	"runtime"
)

// CurrentFilename gets the full file path of the caller of this function.
func CurrentFilename() (string, error) {
	_, filename, _, ok := runtime.Caller(1) // Caller 1 will get this function callers path.
	if !ok {
		return "", errors.New("unable to get the current filename")
	}
	return filename, nil
}

// CurrentDirname gets the full directory path of the caller of this function.
func CurrentDirname() (string, error) {
	_, filename, _, ok := runtime.Caller(1) // Caller 1 will get this function callers path.
	if !ok {
		return "", errors.New("unable to get the current filename")
	}
	return filepath.Dir(filename), nil
}
