package references

import (
	"fmt"
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
	"github.com/zjmeow/zjvm/rtda/heap"
)

type AThrow struct {
	base.NoOperandsInstruction
}

func (a *AThrow) Execute(frame *rtda.Frame) {
	ex := frame.OperandStack().PopRef()
	if ex == nil {
		panic("java.lang.NullPointException")
	}
	thread := frame.Thread()
	if !findAndGotoExceptionHandler(thread, ex) {
		handleUncaughtException(thread, ex)
	}
}

func findAndGotoExceptionHandler(thread *rtda.Thread, ex *heap.Object) bool {
	for thread.StackIsEmpty() {
		frame := thread.TopFrame()
		pc := frame.NextPc() - 1
		handlerPc := frame.Method().FindExceptionHandler(ex.Class(), pc)
		if handlerPc > 0 {
			stack := frame.OperandStack()
			stack.Clear()
			stack.PushRef(ex)
			frame.SetNextPc(handlerPc)
			return true
		}
	}
	return false
}

func handleUncaughtException(thread *rtda.Thread, ex *heap.Object) {
	thread.ClearStack()
	// todo 打印详细信息
	fmt.Println(ex)
}
