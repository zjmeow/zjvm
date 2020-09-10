package lang

import (
	"github.com/zjmeow/zjvm/native"
	"github.com/zjmeow/zjvm/rtda"
	"github.com/zjmeow/zjvm/rtda/heap"
)

func init() {
	native.Register("java/lang/System",
		"arraycopy",
		"(Ljava/lang/Object;ILjava/lang/Object;II)V", arrayCopy)
}
func arrayCopy(frame *rtda.Frame) {
	vars := frame.LocalVars()
	src := vars.GetRef(0)
	srcPos := vars.GetInt(1)
	dest := vars.GetRef(2)
	destPos := vars.GetInt(3)
	length := vars.GetInt(4)
	if src == nil || dest == nil {
		panic("java.lang.NullPointerException")
	}
	if checkArray(src, dest) {
		panic("java.lang.ArrayStoreException")
	}
	if srcPos < 0 || destPos < 0 || length < 0 ||
		srcPos+length > src.ArrayLength() ||
		destPos+length > dest.ArrayLength() {
		panic("java.lang.IndexOutOfBoundsException")
	}
	heap.ArrayCopy(src, dest, srcPos, destPos, length)
}
func checkArray(src, dest *heap.Object) bool {
	srcClass := src.Class()
	destClass := dest.Class()
	if !(srcClass.IsArray() && destClass.IsArray()) {
		return false
	}
	if srcClass.ComponentClass().IsPrimitive() || destClass.IsPrimitive() {
		return srcClass == destClass
	}
	return true

}
