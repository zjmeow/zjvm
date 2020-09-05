package references

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
	"github.com/zjmeow/zjvm/rtda/heap"
)

type ANewArray struct {
	base.Index16Instruction
}

// 根据index到常量池里找到类型，pop int 参数为初始化数组的长度
func (a *ANewArray) Execute(frame *rtda.Frame) {
	cp := frame.ConstantPool()
	// 这里拿到是原始类型，比如说Object，而不是Object[]
	classRef := cp.GetConstant(a.Index).(*heap.ClassRef)
	componentClass := classRef.ResolveClass()
	stack := frame.OperandStack()
	count := stack.PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}
	// 调用ArrayClass把类转成数组类型
	arrayClass := componentClass.ArrayClass()
	array := arrayClass.NewArray(uint(count))
	stack.PushRef(array)
}
