package references

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
	"github.com/zjmeow/zjvm/rtda/heap"
)

type MultiANewArray struct {
	index      uint16
	dimensions uint8
}

func (n *MultiANewArray) FetchOperands(reader *base.BytecodeReader) {
	n.index = reader.ReadUint16()
	n.dimensions = reader.ReadUint8()
}

func (n *MultiANewArray) Execute(frame *rtda.Frame) {
	cp := frame.ConstantPool()
	classRef := cp.GetConstant(uint(n.index)).(*heap.ClassRef)
	arrayClass := classRef.ResolveClass()
	stack := frame.OperandStack()
	counts := popAndCheckCounts(stack, int(n.dimensions))
	array := newMultiDimensionalArray(counts, arrayClass)
	stack.PushRef(array)

}
func popAndCheckCounts(stack *rtda.OperandStack, dimensions int) []int32 {
	counts := make([]int32, dimensions)
	for i := dimensions - 1; i >= 0; i-- {
		counts[i] = stack.PopInt()
		if counts[i] < 0 {
			panic("java.lang.NegativeArraySizeException")
		}
	}
	return counts

}
func newMultiDimensionalArray(counts []int32, arrayClass *heap.Class) *heap.Object {
	count := uint(counts[0])
	array := arrayClass.NewArray(count)
	if len(counts) > 1 {
		refs := array.Refs()
		for i := range refs {
			refs[i] = newMultiDimensionalArray(counts[1:], arrayClass.ComponentClass())
		}
	}
	return array
}
