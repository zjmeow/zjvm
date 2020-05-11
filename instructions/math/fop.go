package math

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
	"math"
)

type FOp struct {
	base.NoOperandsInstruction
	op  func(a, b float32) float32
	div bool
}

func (ins *FOp) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopFloat()
	v2 := stack.PopFloat()
	if ins.div && v2 == 0 {
		panic("div by zero")
	}
	stack.PushFloat(ins.op(v1, v2))
}

type FNeg struct {
	base.NoOperandsInstruction
}

func (ins *FNeg) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v := stack.PopFloat()
	stack.PushFloat(-v)
}

func NewFNeg() *DNeg { return &DNeg{} }
func NewFAdd() *FOp  { return &FOp{op: fAdd} }
func NewFSub() *FOp  { return &FOp{op: fSub} }
func NewFMul() *FOp  { return &FOp{op: fMul} }
func NewFDiv() *FOp  { return &FOp{op: fDiv, div: true} }
func NewFRem() *FOp  { return &FOp{op: fRem, div: true} }

func fAdd(a, b float32) float32 { return a + b }
func fSub(a, b float32) float32 { return a - b }
func fMul(a, b float32) float32 { return a * b }
func fDiv(a, b float32) float32 { return a / b }
func fRem(a, b float32) float32 { return float32(math.Mod(float64(a), float64(b))) }
