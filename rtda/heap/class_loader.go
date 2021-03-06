package heap

import (
	"fmt"
	"github.com/zjmeow/zjvm/classfile"
	"github.com/zjmeow/zjvm/classpath"
	"strings"
)

// 类加载的过程：加载→连接→初始化
type ClassLoader struct {
	cp       *classpath.ClassPath
	classMap map[string]*Class
}

func NewClassLoader(cp *classpath.ClassPath) *ClassLoader {
	loader := &ClassLoader{
		cp:       cp,
		classMap: make(map[string]*Class),
	}
	loader.loadBasicClass()

	return loader
}
func (cl *ClassLoader) loadBasicClass() {
	jlClass := cl.LoadClass("java/lang/Class")
	for _, class := range cl.classMap {
		if class.jClass == nil {
			class.jClass = jlClass.NewObject()
			class.jClass.SetExtra(class)
		}
	}
}
func (cl *ClassLoader) loadPrimitiveClasses() {
	for primitiveType, _ := range primitiveTypes {
		cl.loadPrimitiveClass(primitiveType)
	}
}

func (cl *ClassLoader) loadPrimitiveClass(className string) {
	class := &Class{
		AccessFlags: classfile.AccPublic,
		name:        className,
		classLoader: cl,
		initStarted: true,
	}
	class.jClass = cl.classMap["java/lang/Class"].NewObject()
	class.jClass.SetExtra(cl)
	cl.classMap[className] = class
}

func (cl *ClassLoader) LoadClass(name string) *Class {
	if class, ok := cl.classMap[name]; ok {
		return class
	}
	var classRes *Class
	// 如果是数组
	if strings.HasPrefix(name, "[") {
		classRes = cl.loadArrayClass(name)
	} else {
		classRes = cl.loadNonArrayClass(name)
	}
	// 给类设置个关联对象
	if jlClass, ok := cl.classMap["java/lang/Class"]; ok {
		classRes.jClass = jlClass.NewObject()
		classRes.jClass.SetExtra(classRes)
	}
	return classRes

}

// 数组类初始化比较简单创建好类放进map中即可
func (cl *ClassLoader) loadArrayClass(name string) *Class {
	class := &Class{
		AccessFlags: classfile.AccPublic,
		name:        name,
		classLoader: cl,
		initStarted: true,
		superClass:  cl.LoadClass("java/lang/Object"),
		interfaces: []*Class{
			cl.LoadClass("java/lang/Cloneable"),
			cl.LoadClass("java/io/Serializable"),
		},
	}
	cl.classMap[name] = class
	return class
}

func (cl *ClassLoader) loadNonArrayClass(name string) *Class {
	data, entry := cl.readClass(name)
	class := cl.defineClass(data)
	linkClass(class)
	fmt.Printf("[Loaded %s from %s]\n", name, entry)
	return class
}

func (cl *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := cl.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + name)
	}
	return data, entry
}

func (cl *ClassLoader) defineClass(data []byte) *Class {
	class := cl.parseClass(data)
	class.classLoader = cl
	resolveInterfaces(class)
	resolveSuperClass(class)
	cl.classMap[class.name] = class
	return class
}

func (cl *ClassLoader) parseClass(data []byte) *Class {
	cf, err := classfile.Parse(data)
	if err != nil {
		panic("java.lang.ClassFormatError")
	}
	return newClass(cf)
}

func resolveSuperClass(class *Class) {
	// 非object类则需要递归加载父类
	if class.name != "java/lang/Object" {
		class.superClass = class.classLoader.LoadClass(class.superClassName)
	}
}

// 有实现接口则需要递归加载全部的接口
func resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount <= 0 {
		return
	}
	interfaces := make([]*Class, interfaceCount)
	for i, name := range class.interfaceNames {
		interfaces[i] = class.classLoader.LoadClass(name)
	}
}

// 连接 包括验证和准备阶段
func linkClass(class *Class) {
	verify(class)
	prepare(class)
}

// 验证过程，主要验证二进制流是否符合虚拟机规范，这里不实现
func verify(class *Class) {

}

// 准备阶段，主要是为域分配空间
func prepare(class *Class) {
	calcInstanceFieldSlotIds(class)
	calcStaticFieldSlotIds(class)
	allocAndInitStaticVars(class)
}

func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}
	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.instanceSlotCount = slotId
}
func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	if class.superClass != nil {
		slotId = class.superClass.staticFieldCount
	}
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.staticFieldCount = slotId
}
func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticFieldCount)
	// 初始化常量
	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field)
		}
	}

}

// 从常量池中拿数据初始化到final域上
func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.constValueIndex
	slotId := field.slotId
	if cpIndex > 0 {
		switch field.descriptor {
		case "Z", "B", "C", "S", "I":
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId, val)
		case "J":
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId, val)
		case "F":
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId, val)
		case "D":
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId, val)
		case "Ljava/lang/String;":
			goStr := cp.GetConstant(cpIndex).(string)
			jStr := JString(class.classLoader, goStr)
			vars.SetRef(slotId, jStr)
		default:
			panic("can not init type")
		}
	}
}
