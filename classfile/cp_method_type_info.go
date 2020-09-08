package classfile

type ConstantMethodTypeInfo struct {
	DescriptorIndex uint16
}

func (ci *ConstantMethodTypeInfo) readInfo(reader *ClassReader) {
	ci.DescriptorIndex = reader.readUint16()

}
