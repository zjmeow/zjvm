package heap

import "github.com/zjmeow/zjvm/classfile"

type MemberRef struct {
	SymRef
	name       string
	descriptor string
}

func (mr *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMemberRefInfo) {
	mr.className = refInfo.ClassName()
	mr.name, mr.descriptor = refInfo.NameAndDescriptor()
}
func (mr *MemberRef) Name() string {
	return mr.name
}
