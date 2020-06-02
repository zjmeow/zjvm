package control

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
)

type Goto struct {
	base.BranchInstruction
}

func (ins *Goto) Execute(frame *rtda.Frame) {
	base.Branch(frame, ins.Offset)
}

type GotoW struct {
	base.BranchInstruction
}

func (ins *GotoW) Execute(frame *rtda.Frame) {
	base.Branch(frame, ins.Offset)
}

func (ins *GotoW) FetchOperands(reader *base.BytecodeReader) {
	ins.Offset = int(reader.ReadInt32())
}
