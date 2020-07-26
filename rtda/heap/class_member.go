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
