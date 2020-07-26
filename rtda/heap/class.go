package heap

import "github.com/zjmeow/zjvm/classfile"

type Class struct {
	classfile.AccessFlags
	Name           string // thisClassName
	superClassName string
	interfaceNames []string
}
