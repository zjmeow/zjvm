package math

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
)

type LOp struct {
	base.NoOperandsInstruction
	op  func(a, b int64) int64
	div bool
}

func (ins *LOp) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if ins.div && v2 == 0 {
		panic("div by zero")
	}
	stack.PushLong(ins.op(v1, v2))
}

type LInc struct {
	base.NoOperandsInstruction
}

func (ins *LInc) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v := stack.PopLong()
	stack.PushLong(v + 1)
}

type LNeg struct {
	base.NoOperandsInstruction
}

func (ins *LNeg) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v := stack.PopLong()
	stack.PushLong(-v)
}

func NewLInc() *IInc { return &IInc{} }
func NewLNeg() *INeg { return &INeg{} }
func NewLAdd() *LOp  { return &LOp{op: lAdd} }
func NewLSub() *LOp  { return &LOp{op: lSub} }
func NewLMul() *LOp  { return &LOp{op: lMul} }
func NewLDiv() *LOp  { return &LOp{op: lDiv, div: true} }
func NewLRem() *LOp  { return &LOp{op: lRem, div: true} }
func NewLShl() *LOp  { return &LOp{op: lShl} }
func NewLShr() *LOp  { return &LOp{op: lShr} }
func NewLUShr() *LOp { return &LOp{op: lUshr} }
func NewLAnd() *LOp  { return &LOp{op: lAnd} }
func NewLOr() *LOp   { return &LOp{op: lOr} }
func NewLXor() *LOp  { return &LOp{op: lXor} }

func lAdd(a, b int64) int64  { return a + b }
func lSub(a, b int64) int64  { return a - b }
func lMul(a, b int64) int64  { return a * b }
func lDiv(a, b int64) int64  { return a / b }
func lRem(a, b int64) int64  { return a % b }
func lAnd(a, b int64) int64  { return a & b }
func lOr(a, b int64) int64   { return a | b }
func lXor(a, b int64) int64  { return a ^ b }
func lShl(a, b int64) int64  { return a << (b & 0x1f) }
func lShr(a, b int64) int64  { return a >> (b & 0x1f) }
func lUshr(a, b int64) int64 { return int64(uint64(a) >> (b & 0x3f)) }
