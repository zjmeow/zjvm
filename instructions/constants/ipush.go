package constants

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
)

type BIPush struct{ val int8 }
type SIPush struct{ val int16 }

func (ins *BIPush) FetchOperands(reader *base.BytecodeReader) {
	ins.val = reader.ReadInt8()
}
func (ins *BIPush) Execute(frame *rtda.Frame) {
	i := int32(ins.val)
	frame.OperandStack().PushInt(i)
}

func (ins *SIPush) FetchOperands(reader *base.BytecodeReader) {
	ins.val = reader.ReadInt16()
}
func (ins *SIPush) Execute(frame *rtda.Frame) {
	i := int32(ins.val)
	frame.OperandStack().PushInt(i)
}
