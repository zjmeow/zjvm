package heap

type Object struct {
	class  *Class
	fields LocalVars
}

func newObject(class *Class) *Object {
	return &Object{
		class:  class,
		fields: newSlots(class.instanceSlotCount),
	}
}
