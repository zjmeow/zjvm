package heap

type Class struct {
	AccessFlags    uint16
	Name           string // thisClassName
	superClassName string
	interfaceNames []string
}
