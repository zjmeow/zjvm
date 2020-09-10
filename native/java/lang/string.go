package lang

import (
	"github.com/zjmeow/zjvm/native"
	"github.com/zjmeow/zjvm/rtda"
	"github.com/zjmeow/zjvm/rtda/heap"
)

func init() {
	native.Register("java/lang/String", "intern", "()Ljava/lang/String;", inter)
}
func inter(frame *rtda.Frame) {
	this := frame.LocalVars().GetRef(0)
	interned := heap.InternString(this)
	frame.OperandStack().PushRef(interned)
}
