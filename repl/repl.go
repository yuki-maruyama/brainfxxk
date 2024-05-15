package repl

import (
	"github.com/yuki-maruyama/brainfxxk/interpreter"
	"bufio"
	"fmt"
	"io"
	"os"
)

func Start(in io.Reader, out io.Writer) {
	 scanner := bufio.NewScanner(in)
	 for {
		fmt.Printf(">> ")
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		interpreter.Run(line, os.Stdin, os.Stdout)
	 }
}