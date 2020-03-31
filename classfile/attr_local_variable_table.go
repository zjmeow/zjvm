package classfile

type LocalVariableTableAttribute struct {
	localVariableTable []*LocalVariableTableEntry
}

type LocalVariableTableEntry struct {
	StartPc         uint16
	Length          uint16
	NameIndex       uint16
	DescriptorIndex uint16
	Index           uint16
}

func (lv *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	count := reader.readUint16()
	lv.localVariableTable = make([]*LocalVariableTableEntry, count)
	for i := range lv.localVariableTable {
		lv.localVariableTable[i] = &LocalVariableTableEntry{
			StartPc:         reader.readUint16(),
			Length:          reader.readUint16(),
			NameIndex:       reader.readUint16(),
			DescriptorIndex: reader.readUint16(),
			Index:           reader.readUint16(),
		}
	}
}
