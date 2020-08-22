package references

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
	"github.com/zjmeow/zjvm/rtda/heap"
)

type New struct {
	base.Index16Instruction
}

// 操作数是一个16位的索引，索引可以用来在常量池里找到一个类符号引用
// 解析这个符号引用，然后创建对象，把对象引入栈顶，new工作就完成了
func (n *New) Execute(frame *rtda.Frame) {
	cp := frame.ConstantPool()
	classRef := cp.GetConstant(n.Index).(*heap.ClassRef)
	class := classRef.ResolveClass()
	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}
	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}
