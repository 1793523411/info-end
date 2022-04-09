package utils

import (
	"log"
	"os"
	"path/filepath"
)

func GetCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	log.Println("path", res)
	return res
}
