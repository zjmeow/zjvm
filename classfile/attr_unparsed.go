package classfile

type UnparsedAttribute struct {
	name   string
	length uint32
	info   []byte
}

func (a *UnparsedAttribute) readInfo(reader *ClassReader) {
	a.info = reader.readBytes(a.length)
}
