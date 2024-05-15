package main

import (
	"github.com/yuki-maruyama/brainfxxk/interpreter"
	"github.com/yuki-maruyama/brainfxxk/repl"
	"flag"
	"fmt"
	"os"
)

func main(){
	flag.Parse()
	args := flag.Args()

	if len(args) != 0 {
		file, err := os.ReadFile(args[0])
		if err != nil {
			fmt.Print("Failed to open file")
		}
		script := string(file)
		interpreter.Run(script, os.Stdin, os.Stdout)
	} else {
		fmt.Printf("brainfxxk repl\n")
		repl.Start(os.Stdin, os.Stdout)
	}
}