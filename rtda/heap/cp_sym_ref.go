package heap

type SymRef struct {
	cp        *ConstantPool
	className string
	class     *Class
}

func (f *SymRef) resolveClass() *Class {
	if f.class == nil {
		f.resolveClassRef()
	}
	return f.class
}

func (f *SymRef) resolveClassRef() {
	c1 := f.cp.class
	c2 := c1.classLoader.LoadClass(f.className)
	if c2.isAccessibleTo(c1) {
		panic("java.lang.IllegalAccessError")
	}
	f.class = c2
}
