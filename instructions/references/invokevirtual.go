package references

import (
	"fmt"
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
	"github.com/zjmeow/zjvm/rtda/heap"
)

// 调用实例对象的方法
type InvokeVirtual struct {
	base.Index16Instruction
}

func (ins *InvokeVirtual) Execute(frame *rtda.Frame) {
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	methodRef := cp.GetConstant(ins.Index).(*heap.MethodRef)
	resolveMethod := methodRef.ResolvedMethod()
	if resolveMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	// 拿到调用方的引用，如果为空抛出空异常
	ref := frame.OperandStack().GetRefFromTop(resolveMethod.ArgSlotCount() - 1)
	if methodRef.Name() == "println" {
		goStr := heap.GoString(frame.OperandStack().PopRef())
		fmt.Println(goStr)
		return
	}
	// 判断权限
	if resolveMethod.IsProtected() &&
		resolveMethod.Class().IsSubClass(currentClass) &&
		resolveMethod.Class().GetPackageName() != currentClass.GetPackageName() &&
		ref.Class() != currentClass &&
		!ref.Class().IsSubClass(currentClass) {
		panic("java.lang.IllegalAccessError")
	}
	methodToBeInvoked := heap.LookupMethodInClass(ref.Class(), methodRef.Name(), methodRef.Descriptor())
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	base.InvokeMethod(frame, methodToBeInvoked)
}
