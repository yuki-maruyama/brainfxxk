package main 

import (
	"fmt"
	"os"
	"brainfxxk/repl"
)

func main(){
	fmt.Printf("brainfxxk repl\n")
	repl.Start(os.Stdin, os.Stdout)
}