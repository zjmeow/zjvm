package heap

import "github.com/zjmeow/zjvm/classfile"

type ClassMember struct {
	classfile.AccessFlags
	name       string
	descriptor string
	class      *Class
}

func (c *ClassMember) copyMemberInfo(info *classfile.MemberInfo) {
	c.AccessFlags = info.AccessFlags()
	c.name = info.Name()
	c.descriptor = info.Descriptor()
}

func (c *ClassMember) isAccessibleTo(class *Class) bool {
	if class.IsPublic() {
		return true
	}
	this := c.class
	if c.IsProtected() {
		return class == this || class.IsSubClass(this)
	}
	if !c.IsPrivate() {
		return this.GetPackageName() == class.GetPackageName()
	}
	return this == class
}

func (c *ClassMember) Class() *Class {
	return c.class
}
func (c *ClassMember) Name() string {
	return c.name
}
func (c *ClassMember) Descriptor() string {
	return c.descriptor
}
