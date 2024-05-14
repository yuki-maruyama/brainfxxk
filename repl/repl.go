package repl

import (
	"brainfxxk/lexar"
	"bufio"
	"fmt"
	"brainfxxk/token"
	"io"
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
		l := lexar.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	 }
}