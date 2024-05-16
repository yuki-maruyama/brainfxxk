package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"

	"github.com/yuki-maruyama/brainfxxk/ast"
	"github.com/yuki-maruyama/brainfxxk/interpreter"
	"github.com/yuki-maruyama/brainfxxk/lexar"
	"github.com/yuki-maruyama/brainfxxk/parser"
	"github.com/yuki-maruyama/brainfxxk/repl"
)

func main(){
	flag.Parse()
	args := flag.Args()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	if len(args) != 0 {
		file, err := os.ReadFile(args[0])
		if err != nil {
			fmt.Print("Failed to open file")
		}
		script := string(file)
		config := &interpreter.Config {
			MemorySize: 65536,
			MaxStep: 100000,

			Reader: os.Stdin,
			Writer: os.Stdout,
		}
		err = interpreter.Run(ctx, script, config)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		fmt.Printf("brainfxxk repl\n")
		repl.Start(ctx, os.Stdin, os.Stdout)
	}
}