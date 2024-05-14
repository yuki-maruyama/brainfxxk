package repl

import (
	"brainfxxk/lexar"
	"bufio"
	"fmt"
	"brainfxxk/parser"
	"brainfxxk/ast"
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
		p := parser.New(l)
		ast := p.ParseProgram()

		// for ast := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		// 	fmt.Printf("%+v\n", tok)
		// }
		printAST(ast, "")
	 }
}

func printAST(node ast.Node, indent string) {
	switch n := node.(type) {
	case *ast.Program:
		fmt.Println(indent + "Program")
		for _, child := range n.Body {
			printAST(child, indent+"  ")
		}
	case *ast.MoveRight:
		fmt.Println(indent + "MoveRight")
	case *ast.MoveLeft:
		fmt.Println(indent + "MoveLeft")
	case *ast.Increment:
		fmt.Println(indent + "Increment")
	case *ast.Decrement:
		fmt.Println(indent + "Decrement")
	case *ast.Output:
		fmt.Println(indent + "Output")
	case *ast.Input:
		fmt.Println(indent + "Input")
	case *ast.Loop:
		fmt.Println(indent + "Loop")
		for _, child := range n.Body {
			printAST(child, indent+"  ")
		}
	default:
		fmt.Println(indent + "Unknown node")
	}
}