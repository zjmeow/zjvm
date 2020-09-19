package reserved

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/native"
	_ "github.com/zjmeow/zjvm/native/java/io"
	_ "github.com/zjmeow/zjvm/native/java/lang"
	_ "github.com/zjmeow/zjvm/native/sun/misc"
	"github.com/zjmeow/zjvm/rtda"
)

type InvokeNative struct {
	base.NoOperandsInstruction
}

func (i *InvokeNative) Execute(frame *rtda.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	descriptor := method.Descriptor()
	methodName := method.Name()
	nativeMethod := native.FindNativeMethod(className, methodName, descriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + "." + descriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
	}
	nativeMethod(frame)
}
