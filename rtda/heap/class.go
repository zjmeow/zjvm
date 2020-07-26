package heap

import "github.com/zjmeow/zjvm/classfile"

type Class struct {
	AccessFlags    classfile.AccessFlags
	Name           string // thisClassName
	superClassName string
	interfaceNames []string
}
