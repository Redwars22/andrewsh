package main

import (
	"os"
	"runtime"
)

func changeDirectoryWindows(dir string) error {
	return os.Chdir(dir)
}

func changeDirectory(dir string) error {
	if runtime.GOOS == "windows" {
		return changeDirectoryWindows(dir)
	}
	return os.Chdir(dir)
}