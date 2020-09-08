package classfile

type ConstantMethodHandleInfo struct {
	BootstrapMethodAttrIndex uint8
	NameAndTypeIndex         uint16
}

func (ci *ConstantMethodHandleInfo) readInfo(reader *ClassReader) {
	ci.BootstrapMethodAttrIndex = reader.readUint8()
	ci.NameAndTypeIndex = reader.readUint16()
}
