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

func (ci *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	ci.val = int32(bytes)
}
func (ci *ConstantIntegerInfo) Value() int32 {
	return ci.val
}

func (ci *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	ci.val = int64(bytes)
}

func (ci *ConstantLongInfo) Value() int64 {
	return ci.val
}

func (ci *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	ci.val = math.Float32frombits(bytes)
}

func (ci *ConstantFloatInfo) Value() float32 {
	return ci.val
}

func (ci *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	ci.val = math.Float64frombits(bytes)
}

func (ci *ConstantDoubleInfo) Value() float64 {
	return ci.val
}
