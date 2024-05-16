package interpreter

import (
	"context"
	"fmt"
	"io"

	"github.com/yuki-maruyama/brainfxxk/ast"
	"github.com/yuki-maruyama/brainfxxk/parser"
)

type Interpreter struct {
	Program *ast.Program
	Reader  io.Reader
	Writer  io.Writer

	Memory  []byte
	Cursor  int
}

func Run(ctx context.Context, script string, config Config) error{
	p, err := parser.Parse(script)
	if err != nil {
		return err
	}
	
	return NewInterpreter(p, config).Run(ctx)
}

func NewInterpreter(p *ast.Program, config Config) *Interpreter{
	return &Interpreter{
		Program: p,
		Reader: config.Reader,
		Writer: config.Writer,
		Memory: make([]byte, config.MemorySize),
		Cursor: 0,
	}
}

func (i *Interpreter) Run(ctx context.Context) error{
	err := i.runExpressions(ctx, i.Program.Body)
	return err
} 

func (i *Interpreter) runExpressions(ctx context.Context, exprs []ast.Node) error{
	for _, expr := range exprs {
		select{
		case <- ctx.Done():
			return ctx.Err()
		default:
			err := i.runExpression(ctx, expr)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (i *Interpreter) runExpression(ctx context.Context, expr ast.Node) error{
	switch e := expr.(type) {
	case *ast.Increment:
		i.Memory[i.Cursor]++
	case *ast.Decrement:
		i.Memory[i.Cursor]--
	case *ast.MoveRight:
		i.Cursor++
		if i.Cursor < 0 || i.Cursor >= len(i.Memory){
			return fmt.Errorf("cursor range exceeds valid bounds")
		}
	case *ast.MoveLeft:
		i.Cursor--
		if i.Cursor < 0 || i.Cursor >= len(i.Memory){
			return fmt.Errorf("cursor range exceeds valid bounds")
		}
	case *ast.Input:
		b := make([]byte, 1)
		if i.Reader == nil {
			return fmt.Errorf("input error")
		}
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
		if i.Writer == nil {
			return fmt.Errorf("output error")
		}
		if _, err := i.Writer.Write(b); err != nil{
			return err
		}
	case *ast.Loop:
		for i.Memory[i.Cursor] != 0 {
			if err := i.runExpressions(ctx, e.Body); err != nil{
				return err
			}
			select {
			case <- ctx.Done():
				return ctx.Err()
			default:
			}
		}
	}

	return nil
}