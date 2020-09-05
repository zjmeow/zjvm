package references

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
	"github.com/zjmeow/zjvm/rtda/heap"
)

const (
	ArrayRef     = 0 //这个不是虚拟机定义的，是我们自己定义的
	ArrayBoolean = 4
	ArrayChar    = 5
	ArrayFloat   = 6
	ArrayDouble  = 7
	ArrayByte    = 8
	ArrayShort   = 9
	ArrayInt     = 10
	ArrayLong    = 11
)

type NewArray struct {
	arrayType uint8
}

func (n *NewArray) FetchOperands(reader *base.BytecodeReader) {
	n.arrayType = reader.ReadUint8()
}
func (n *NewArray) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	count := stack.PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}
	classLoader := frame.Method().Class().ClassLoader()
	arrayClass := getArrayClass(classLoader, n.arrayType)
	array := arrayClass.NewArray(uint(count))
	stack.PushRef(array)

}
func getArrayClass(loader *heap.ClassLoader, arrayType uint8) *heap.Class {
	switch arrayType {
	case ArrayBoolean:
		return loader.LoadClass("[Z")
	case ArrayByte:
		return loader.LoadClass("[B")
	case ArrayChar:
		return loader.LoadClass("[C")
	case ArrayShort:
		return loader.LoadClass("[S")
	case ArrayInt:
		return loader.LoadClass("[I")
	case ArrayLong:
		return loader.LoadClass("[J")
	case ArrayFloat:
		return loader.LoadClass("[F")
	case ArrayDouble:
		return loader.LoadClass("[D")
	default:
		panic("Invalid array type")
	}
}
