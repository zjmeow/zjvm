package heap

func (o *Object) Clone() *Object {
	return &Object{
		class: o.class,
		data:  o.cloneData(),
	}
}

// TODO
func (o *Object) cloneData() interface{} {
	return nil
}
