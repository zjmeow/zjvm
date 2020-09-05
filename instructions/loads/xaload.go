package loads

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/instructions/references"
	"github.com/zjmeow/zjvm/rtda"
	"github.com/zjmeow/zjvm/rtda/heap"
)

func NewIALoad() *ArrayLoad { return &ArrayLoad{arrayType: references.ArrayInt} }
func NewLALoad() *ArrayLoad { return &ArrayLoad{arrayType: references.ArrayLong} }
func NewFALoad() *ArrayLoad { return &ArrayLoad{arrayType: references.ArrayFloat} }
func NewDALoad() *ArrayLoad { return &ArrayLoad{arrayType: references.ArrayDouble} }
func NewAALoad() *ArrayLoad { return &ArrayLoad{arrayType: references.ArrayRef} }
func NewBALoad() *ArrayLoad { return &ArrayLoad{arrayType: references.ArrayByte} }
func NewCALoad() *ArrayLoad { return &ArrayLoad{arrayType: references.ArrayChar} }
func NewSALoad() *ArrayLoad { return &ArrayLoad{arrayType: references.ArrayShort} }

// 从数组中load出数据
type ArrayLoad struct {
	base.NoOperandsInstruction
	arrayType int
}

func (a *ArrayLoad) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNull(arrRef)
	checkIndex(arrRef.ArrayLength(), index)
	switch a.arrayType {
	case references.ArrayByte:
		stack.PushInt(int32(arrRef.Bytes()[index]))
	case references.ArrayChar:
		stack.PushInt(int32(arrRef.Chars()[index]))
	case references.ArrayShort:
		stack.PushInt(int32(arrRef.Shorts()[index]))
	case references.ArrayInt:
		stack.PushInt(arrRef.Ints()[index])
	case references.ArrayLong:
		stack.PushLong(arrRef.Longs()[index])
	case references.ArrayFloat:
		stack.PushFloat(arrRef.Floats()[index])
	case references.ArrayDouble:
		stack.PushDouble(arrRef.Doubles()[index])
	case references.ArrayRef:
		stack.PushRef(arrRef.Refs()[index])
	}
}

func checkNotNull(object *heap.Object) {
	if object == nil {
		panic("java.lang.NullPointerException")
	}
}
func checkIndex(length int32, index int32) {
	if index < 0 || index >= length {
		panic("ArrayIndexOutOfBoundsException")
	}
}
