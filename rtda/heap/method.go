package heap

import "github.com/zjmeow/zjvm/classfile"

type Method struct {
	ClassMember
	maxStack  uint16
	maxLocals uint16
	code      []byte
}

func newMethod(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, method := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(method)
		methods[i].copyAttributes(method)
	}
	return methods
}

func (m *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		m.maxStack = codeAttr.MaxStack()
		m.maxLocals = codeAttr.MaxLocals()
		m.code = codeAttr.Code()
	}
}
func (m *Method) MaxStack() uint16 {
	return m.maxStack
}
func (m *Method) MaxLocals() uint16 {
	return m.maxLocals
}
func (m *Method) Code() []byte {
	return m.code
}
