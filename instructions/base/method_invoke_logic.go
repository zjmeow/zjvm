package base

import (
	"github.com/zjmeow/zjvm/native"
	"github.com/zjmeow/zjvm/rtda"
	"github.com/zjmeow/zjvm/rtda/heap"
)

func InvokeMethod(invokerFrame *rtda.Frame, method *heap.Method) {
	thread := invokerFrame.Thread()
	newFrame := thread.NewFrame(method)
	thread.PushFrame(newFrame)
	argSlot := int(method.ArgSlotCount())
	if argSlot > 0 {
		for i := argSlot - 1; i >= 0; i-- {
			slot := invokerFrame.OperandStack().PopSlot()
			newFrame.LocalVars().SetSlot(uint(i), slot)
		}
	}
	if method.IsNative() {
		nativeMethod := native.FindNativeMethod(method.Class().Name(), method.Name(), method.Descriptor())
		nativeMethod(newFrame)
	}
}
