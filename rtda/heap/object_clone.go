package heap

import "reflect"

func (o *Object) Clone() *Object {
	return &Object{
		class: o.class,
		data:  o.cloneData(),
	}
}

func (o *Object) cloneData() interface{} {
	fields1 := reflect.ValueOf(o.Fields)
	fields2 := reflect.MakeSlice(fields1.Type(), fields1.Len(), fields1.Len())
	reflect.Copy(fields2, fields1)
	return NewDefaultObject(o.Class(), fields2.Interface())

}
