package heap

import "github.com/zjmeow/zjvm/classfile"

type Field struct {
	ClassMember
	constValueIndex uint
	slotId          uint
}

func newField(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField)
	}
	return fields
}

func (f *Field) copyAttributes(field *classfile.MemberInfo) {
	if valAttr := field.ConstantValueAttribute(); valAttr != nil {
		f.constValueIndex = uint(valAttr.ConstantValueIndex())
	}
}

func (f *Field) isLongOrDouble() bool {
	return f.descriptor == "J" || f.descriptor == "D"
}

func (f *Field) SlotId() uint {
	return f.slotId
}
