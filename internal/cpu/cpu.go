package cpu

import (
	"github.com/onrcayci/goshell/internal/interpreter"
	"github.com/onrcayci/goshell/internal/parser"
	"github.com/onrcayci/goshell/internal/ram"
)

type CPU struct {
	IP     int
	IR     string
	Quanta int
}

func New() *CPU {
	newCPU := CPU{}
	return &newCPU
}

func (c *CPU) Run() {
	for c.Quanta != 0 {
		c.IR = ram.RAM[c.IP]
		c.IP++
		c.Quanta--
		argc, argv := parser.ParseInput(c.IR)
		interpreter.Interpreter(argc, argv)
	}
}
