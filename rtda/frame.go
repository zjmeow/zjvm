package rtda

import "github.com/zjmeow/zjvm/rtda/heap"

type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
	thread       *Thread
	method       *heap.Method
	nextPc       int
}

func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:       thread,
		method:       method,
		localVars:    newLocalVars(uint(method.MaxLocals())),
		operandStack: newOperandStack(uint(method.MaxStack())),
	}
}

func (f *Frame) OperandStack() *OperandStack {
	return f.operandStack
}

func (f *Frame) LocalVars() LocalVars {
	return f.localVars
}
func (f *Frame) Thread() *Thread {
	return f.thread
}

// TODO
func (f *Frame) SetNextPc(pc int) {
	f.nextPc = pc
}
func (f *Frame) NextPc() int {
	return f.nextPc
}

func (f *Frame) ConstantPool() *heap.ConstantPool {
	return f.method.Class().ConstantPool()
}
func (f *Frame) Method() *heap.Method {
	return f.method
}
