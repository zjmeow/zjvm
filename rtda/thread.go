package rtda

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

func (t *Thread) NewFrame(maxLocals, maxStack uint) *Frame {
	return newFrame(t, maxLocals, maxStack)
}

func (t *Thread) PopFrame() *Frame {
	return t.stack.pop()
}

func (t *Thread) PushFrame(frame *Frame) {
	t.stack.push(frame)
}
