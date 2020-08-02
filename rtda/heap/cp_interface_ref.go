package heap

import "github.com/zjmeow/zjvm/classfile"

type InterfaceRef struct {
	MemberRef
	method *Method
}

func newInterfaceMethodRef(cp *ConstantPool, refInfo *classfile.ConstantInterfaceRefInfo) *InterfaceRef {
	ref := &InterfaceRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberRefInfo)
	return ref
}
