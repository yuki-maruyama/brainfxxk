package ast

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