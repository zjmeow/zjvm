package classfile

type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	count := int(reader.readUint16())
	cp := make([]ConstantInfo, count)
	for i := 1; i < count; i++ {
		cp[i] = readConstantInfo(reader, cp)
	}
	return cp
}

func (cp *ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := (*cp)[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index")
}

func (cp *ConstantPool) getUtf8(index uint16) string {
	return ""
}
