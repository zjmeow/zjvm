package stores

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/instructions/references"
	"github.com/zjmeow/zjvm/rtda"
	"github.com/zjmeow/zjvm/rtda/heap"
)

func NewIAStore() *ArrayStore { return &ArrayStore{arrayType: references.ArrayInt} }
func NewLAStore() *ArrayStore { return &ArrayStore{arrayType: references.ArrayLong} }
func NewFAStore() *ArrayStore { return &ArrayStore{arrayType: references.ArrayFloat} }
func NewDAStore() *ArrayStore { return &ArrayStore{arrayType: references.ArrayDouble} }
func NewAAStore() *ArrayStore { return &ArrayStore{arrayType: references.ArrayRef} }
func NewBAStore() *ArrayStore { return &ArrayStore{arrayType: references.ArrayByte} }
func NewCAStore() *ArrayStore { return &ArrayStore{arrayType: references.ArrayChar} }
func NewSAStore() *ArrayStore { return &ArrayStore{arrayType: references.ArrayShort} }

// 从数组中Store出数据
type ArrayStore struct {
	base.NoOperandsInstruction
	arrayType int
}

func (a *ArrayStore) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNull(arrRef)
	checkIndex(arrRef.ArrayLength(), index)
	switch a.arrayType {
	case references.ArrayByte:
		arrRef.Bytes()[index] = int8(stack.PopInt())
	case references.ArrayChar:
		arrRef.Chars()[index] = uint16(stack.PopInt())
	case references.ArrayShort:
		arrRef.Shorts()[index] = int16(stack.PopInt())
	case references.ArrayInt:
		arrRef.Ints()[index] = stack.PopInt()
	case references.ArrayLong:
		arrRef.Longs()[index] = stack.PopLong()
	case references.ArrayFloat:
		arrRef.Floats()[index] = stack.PopFloat()
	case references.ArrayDouble:
		arrRef.Doubles()[index] = stack.PopDouble()
	case references.ArrayRef:
		arrRef.Refs()[index] = stack.PopRef()
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
