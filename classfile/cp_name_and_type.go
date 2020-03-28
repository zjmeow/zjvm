package classfile

type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (ci *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	ci.nameIndex = reader.readUint16()
	ci.descriptorIndex = reader.readUint16()
}
