package value

import "github.com/evanphx/m13/insn"

type Code struct {
	NumRefs      int
	NumRegs      int
	Instructions []insn.Instruction
	Literals     []string
	SubCode      []*Code
}
