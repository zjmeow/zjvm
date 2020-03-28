package classfile

type ConstantClassInfo struct {
	cp    ConstantPool
	index uint16
}

func (ci *ConstantClassInfo) readInfo(reader *ClassReader) {
	ci.index = reader.readUint16()
}

func (ci *ConstantClassInfo) Name() string {
	return ci.cp.getUtf8(ci.index)
}
