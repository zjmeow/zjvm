package classfile

import "math"

type ConstantIntegerInfo struct {
	val int32
}
type ConstantLongInfo struct {
	val int64
}

type ConstantFloatInfo struct {
	val float32
}

type ConstantDoubleInfo struct {
	val float64
}

type ConstantUtf8Info struct {
	val string
}

func (ci *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	ci.val = int32(bytes)
}

func (ci *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	ci.val = int64(bytes)
}

func (ci *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	ci.val = math.Float32frombits(bytes)
}

func (ci *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	ci.val = math.Float64frombits(bytes)
}

// TODO:replace utf8 with MUTF-8
func (ci *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	ci.val = string(bytes)
}
