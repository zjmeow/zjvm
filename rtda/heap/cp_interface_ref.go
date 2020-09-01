package heap

import "github.com/zjmeow/zjvm/classfile"

type InterfaceRef struct {
	MemberRef
	method *Method
}

func newInterfaceRef(cp *ConstantPool, refInfo *classfile.ConstantInterfaceRefInfo) *InterfaceRef {
	ref := &InterfaceRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberRefInfo)
	return ref
}

func (r *InterfaceRef) ResolvedInterface() *Method {
	if r.method == nil {
		r.resolveInterfaceRef()
	}
	return r.method
}

func (r *InterfaceRef) resolveInterfaceRef() {
	// 被引用的类
	class := r.ResolveClass()
	// 接口的方法不能调用
	if !class.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	method := lookupInterfacesMethod(class, r.name, r.descriptor)
	// 找不到方法
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	// 没有访问权限
	if !method.isAccessibleTo(r.cp.class) {
		panic("java.lang.IllegalAccessError")
	}
	r.method = method
}
func lookupInterfacesMethod(iface *Class, name, descriptor string) *Method {
	for _, method := range iface.methods {
		if method.name == name && method.descriptor == descriptor {
			return method
		}
	}
	return LookupMethodInterfaces(iface.interfaces, name, descriptor)
}
