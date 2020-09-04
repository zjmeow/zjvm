package references

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
	"github.com/zjmeow/zjvm/rtda/heap"
)

type InvokeInterface struct {
	index uint
	// count unit8 由于历史原因虽然没用但是仍然有这段
	// zero unit8 为了兼容某些虚拟机的字段，在本虚拟机中没用
}

func (i *InvokeInterface) FetchOperands(reader *base.BytecodeReader) {
	i.index = uint(reader.ReadInt16())
	// 丢弃两个由于历史原因没用到的字段
	reader.ReadUint8()
	reader.ReadUint8()
}

func (i *InvokeInterface) Execute(frame *rtda.Frame) {
	cp := frame.ConstantPool()
	methodRef := cp.GetConstant(i.index).(*heap.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()
	if resolvedMethod.IsStatic() || resolvedMethod.IsPrivate() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if ref == nil {
		panic("java.lang.NullPointException")
	}
	if !ref.Class().IsImplements(methodRef.ResolveClass()) {
		panic("java.lang.IncompatibleClassChangeError")
	}
	methodToBeInvoked := heap.LookupMethodInClass(ref.Class(), methodRef.Name(), methodRef.Descriptor())
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	// 接口方法必须是public的
	if !methodToBeInvoked.IsPublic() {
		panic("java.lang.IllegalAccessError")
	}
	base.InvokeMethod(frame, methodToBeInvoked)
}
