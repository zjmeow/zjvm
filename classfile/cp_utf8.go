package classfile

type ConstantUtf8Info struct {
	val string
}

// TODO:replace utf8 with MUTF-8
func (ci *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	ci.val = string(bytes)
}
