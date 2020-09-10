package lang

import (
	"github.com/zjmeow/zjvm/native"
	"github.com/zjmeow/zjvm/rtda"
)

func init() {
	native.Register("java/lang/Object", "getClass", "()Ljava/lang/Class;", getC)
	native.Register("java/lang/Object", "hashCode", "()I", hashCode)
	native.Register("java/lang/Object", "clone", "()Ljava/lang/Object;", clone)

}

func getC(frame *rtda.Frame) {

}

func hashCode(frame *rtda.Frame) {

}
func clone(frame *rtda.Frame) {

}
