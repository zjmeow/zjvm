package heap

func (c *Class) NewArray(count uint) *Object {
	if !c.IsArray() {
		panic("Not array class:" + c.name)
	}
	switch c.name {
	case "[Z":
		NewDefaultObject(c, make([]int8, count))
	case "[B":
		return NewDefaultObject(c, make([]int8, count))
	case "[C":
		return NewDefaultObject(c, make([]uint16, count))
	case "[S":
		return NewDefaultObject(c, make([]int16, count))
	case "[I":
		return NewDefaultObject(c, make([]int32, count))
	case "[J":
		return NewDefaultObject(c, make([]int64, count))
	case "[F":
		return NewDefaultObject(c, make([]float32, count))

	case "[D":
		return NewDefaultObject(c, make([]float64, count))
	default:
		return NewDefaultObject(c, make([]*Object, count))
	}
	return NewDefaultObject(c, make([]*Object, count))
}
