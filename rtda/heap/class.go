package heap

import (
	"github.com/zjmeow/zjvm/classfile"
)

type Class struct {
	classfile.AccessFlags
	name               string // thisClassName
	superClassName     string
	interfaceNames     []string
	fields             []*Field
	classLoader        *ClassLoader
	superClass         *Class
	staticVars         LocalVars
	constantPool       *ConstantPool
	instanceFieldCount uint
	staticFieldCount   uint
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.AccessFlags = cf.AccessFlags()
	//class.name = cf.ClassName()
	// todo class name未实现

	return class
}
