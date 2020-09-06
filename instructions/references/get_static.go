package references

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
	"github.com/zjmeow/zjvm/rtda/heap"
)

type GetStatic struct {
	base.Index16Instruction
}

func (g *GetStatic) Execute(frame *rtda.Frame) {
	method := frame.Method()
	class := method.Class()
	cp := frame.ConstantPool()
	fieldRef := cp.GetConstant(g.Index).(*heap.FieldRef)
	field := fieldRef.ResolveField()
	if !field.Class().InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), field.Class())
		return
	}
	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	// 需要初始化的对象在栈的第一个位置
	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := class.StaticVars()
	stack := frame.OperandStack()
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
