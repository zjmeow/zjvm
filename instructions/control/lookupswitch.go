package control

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
)

type LookupSwitch struct {
	defaultOffset int32
	npairs        int32
	matchOffsets  []int32
}

func (ins *LookupSwitch) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	ins.defaultOffset = reader.ReadInt32()
	ins.npairs = reader.ReadInt32()
	ins.matchOffsets = reader.ReadInt32s(ins.npairs * 2)
}

func (ins *LookupSwitch) Execute(frame *rtda.Frame) {
	key := frame.OperandStack().PopInt()
	for i := int32(0); i < ins.npairs*2; i += 2 {
		if ins.matchOffsets[i] == key {
			base.Branch(frame, int(ins.matchOffsets[i+1]))
			return
		}
	}
	base.Branch(frame, int(ins.defaultOffset))
}
