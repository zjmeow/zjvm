package references

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
	"github.com/zjmeow/zjvm/rtda/heap"
)

type GetField struct {
	base.Index16Instruction
}

func (g *GetField) Execute(frame *rtda.Frame) {
	method := frame.Method()
	class := method.Class()
	cp := frame.ConstantPool()
	fieldRef := cp.GetConstant(g.Index).(*heap.FieldRef)
	field := fieldRef.ResolveField()
	fieldClass := field.Class()
	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	// 如果是final，则需要看是否是在本class的 clinit 中初始化的，如果不是就要抛出异常
	// clinit 初始化能保证是线程安全的，所以会被用来做单例初始化
	if field.IsFinal() {
		if class != fieldClass || method.Name() != "<clinit>" {
			panic("java.lang.IllegalAccessError")
		}
	}
	// 需要初始化的对象在栈的第一个位置
	descriptor := field.Descriptor()
	slotId := field.SlotId()
	stack := frame.OperandStack()
	ref := stack.PopRef()
	if ref == nil {
		panic("java.lang.NullPointException")
	}
	slots := ref.Fields()
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		stack.PushInt(slots.GetInt(slotId))
	case 'F':
		stack.PushFloat(slots.GetFloat(slotId))
	case 'J':
		stack.PushLong(slots.GetLong(slotId))
	case 'D':
		stack.PushDouble(slots.GetDouble(slotId))
	case 'L', '[':
		stack.PushRef(slots.GetRef(slotId))
	}

}
