package classfile

type ConstantStringInfo struct {
	cp    ConstantPool
	index uint16
}

func (ci *ConstantStringInfo) readInfo(reader *ClassReader) {
	ci.index = reader.readUint16()
}
func (ci *ConstantStringInfo) String() string {
	return ci.cp.getUtf8(ci.index)
}
