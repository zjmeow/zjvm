package classfile

const (
	CONSTANT_Class              = 7
	CONSTANT_Fieldref           = 9
	CONSTANT_Methodref          = 10
	CONSTANT_InterfaceMethodref = 11
	CONSTANT_String             = 8
	CONSTANT_Integer            = 3
	CONSTANT_Float              = 4
	CONSTANT_Long               = 5
	CONSTANT_Double             = 6
	CONSTANT_NameAndType        = 12
	CONSTANT_Utf8               = 1
	CONSTANT_MethodHandle       = 15
	CONSTANT_MethodType         = 16
	CONSTANT_InvokeDynamic      = 18
)

type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

func newTagInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_Integer:
		return &ConstantIntegerInfo{}
	case CONSTANT_Long:
		return &ConstantLongInfo{}
	case CONSTANT_Float:
		return &ConstantFloatInfo{}
	case CONSTANT_Double:
		return &ConstantDoubleInfo{}
	case CONSTANT_Utf8:
		return &ConstantUtf8Info{}
	case CONSTANT_Class:
		return &ConstantClassInfo{cp: cp}
	case CONSTANT_String:
		return &ConstantStringInfo{cp: cp}
	case CONSTANT_Fieldref:
		return &ConstantFieldRefInfo{ConstantMemberRefInfo{cp: cp}}
	case CONSTANT_Methodref:
		return &ConstantMethodRefInfo{ConstantMemberRefInfo{cp: cp}}
	case CONSTANT_InterfaceMethodref:
		return &ConstantInterfaceRefInfo{ConstantMemberRefInfo{cp: cp}}
	case CONSTANT_NameAndType:
		return &ConstantNameAndTypeInfo{}
	case CONSTANT_InvokeDynamic:
		return &ConstantInvokeDynamicInfo{}
	case CONSTANT_MethodHandle:
		return &ConstantMethodHandleInfo{}
	case CONSTANT_MethodType:
		return &ConstantMethodTypeInfo{}
	}
	panic("java.lang.ClassFormatError: constant pool tag")
	//return nil
}

func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	c := newTagInfo(tag, cp)
	c.readInfo(reader)
	return c
}
