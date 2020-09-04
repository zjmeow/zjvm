package references

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
	"github.com/zjmeow/zjvm/rtda/heap"
)

// 指令用于调用一些需要特殊处理的实例方法，包括实例初始化方法、私有方法和父类方法
// 因为不需要动态绑定，所以使用这个方法会加快调用速度
type InvokeSpecial struct {
	base.Index16Instruction
}

func (ins *InvokeSpecial) Execute(frame *rtda.Frame) {
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	methodRef := cp.GetConstant(ins.Index).(*heap.MethodRef)
	resolvedClass := methodRef.ResolveClass()
	resolveMethod := methodRef.ResolvedMethod()
	if resolveMethod.Name() == "<init>" && resolveMethod.Class() != resolvedClass {
		panic("java.lang.NoSuchMethodError")
	}
	if resolveMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	// 弹出this引用
	ref := frame.OperandStack().GetRefFromTop(resolveMethod.ArgSlotCount())
	if ref == nil {
		panic("java.lang.NullPointException")
	}
	// 如果是在子类里调用的，需要检查父类的方法权限是不是protect以上的
	if resolveMethod.IsProtected() &&
		resolveMethod.Class().IsSubClass(currentClass) &&
		resolveMethod.Class().GetPackageName() != currentClass.GetPackageName() &&
		ref.Class() != currentClass &&
		!ref.Class().IsSubClass(currentClass) {
		panic("java.lang.IllegalAccessError")
	}
	methodToBeInvoked := resolveMethod
	// 如果该类子类，被转换成了父类，仍然去调用子类的方法，init方法除外
	if currentClass.IsSuper() && currentClass.IsSubClass(currentClass) && resolveMethod.Name() != "<init>" {
		methodToBeInvoked = heap.LookupMethodInClass(currentClass.SuperClass(), methodRef.Name(),
			methodRef.Descriptor())
	}

	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	base.InvokeMethod(frame, methodToBeInvoked)
}
