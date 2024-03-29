package main

import (
	"os"
	"os/exec"
	"path/filepath"
)

type Emitter struct {
	code     string
	header   string
	fullPath string
}

func (e *Emitter) emit(code string) {
	e.code = e.code + code
}

func (e *Emitter) emitLine(code string) {
	e.code = e.code + code + "\n"
}

func (e *Emitter) headerLine(code string) {
	e.header = e.header + code + "\n"
}

func (e *Emitter) writeFile() {
	code := e.header + e.code
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	filepath.Dir(ex)
	// println(exPath)
	os.Mkdir("/home/dev/urdu/output", 0755)
	os.WriteFile(e.fullPath, []byte(code), 0755)
	exec.Command("gcc", e.fullPath, "-o", "/home/dev/urdu/output/out").Run()
	println("Compiled successfully!")
	println("Run with \n ./output/out")

}
