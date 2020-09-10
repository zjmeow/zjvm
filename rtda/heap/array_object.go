package heap

func (o *Object) Bytes() []int8 {
	return o.data.([]int8)
}
func (o *Object) Shorts() []int16 {
	return o.data.([]int16)
}
func (o *Object) Chars() []uint16 {
	return o.data.([]uint16)
}
func (o *Object) Ints() []int32 {
	return o.data.([]int32)
}
func (o *Object) Longs() []int64 {
	return o.data.([]int64)
}
func (o *Object) Floats() []float32 {
	return o.data.([]float32)
}

func (o *Object) Doubles() []float64 {
	return o.data.([]float64)
}
func (o *Object) Refs() []*Object {
	return o.data.([]*Object)
}
func (o *Object) ArrayLength() int32 {
	switch o.data.(type) {
	case []int8:
		return int32(len(o.Bytes()))
	case []int16:
		return int32(len(o.Shorts()))
	case []uint16:
		return int32(len(o.Chars()))
	case []int32:
		return int32(len(o.Ints()))
	case []float32:
		return int32(len(o.Floats()))
	case []float64:
		return int32(len(o.Doubles()))
	case []*Object:
		return int32(len(o.Refs()))
	default:
		panic("Not an array")
	}
}

func ArrayCopy(src, dst *Object, srcPos, dstPos, length int32) {
	switch src.data.(type) {
	case []int32:
		srcCopy := src.data.([]int32)[srcPos : srcPos+length]
		dstCopy := dst.data.([]int32)[dstPos : dstPos+length]
		copy(dstCopy, srcCopy)
	case []int8:
		srcCopy := src.data.([]int8)[srcPos : srcPos+length]
		dstCopy := dst.data.([]int8)[dstPos : dstPos+length]
		copy(dstCopy, srcCopy)
	case []int16:
		srcCopy := src.data.([]int16)[srcPos : srcPos+length]
		dstCopy := dst.data.([]int16)[dstPos : dstPos+length]
		copy(dstCopy, srcCopy)
	case []uint16:
		srcCopy := src.data.([]uint16)[srcPos : srcPos+length]
		dstCopy := dst.data.([]uint16)[dstPos : dstPos+length]
		copy(dstCopy, srcCopy)
	case []float32:
		srcCopy := src.data.([]float32)[srcPos : srcPos+length]
		dstCopy := dst.data.([]float32)[dstPos : dstPos+length]
		copy(dstCopy, srcCopy)
	case []float64:
		srcCopy := src.data.([]float64)[srcPos : srcPos+length]
		dstCopy := dst.data.([]float64)[dstPos : dstPos+length]
		copy(dstCopy, srcCopy)
	case []*Object:
		srcCopy := src.data.([]*Object)[srcPos : srcPos+length]
		dstCopy := dst.data.([]*Object)[dstPos : dstPos+length]
		copy(dstCopy, srcCopy)
	}

}
