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
	switch a.arrayType {
	case references.ArrayByte:
		val := stack.PopInt()
		index, arrRef := checkAndGetRef(stack)
		arrRef.Bytes()[index] = int8(val)
	case references.ArrayChar:
		val := stack.PopInt()
		index, arrRef := checkAndGetRef(stack)
		arrRef.Chars()[index] = uint16(val)
	case references.ArrayShort:
		val := stack.PopInt()
		index, arrRef := checkAndGetRef(stack)
		arrRef.Shorts()[index] = int16(val)
	case references.ArrayInt:
		val := stack.PopInt()
		index, arrRef := checkAndGetRef(stack)
		arrRef.Ints()[index] = val
	case references.ArrayLong:
		val := stack.PopLong()
		index, arrRef := checkAndGetRef(stack)
		arrRef.Longs()[index] = val
	case references.ArrayFloat:
		val := stack.PopFloat()
		index, arrRef := checkAndGetRef(stack)
		arrRef.Floats()[index] = val
	case references.ArrayDouble:
		val := stack.PopDouble()
		index, arrRef := checkAndGetRef(stack)
		arrRef.Doubles()[index] = val
	case references.ArrayRef:
		val := stack.PopRef()
		index, arrRef := checkAndGetRef(stack)
		arrRef.Refs()[index] = val
	}
}

func checkAndGetRef(stack *rtda.OperandStack) (int32, *heap.Object) {
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNull(arrRef)
	checkIndex(arrRef.ArrayLength(), index)
	return index, arrRef
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
