package heap

type Object struct {
	class *Class
	data  interface{}
	extra interface{}
}

func newObject(class *Class) *Object {
	return &Object{
		class: class,
		data:  newSlots(class.instanceSlotCount),
	}
}

func NewDefaultObject(class *Class, data interface{}) *Object {
	return &Object{
		class: class,
		data:  data,
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
	o.extra = extra
}
func (o *Object) Extra() interface{} {
	return o.extra
}
func (o *Object) GetRefVar(name, descriptor string) *Object {
	field := o.class.GetField(name, descriptor, false)
	slots := o.data.(LocalVars)
	return slots.GetRef(field.slotId)
}
func (o *Object) Data() interface{} {
	return o.data
}
