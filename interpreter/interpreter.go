package interpreter

import (
	"brainfxxk/ast"
	"brainfxxk/parser"
	"fmt"
	"io"
	"os"
)

const MAX_MEM = 1024

type Interpreter struct {
	Program *ast.Program
	Reader  *os.File
	Writer  *os.File

	Memory  []byte
	Cursor  int
}

func Run(script string, reader *os.File, writer *os.File) error{
	p, err := parser.Parse(script)
	if err != nil {
		return err
	}
	
	return NewInterpreter(p, reader, writer).Run()
}

func NewInterpreter(p *ast.Program, r *os.File, w *os.File) *Interpreter{
	return &Interpreter{
		Program: p,
		Reader: r,
		Writer: w,
		Memory: make([]byte, MAX_MEM),
		Cursor: 0,
	}
}

func (i *Interpreter) Run() error{
	err := i.runExpressions(i.Program.Body)
	return err
} 

func (i *Interpreter) runExpressions(exprs []ast.Node) error{
	for _, expr := range exprs {
		err := i.runExpression(expr)
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *Interpreter) runExpression(expr ast.Node) error{
	switch e := expr.(type) {
	case *ast.Increment:
		i.Memory[i.Cursor]++
	case *ast.Decrement:
		i.Memory[i.Cursor]--
	case *ast.MoveRight:
		i.Cursor++
		if i.Cursor < 0 || i.Cursor > MAX_MEM{
			return fmt.Errorf("cursor range exceeds valid bounds")
		}
	case *ast.MoveLeft:
		i.Cursor--
		if i.Cursor < 0 || i.Cursor > MAX_MEM{
			return fmt.Errorf("cursor range exceeds valid bounds")
		}
	case *ast.Input:
		b := make([]byte, 1)
		_, err := i.Reader.Read(b)
		if err != nil {
			if err == io.EOF{
				return fmt.Errorf("EOF")
			} else {
				return err
			}
		}
		i.Memory[i.Cursor] = b[0]
	case *ast.Output:
		b := []byte{i.Memory[i.Cursor]}
		i.Writer.Write(b)
	case *ast.Loop:
		for i.Memory[i.Cursor] != 0 {
			if err := i.runExpressions(e.Body); err != nil{
				return err
			}
		}
	}

	return nil
}