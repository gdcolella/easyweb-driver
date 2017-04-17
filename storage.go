package main

import (
	"io/ioutil"
	"os"
)

var FILE = "./storage.txt"

type Contents struct {
	Content string
}

func read() Contents {
	file, e := ioutil.ReadFile(FILE)
	if e != nil {
		return Contents{Content: ""}
	}
	return Contents{Content: string(file)}
}

func write(s string) {
	ioutil.WriteFile(FILE, []byte(s), os.FileMode(0777))
}
