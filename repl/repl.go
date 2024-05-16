package repl

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/yuki-maruyama/brainfxxk/interpreter"
)

func Start(ctx context.Context, in io.Reader, out io.Writer) {
	 scanner := bufio.NewScanner(in)
	 for {
		fmt.Printf(">> ")
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		config := &interpreter.Config {
			MemorySize: 1024,
			MaxStep: 100000,

			Reader: os.Stdin,
			Writer: os.Stdout,
		}

		err := interpreter.Run(ctx, line, config); if err != nil {
			fmt.Println("Error: ", err)
		}
	 }
}