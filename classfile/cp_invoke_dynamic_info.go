package classfile

type ConstantInvokeDynamicInfo struct {
	BootstrapMethodAttrIndex uint16
	NameAndTypeIndex         uint16
}

func (ci *ConstantInvokeDynamicInfo) readInfo(reader *ClassReader) {
	ci.BootstrapMethodAttrIndex = reader.readUint16()
	ci.NameAndTypeIndex = reader.readUint16()
}
