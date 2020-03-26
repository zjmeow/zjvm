package classfile

type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) {}

func (cp *ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := (*cp)[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index")
}

func (cp *ConstantPool) getUtf8(index uint16) string {
	return ""
}
