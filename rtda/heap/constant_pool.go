package heap

import (
	"fmt"
	"github.com/zjmeow/zjvm/classfile"
)

type Constant interface {
}
type ConstantPool struct {
	class     *Class
	constants []Constant
}

func (c *ConstantPool) GetConstant(index uint) Constant {
	if constant := c.constants[index]; constant != nil {
		return constant
	}
	panic(fmt.Sprintf("No contants at index %d", index))
}

func newConstantPool(class *Class, cfCp classfile.ConstantPool) *ConstantPool {
	cpCount := len(cfCp)
	constants := make([]Constant, cpCount)
	rtCp := &ConstantPool{class, constants}
	for i := 1; i < cpCount; i++ {
		cpInfo := cfCp[i]
		switch cpInfo.(type) {
		case *classfile.ConstantIntegerInfo:
			intInfo := cpInfo.(*classfile.ConstantIntegerInfo)
			constants[i] = intInfo.Value()
		case *classfile.ConstantLongInfo:
			longInfo := cpInfo.(*classfile.ConstantLongInfo)
			constants[i] = longInfo.Value()
			i++
		case *classfile.ConstantFloatInfo:
			floatInfo := cpInfo.(*classfile.ConstantFloatInfo)
			constants[i] = floatInfo.Value()
		case *classfile.ConstantDoubleInfo:
			doubleInfo := cpInfo.(*classfile.ConstantDoubleInfo)
			constants[i] = doubleInfo.Value()
			i++
		case *classfile.ConstantStringInfo:
			stringInfo := cpInfo.(*classfile.ConstantStringInfo)
			constants[i] = stringInfo.String()
		case *classfile.ConstantUtf8Info:
			utf8Info := cpInfo.(*classfile.ConstantUtf8Info)
			constants[i] = utf8Info.String()
		case *classfile.ConstantClassInfo:
			classInfo := cpInfo.(*classfile.ConstantClassInfo)
			constants[i] = newClassRef(rtCp, classInfo)
		case *classfile.ConstantFieldRefInfo:
			fieldInfo := cpInfo.(*classfile.ConstantFieldRefInfo)
			constants[i] = newFieldRef(rtCp, fieldInfo)
		case *classfile.ConstantMethodRefInfo:
			methodInfo := cpInfo.(*classfile.ConstantMethodRefInfo)
			constants[i] = newMethodRef(rtCp, methodInfo)
		case *classfile.ConstantInterfaceRefInfo:
			interfaceInfo := cpInfo.(*classfile.ConstantInterfaceRefInfo)
			constants[i] = newInterfaceMethodRef(rtCp, interfaceInfo)
		case *classfile.ConstantNameAndTypeInfo:
			nameAndType := cpInfo.(*classfile.ConstantNameAndTypeInfo)
			constants[i] = nameAndType
		}

	}

	return rtCp
}
