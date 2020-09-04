package classfile

type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	count := int(reader.readUint16())
	cp := make([]ConstantInfo, count)
	for i := 1; i < count; i++ {
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
	return cp
}

func (cp *ConstantPool) getClassName(index uint16) string {
	classInfo := cp.getConstantInfo(index).(*ConstantClassInfo)
	return cp.getUtf8(classInfo.index)
}

func (cp *ConstantPool) getNameAndType(index uint16) (string, string) {
	info := cp.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := cp.getUtf8(info.nameIndex)
	tp := cp.getUtf8(info.descriptorIndex)
	return name, tp
}

func (cp *ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := (*cp)[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index")
}

func (cp *ConstantPool) getUtf8(index uint16) string {
	return cp.getConstantInfo(index).(*ConstantUtf8Info).val
}
