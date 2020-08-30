package heap

import "github.com/zjmeow/zjvm/classfile"

type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *ConstantPool, refInfo *classfile.ConstantMethodRefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberRefInfo)
	return ref
}

func (m *MethodRef) ResolvedMethod() *Method {
	if m.method == nil {
		m.resolveMethodRef()
	}
	return m.method
}

//
func (m *MethodRef) resolveMethodRef() {
	// 被引用的类
	class := m.ResolveClass()
	// 接口的方法不能调用
	if class.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	method := lookupMethod(class, m.name, m.descriptor)
	// 找不到方法
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	// 没有访问权限
	if !method.isAccessibleTo(m.cp.class) {
		panic("java.lang.IllegalAccessError")
	}
	m.method = method
}

func lookupMethod(class *Class, name, descriptor string) *Method {
	method := LookupMethodInClass(class, name, descriptor)
	if method == nil {
		method = LookupMethodInterfaces(class.interfaces, name, descriptor)
	}
	return method
}
