package base

type BytecodeReader struct {
	code []byte
	pc   int
}

func (br *BytecodeReader) Reset(code []byte, pc int) {
	br.code = code
	br.pc = pc
}

func (br *BytecodeReader) ReadUint8() uint8 {
	res := br.code[br.pc]
	br.pc++
	return res
}

func (br *BytecodeReader) ReadInt8() int8 {
	return int8(br.ReadUint8())
}

func (br *BytecodeReader) ReadUint16() uint16 {
	high := uint16(br.ReadUint8())
	low := uint16(br.ReadUint8())
	return (high << 8) | low
}

func (br *BytecodeReader) ReadInt16() int16 {
	return int16(br.ReadUint16())
}

func (br *BytecodeReader) ReadInt32() int32 {
	high := int32(br.ReadUint16())
	low := int32(br.ReadUint16())
	return (high << 16) | low
}

func (br *BytecodeReader) ReadInt32s(count int32) []int32 {
	res := make([]int32, count)
	for i := range res {
		res[i] = br.ReadInt32()
	}
	return res
}

func (br *BytecodeReader) SkipPadding() {
	for br.pc%4 != 0 {
		br.ReadUint8()
	}
}
func (br *BytecodeReader) Pc() int {
	return br.pc
}
