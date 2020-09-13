package lang

import (
	"github.com/zjmeow/zjvm/native"
	"github.com/zjmeow/zjvm/rtda"
	"github.com/zjmeow/zjvm/rtda/heap"
)

func init() {
	native.Register("java/lang/Class", "getPrimitiveClass",
		"(Ljava/lang/String;)Ljava/lang/Class;", getPrimitiveClass)
	native.Register("java/lang/Class", "getName0", "()Ljava/lang/String;", getName0)
	native.Register("java/lang/Class", "desiredAssertionStatus0",
		"(Ljava/lang/Class;)Z", desiredAssertionStatus0)

}
func getPrimitiveClass(frame *rtda.Frame) {
	nameObj := frame.LocalVars().GetRef(0)
	name := heap.GoString(nameObj)
	loader := frame.Method().Class().ClassLoader()
	jStr := heap.JString(loader, name)
	frame.OperandStack().PushRef(jStr)
}

func getName0(frame *rtda.Frame) {
	this := frame.LocalVars().GetRef(0)
	class := this.Extra().(*heap.Class)
	name := class.Name()
	nameObj := heap.JString(class.ClassLoader(), name)
	frame.OperandStack().PushRef(nameObj)

}

func desiredAssertionStatus0(frame *rtda.Frame) {
	// push bool false
	frame.OperandStack().PushInt(0)
}
