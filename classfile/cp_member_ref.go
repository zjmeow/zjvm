package classfile

type ConstantMemberRefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (ci *ConstantMemberRefInfo) readInfo(reader *ClassReader) {
	ci.classIndex = reader.readUint16()
	ci.nameAndTypeIndex = reader.readUint16()
}

func (ci *ConstantMemberRefInfo) ClassName() string {
	return ci.cp.getClassName(ci.classIndex)
}
func (ci *ConstantMemberRefInfo) NameAndDescriptor() (string, string) {
	return ci.cp.getNameAndType(ci.nameAndTypeIndex)
}

type ConstantFieldRefInfo struct {
	ConstantMemberRefInfo
}
type ConstantMethodRefInfo struct {
	ConstantMemberRefInfo
}
type ConstantInterfaceRefInfo struct {
	ConstantMemberRefInfo
}
