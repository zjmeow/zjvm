package references

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
	"github.com/zjmeow/zjvm/rtda/heap"
)

const (
	BOOLEAN = 4
	CHAR    = 5
	FLOAT   = 6
	DOUBLE  = 7
	BYTE    = 8
	SHORT   = 9
	INT     = 10
	LONG    = 11
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
	case BOOLEAN:
		return loader.LoadClass("[Z")
	case BYTE:
		return loader.LoadClass("[B")
	case CHAR:
		return loader.LoadClass("[C")
	case SHORT:
		return loader.LoadClass("[S")
	case INT:
		return loader.LoadClass("[I")
	case LONG:
		return loader.LoadClass("[J")
	case FLOAT:
		return loader.LoadClass("[F")
	case DOUBLE:
		return loader.LoadClass("[D")
	default:
		panic("Invalid array type")
	}
}
