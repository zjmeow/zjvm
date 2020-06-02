package math

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
)

type IOp struct {
	base.NoOperandsInstruction
	op  func(a, b int32) int32
	div bool
}

func (ins *IOp) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	if ins.div && v2 == 0 {
		panic("div by zero")
	}
	stack.PushInt(ins.op(v1, v2))
}

type IInc struct {
	Index uint
	Const int32
}

func (ins *IInc) FetchOperands(reader *base.BytecodeReader) {
	ins.Index = uint(reader.ReadUint8())
	ins.Const = int32(reader.ReadInt8())
}
func (ins *IInc) Execute(frame *rtda.Frame) {
	val := frame.LocalVars().GetInt(ins.Index)
	val += ins.Const
	frame.LocalVars().SetInt(ins.Index, val)
}

type INeg struct {
	base.NoOperandsInstruction
}

func (ins *INeg) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v := stack.PopInt()
	stack.PushInt(-v)
}

func NewIInc() *IInc { return &IInc{} }
func NewINeg() *INeg { return &INeg{} }
func NewIAdd() *IOp  { return &IOp{op: iAdd} }
func NewISub() *IOp  { return &IOp{op: iSub} }
func NewIMul() *IOp  { return &IOp{op: iMul} }
func NewIDiv() *IOp  { return &IOp{op: iDiv, div: true} }
func NewIRem() *IOp  { return &IOp{op: iRem, div: true} }
func NewIShl() *IOp  { return &IOp{op: iShl} }
func NewIShr() *IOp  { return &IOp{op: iShr} }
func NewIUShr() *IOp { return &IOp{op: iUshr} }
func NewIAnd() *IOp  { return &IOp{op: iAnd} }
func NewIOr() *IOp   { return &IOp{op: iOr} }
func NewIXor() *IOp  { return &IOp{op: iXor} }

func iAdd(a, b int32) int32  { return a + b }
func iSub(a, b int32) int32  { return a - b }
func iMul(a, b int32) int32  { return a * b }
func iDiv(a, b int32) int32  { return a / b }
func iRem(a, b int32) int32  { return a % b }
func iAnd(a, b int32) int32  { return a & b }
func iOr(a, b int32) int32   { return a | b }
func iXor(a, b int32) int32  { return a ^ b }
func iShl(a, b int32) int32  { return a << (b & 0x1f) }
func iShr(a, b int32) int32  { return a >> (b & 0x1f) }
func iUshr(a, b int32) int32 { return int32(uint32(a) >> (b & 0x1f)) }
