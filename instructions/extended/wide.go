package extended

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/instructions/loads"
	"github.com/zjmeow/zjvm/instructions/math"
	"github.com/zjmeow/zjvm/instructions/stores"
)

type Wide struct {
	modifiedInstruction base.Instruction
}

func (ins *Wide) FetchOperands(reader *base.BytecodeReader) {
	op := reader.ReadUint8()
	switch op {
	case 0x15:
		instr := &loads.ILOAD{}
		instr.Index = uint(reader.ReadUint16())
		ins.modifiedInstruction = instr
	case 0x16:
		instr := &loads.LLOAD{}
		instr.Index = uint(reader.ReadUint16())
		ins.modifiedInstruction = instr
	case 0x17:
		instr := &loads.FLOAD{}
		instr.Index = uint(reader.ReadUint16())
		ins.modifiedInstruction = instr
	case 0x18:
		instr := &loads.DLOAD{}
		instr.Index = uint(reader.ReadUint16())
		ins.modifiedInstruction = instr
	case 0x19:
		instr := &loads.ALOAD{}
		instr.Index = uint(reader.ReadUint16())
		ins.modifiedInstruction = instr
	case 0x36:
		instr := &stores.ISTORE{}
		instr.Index = uint(reader.ReadUint16())
		ins.modifiedInstruction = instr
	case 0x37:
		instr := &stores.LSTORE{}
		instr.Index = uint(reader.ReadUint16())
		ins.modifiedInstruction = instr
	case 0x38:
		instr := &stores.FSTORE{}
		instr.Index = uint(reader.ReadUint16())
		ins.modifiedInstruction = instr
	case 0x39:
		instr := &stores.DSTORE{}
		instr.Index = uint(reader.ReadUint16())
		ins.modifiedInstruction = instr
	case 0x3a:
		instr := &stores.ASTORE{}
		instr.Index = uint(reader.ReadUint16())
		ins.modifiedInstruction = instr
	case 0x84:
		instr := &math.IInc{}
		instr.Index = uint(reader.ReadUint16())
		instr.Const = int32(reader.ReadUint16())
		ins.modifiedInstruction = instr
	case 0x9a: // TODO ret
		panic("0x9a")
	}

}
