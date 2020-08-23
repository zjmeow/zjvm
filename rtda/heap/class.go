package heap

import (
	"github.com/zjmeow/zjvm/classfile"
	"strings"
)

type Class struct {
	classfile.AccessFlags
	name               string // thisClassName
	superClassName     string
	methods            []*Method
	interfaceNames     []string
	fields             []*Field
	interfaces         []*Class
	classLoader        *ClassLoader
	superClass         *Class
	staticVars         LocalVars
	constantPool       *ConstantPool
	instanceFieldCount uint
	staticFieldCount   uint
	instanceSlotCount  uint
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.AccessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newField(class, cf.Fields())
	class.methods = newMethod(class, cf.Methods())
	return class
}

func (c *Class) getPackageName() string {
	if i := strings.LastIndex(c.name, "/"); i >= 0 {
		return c.name[:i]
	}
	return ""
}
func (c *Class) isAccessibleTo(other *Class) bool {
	if c.IsPublic() {
		return true
	}
	return c.getPackageName() == other.getPackageName()
}
func (c *Class) isSubClass(other *Class) bool {
	for parent := c.superClass; parent != nil; parent = parent.superClass {
		if parent == other {
			return true
		}
	}
	return false
}
func (c *Class) ConstantPool() *ConstantPool {
	return c.constantPool
}

func (c *Class) NewObject() *Object {
	return newObject(c)
}
func (c *Class) StaticVars() LocalVars {
	return c.staticVars
}

func (c *Class) isImplements(iface *Class) bool {
	for class := c; class != nil; class = class.superClass {
		for _, i := range class.interfaces {
			if i == iface || i.isSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}

func (c *Class) isSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range c.interfaces {
		if superInterface == iface || superInterface.isSubInterfaceOf(iface) {
			return true
		}
	}
	return false
}

func (c *Class) isAssignableFrom(otherClass *Class) bool {
	if c == otherClass {
		return true
	}
	if !c.IsInterface() {
		return otherClass.isSubClass(c)
	} else {
		return otherClass.isImplements(c)
	}
}

func (c *Class) GetMainMethod() *Method {
	return c.getStaticMethod("main", "([Ljava/lang/String;)V")
}

func (c *Class) getStaticMethod(name, descriptor string) *Method {
	for _, m := range c.methods {
		if m.Name() == name && m.Descriptor() == descriptor {
			return m
		}
	}
	return nil
}
