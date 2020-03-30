package classfile

type SourceFileAttribute struct {
	cp              ConstantPool
	sourceFileIndex uint16
}

func (sf *SourceFileAttribute) readInfo(reader *ClassReader) {
	sf.sourceFileIndex = reader.readUint16()
}
func (sf *SourceFileAttribute) FileName() string {
	return sf.cp.getUtf8(sf.sourceFileIndex)
}
