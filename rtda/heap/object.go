package heap

type Object struct {
	class *Class
	data  interface{}
}

func newObject(class *Class) *Object {
	return &Object{
		class: class,
		data:  newSlots(class.instanceSlotCount),
	}
}

func (o *Object) Fields() LocalVars {
	return o.data.(LocalVars)
}
func (o *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(o.class)
}
func (o *Object) Class() *Class {
	return o.class
}
func (o *Object) SetRefVar(name, descriptor string, ref *Object) {
	field := o.class.GetField(name, descriptor, false)
	o.Fields().SetRef(field.slotId, ref)
}
func (o *Object) SetExtra(extra interface{}) {

}
