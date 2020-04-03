package rtda

type OperandStack struct {
	size  uint
	slots LocalVars
}

func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots: make([]Slot, maxStack),
		}
	}
	return nil
}

func (os *OperandStack) PushInt(val int32) {
	os.slots.SetInt(os.size, val)
	os.size++
}
func (os *OperandStack) PopInt() int32 {
	os.size--
	return os.slots.GetInt(os.size)
}
func (os *OperandStack) PushFloat(val float32) {
	os.slots.SetFloat(os.size, val)
	os.size++
}
func (os *OperandStack) PopFloat() float32 {
	os.size--
	return os.slots.GetFloat(os.size)
}

func (os *OperandStack) PushLong(val int64) {
	os.slots.SetLong(os.size, val)
	os.size += 2
}

func (os *OperandStack) PopLong() int64 {
	os.size -= 2
	return os.slots.GetLong(os.size)
}

func (os *OperandStack) PushDouble(val float64) {
	os.slots.SetDouble(os.size, val)
	os.size += 2
}

func (os *OperandStack) PopDouble() float64 {
	os.size -= 2
	return os.slots.GetDouble(os.size)
}

func (os *OperandStack) PushRef(val *Object) {
	os.slots.SetRef(os.size, val)
	os.size++
}

func (os *OperandStack) PopRef() *Object {
	os.size--
	ref := os.slots.GetRef(os.size)
	os.slots.SetRef(os.size, nil)
	return ref
}
