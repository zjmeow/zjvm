package lang

import (
	"github.com/zjmeow/zjvm/native"
	"github.com/zjmeow/zjvm/rtda"
)

func init() {
	native.Register("java/lang/Throwable",
		"fillInStackTrace",
		"(I)Ljava/lang/Throwable;",
		fillInStackTrace)
}

func fillInStackTrace(frame *rtda.Frame) {

}
