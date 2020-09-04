package loads

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
)

type LongLoad struct {
	base.NoOperandsInstruction
	index uint
}
type LLOAD struct{ base.Index8Instruction }

func NewLongLoad(index uint) *LongLoad {
	return &LongLoad{
		index: index,
	}
}

func lLoad(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetLong(index)
	frame.OperandStack().PushLong(val)
}
func (ins *LongLoad) Execute(frame *rtda.Frame) {
	lLoad(frame, ins.index)
}

func (ins *LLOAD) Execute(frame *rtda.Frame) {
	lLoad(frame, ins.Index)
}
