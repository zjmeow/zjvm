package heap

import "github.com/zjmeow/zjvm/classfile"

type FieldRef struct {
	MemberRef
	field *Field
}

func newFieldRef(cp *ConstantPool, refInfo *classfile.ConstantFieldRefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberRefInfo)
	return ref
}

func (f *FieldRef) resolveField() *Field {
	if f.field == nil {
		f.resolveFieldRef()
	}
	return f.field
}

func (f *FieldRef) resolveFieldRef() {
	d := f.cp.class
	c := f.ResolveClass()
	field := lookupField(c, f.name, f.descriptor)
	if field == nil {
		panic("java.lang.NoSuchAccessError")
	}
	if !field.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	f.field = field
}

func lookupField(class *Class, name, descriptor string) *Field {
	for _, field := range class.fields {
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}
	for _, inter := range class.interfaces {
		if field := lookupField(inter, name, descriptor); field != nil {
			return field
		}
	}
	if class.superClass != nil {
		return lookupField(class.superClass, name, descriptor)
	}
	return nil
}
