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

func (o *Object) Fields() LocalVars {
	return o.fields
}
func (o *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(o.class)
}
