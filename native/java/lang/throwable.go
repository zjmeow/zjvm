package lang

import (
	"github.com/zjmeow/zjvm/native"
	"github.com/zjmeow/zjvm/rtda"
	"github.com/zjmeow/zjvm/rtda/heap"
)

func init() {
	native.Register("java/lang/Throwable",
		"fillInStackTrace",
		"(I)Ljava/lang/Throwable;",
		fillInStackTrace)
}

func fillInStackTrace(frame *rtda.Frame) {
	//this := frame
	//frame.OperandStack().PushRef(this)
	//stes := createStackTrackElements(this, frame.Thread())
	//this.SetExtra(stes)
}

type StackTraceElement struct {
	fileName   string
	className  string
	methodName string
	lineNumber int
}

func createStackTrackElements(tObj *heap.Object, thread *rtda.Thread) []*StackTraceElement {
	// 需要跳过栈顶两帧，分别在执行fillInStackTrace(int)和fillInStackTrace() 方法
	skip := distanceToObject(tObj.Class()) + 2
	frames := thread.GetFrames()[skip:]
	stes := make([]*StackTraceElement, len(frames))
	for i, frame := range frames {
		stes[i] = createStackTrackElement(frame)
	}
	return stes
}

func createStackTrackElement(frame *rtda.Frame) *StackTraceElement {
	method := frame.Method()
	class := method.Class()
	return &StackTraceElement{
		// todo 文件名和行数未实现
		fileName:   "unknown",
		className:  class.Name(),
		methodName: method.Name(),
		lineNumber: 0,
	}
}

func distanceToObject(class *heap.Class) int {
	distance := 0
	for c := class.SuperClass(); c != nil; c = c.SuperClass() {
		distance++
	}
	return distance
}
