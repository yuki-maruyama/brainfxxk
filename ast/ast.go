package ast

import "fmt"

type Node interface {
	node()
}

type MoveRight struct{}
type MoveLeft struct{}
type Increment struct{}
type Decrement struct{}
type Output struct{}
type Input struct{}
type Loop struct {
	Body []Node
}

type Program struct{
	Body []Node
}

func (*MoveRight) node()    {}
func (*MoveLeft) node()     {}
func (*Increment) node()    {}
func (*Decrement) node()    {}
func (*Output) node()       {}
func (*Input) node()        {}
func (*Loop) node()         {}
func (*Program) node()      {}

func PrintAST(node Node, indent string) {
	switch n := node.(type) {
	case *Program:
		fmt.Println(indent + "Program")
		for _, child := range n.Body {
			PrintAST(child, indent+"  ")
		}
	case *MoveRight:
		fmt.Println(indent + "MoveRight")
	case *MoveLeft:
		fmt.Println(indent + "MoveLeft")
	case *Increment:
		fmt.Println(indent + "Increment")
	case *Decrement:
		fmt.Println(indent + "Decrement")
	case *Output:
		fmt.Println(indent + "Output")
	case *Input:
		fmt.Println(indent + "Input")
	case *Loop:
		fmt.Println(indent + "Loop")
		for _, child := range n.Body {
			PrintAST(child, indent+"  ")
		}
	default:
		fmt.Println(indent + "Unknown node")
	}
}
