package comparisons

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
)

func NewLCMP() *LCmp  { return &LCmp{} }
func NewFCMPG() *FCmp { return &FCmp{g: true} }
func NewFCMPL() *FCmp { return &FCmp{g: false} }
func NewDCMPG() *DCmp { return &DCmp{g: true} }
func NewDCMPL() *DCmp { return &DCmp{g: false} }

type LCmp struct {
	base.NoOperandsInstruction
}

func (ins *LCmp) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	switch {
	case v1 > v2:
		stack.PushInt(1)
	case v1 == v2:
		stack.PushInt(0)
	case v1 < v2:
		stack.PushInt(-1)
	}
}

type FCmp struct {
	base.NoOperandsInstruction
	g bool
}

func (ins *FCmp) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	switch {
	case v1 > v2:
		stack.PushInt(1)
	case v1 == v2:
		stack.PushInt(0)
	case v1 < v2:
		stack.PushInt(-1)
	case ins.g:
		stack.PushInt(1)
	default:
		stack.PushInt(-1)
	}
}

type DCmp struct {
	base.NoOperandsInstruction
	g bool
}

func (ins *DCmp) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	switch {
	case v1 > v2:
		stack.PushInt(1)
	case v1 == v2:
		stack.PushInt(0)
	case v1 < v2:
		stack.PushInt(-1)
	case ins.g:
		stack.PushInt(1)
	default:
		stack.PushInt(-1)
	}
}
