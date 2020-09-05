package references

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
)

type ArrayLength struct {
	base.NoOperandsInstruction
}

func (a *ArrayLength) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	arrayRef := stack.PopRef()
	if arrayRef == nil {
		panic("java.lang.NullPointException")
	}

	length := arrayRef.ArrayLength()
	stack.PushInt(length)
}
