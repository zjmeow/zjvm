package classfile

type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (cv *ConstantValueAttribute) readInfo(reader *ClassReader) {
	cv.constantValueIndex = reader.readUint16()
}
func (cv *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return cv.constantValueIndex
}
