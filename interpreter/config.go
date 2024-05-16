package interpreter

import "io"

type Config struct {
	MemorySize   int
	MaxStep      int

	Reader       io.Reader
	Writer       io.Writer
}