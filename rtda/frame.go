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

func newFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread:       thread,
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
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
