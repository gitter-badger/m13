// Code generated by "stringer -type=Op"; DO NOT EDIT

package insn

import "fmt"

const _Op_name = "NoopStoreIntCopyRegCallNResetReturnGIFCall0Goto"

var _Op_index = [...]uint8{0, 4, 12, 19, 24, 29, 35, 38, 43, 47}

func (i Op) String() string {
	if i < 0 || i >= Op(len(_Op_index)-1) {
		return fmt.Sprintf("Op(%d)", i)
	}
	return _Op_name[_Op_index[i]:_Op_index[i+1]]
}
