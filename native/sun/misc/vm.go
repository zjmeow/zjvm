package misc

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/native"
	"github.com/zjmeow/zjvm/rtda"
)

func init() {
	native.Register("sun/misc/VM", "initialize", "()V", initialize)
}

// private static native void initialize();
func initialize(frame *rtda.Frame) {
	classLoader := frame.Method().Class().ClassLoader()
	jlSysClass := classLoader.LoadClass("java/lang/System")
	initSysClass := jlSysClass.GetStaticMethod("initializeSystemClass", "()V")
	base.InvokeMethod(frame, initSysClass)
}
