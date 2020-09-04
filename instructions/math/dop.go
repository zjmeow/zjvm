package math

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
	"math"
)

type DOp struct {
	base.NoOperandsInstruction
	op  func(a, b float64) float64
	div bool
}

func (ins *DOp) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	if ins.div && v2 == 0 {
		panic("div by zero")
	}
	stack.PushDouble(ins.op(v1, v2))
}

type DNeg struct {
	base.NoOperandsInstruction
}

func (ins *DNeg) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v := stack.PopDouble()
	stack.PushDouble(-v)
}

func NewDNeg() *DNeg { return &DNeg{} }
func NewDAdd() *DOp  { return &DOp{op: dAdd} }
func NewDSub() *DOp  { return &DOp{op: dSub} }
func NewDMul() *DOp  { return &DOp{op: dMul} }
func NewDDiv() *DOp  { return &DOp{op: dDiv, div: true} }
func NewDRem() *DOp  { return &DOp{op: dRem, div: true} }

func dAdd(a, b float64) float64 { return a + b }
func dSub(a, b float64) float64 { return a - b }
func dMul(a, b float64) float64 { return a * b }
func dDiv(a, b float64) float64 { return a / b }
func dRem(a, b float64) float64 { return math.Mod(a, b) }
