package rtda

import "github.com/zjmeow/zjvm/rtda/heap"

type Thread struct {
	pc    int
	stack *Stack
}

func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

func (t *Thread) Pc() int {
	return t.pc
}

func (t *Thread) SetPc(pc int) {
	t.pc = pc
}

func (t *Thread) NewFrame(method *heap.Method) *Frame {
	return newFrame(t, method)
}

func (t *Thread) PopFrame() *Frame {
	return t.stack.pop()
}

func (t *Thread) PushFrame(frame *Frame) {
	t.stack.push(frame)
}

func (t *Thread) TopFrame() (frame *Frame) {
	return t.stack.top()
}
func (t *Thread) StackIsEmpty() bool {
	return t.stack.IsEmpty()
}
func (t *Thread) ClearStack() {
	for !t.stack.IsEmpty() {
		t.PopFrame()
	}
}

func (t *Thread) GetFrames() []*Frame {
	return t.stack.getFrames()
}
