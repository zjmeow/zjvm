package heap

type SymRef struct {
	cp        *ConstantPool
	className string
	class     *Class
}

func (f *SymRef) ResolveClass() *Class {
	if f.class == nil {
		f.resolveClassRef()
	}
	return f.class
}

// 用 c1 的加载器去加载c2，如果不可访问则丢出错误
func (f *SymRef) resolveClassRef() {
	c1 := f.cp.class
	c2 := c1.classLoader.LoadClass(f.className)
	if !c2.isAccessibleTo(c1) {
		panic("java.lang.IllegalAccessError")
	}
	f.class = c2
}
