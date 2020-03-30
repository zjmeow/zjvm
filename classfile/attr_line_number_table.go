package classfile

type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}
type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (ln *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	count := reader.readUint16()
	ln.lineNumberTable = make([]*LineNumberTableEntry, count)
	for i := range ln.lineNumberTable {
		ln.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}

}
