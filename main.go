package main

import (
	"flag"
	"fmt"
	"os"
)

type Source struct {
	source         string
	curChar        string
	curPos         int
	symbols        set
	labelsDeclared set
	labelsGotoed   set
}

func readSource(filename string) string {

	b, err := os.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}
	txt := string(b)
	return txt
	//check how many new lines there are

}

func main() {
	var file string
	flag.StringVar(&file, "f", "", "a file to compile")
	//verbose := flag.Bool("v", false, "verbose mode")
	run := flag.Bool("r", false, "run the program after compiling")
	flag.Parse()

	if len(file) == 0 {
		panic("File missing. Usage: ./compile <filename>.urdu")
	}
	println("Compiling file: ", file)
	if file[len(file)-5:] != ".urdu" {
		panic("File must be of type .urdu")
	}

	codeString := readSource(file)
	source := Source{source: codeString + "\n", curChar: "", curPos: -1, symbols: *NewSet(), labelsDeclared: *NewSet(), labelsGotoed: *NewSet()}
	emitter := Emitter{fullPath: "output/out.c", header: "", code: "", run: *run}
	parser := Parser{source: source, curToken: Token{}, peekToken: Token{}, emitter: emitter}
	do_parsing(parser)

}
