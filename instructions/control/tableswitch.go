package control

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
)

type TableSwitch struct {
	defaultOffset int32
	low           int32
	high          int32
	jumpOffset    []int32
}

func (ins *TableSwitch) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	ins.defaultOffset = reader.ReadInt32()
	ins.low = reader.ReadInt32()
	ins.high = reader.ReadInt32()
	count := ins.high - ins.low + 1
	ins.jumpOffset = reader.ReadInt32s(count)
}

func (ins *TableSwitch) Execute(frame *rtda.Frame) {
	idx := frame.OperandStack().PopInt()
	var offset int
	if idx >= ins.low && idx <= ins.high {
		offset = int(ins.jumpOffset[idx-ins.low])
	} else {
		offset = int(ins.defaultOffset)
	}
	base.Branch(frame, offset)
}
