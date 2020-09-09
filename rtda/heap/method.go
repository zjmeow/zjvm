package heap

import (
	"github.com/zjmeow/zjvm/classfile"
	"github.com/zjmeow/zjvm/instructions"
)

type Method struct {
	ClassMember
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	argSlotCount   uint
	exceptionTable classfile.ExceptionTable
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, method := range cfMethods {
		methods[i] = newMethod(class, method)
	}
	return methods
}

func newMethod(class *Class, cfMethod *classfile.MemberInfo) *Method {
	method := &Method{}
	method.class = class
	method.copyMemberInfo(cfMethod)
	method.copyAttributes(cfMethod)
	md := parseMethodDescriptor(method.descriptor)
	method.calcArgSlotCount(md)
	if method.IsNative() {
		method.injectCodeAttribute(md.returnType)
	}
	return method
}

func newExceptionTable(entries []*classfile.ExceptionTableEntry, cp *ConstantPool) ExceptionTable {
	table := make([]*ExceptionHandler, len(entries))
	for i, entry := range entries {
		table[i] = &ExceptionHandler{
			startPc:   int(entry.StartPc()),
			endPc:     int(entry.EndPc()),
			handlerPc: int(entry.HandlerPc()),
			catchType: getCatchType(uint(entry.CatchType()), cp),
		}
	}
	return table
}

func getCatchType(index uint, cp *ConstantPool) *ClassRef {
	if index == 0 {
		return nil
	}
	return cp.GetConstant(index).(*ClassRef)
}

func (m *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		m.maxStack = codeAttr.MaxStack()
		m.maxLocals = codeAttr.MaxLocals()
		m.code = codeAttr.Code()
		m.exceptionTable = newExceptionTable(codeAttr.E)
	}
}
func (m *Method) MaxStack() uint16 {
	return m.maxStack
}
func (m *Method) MaxLocals() uint16 {
	return m.maxLocals
}
func (m *Method) Code() []byte {
	return m.code
}

func (m *Method) ArgSlotCount() uint {
	return m.argSlotCount
}
func (m *Method) calcArgSlotCount(parsedDescriptor MethodDescriptor) {

	for _, paramType := range parsedDescriptor.parameterTypes {
		m.argSlotCount++
		if paramType == "J" || paramType == "D" {
			m.argSlotCount++
		}
	}
	if !m.IsStatic() {
		m.argSlotCount++
	}
}

// 手动注入自定义的native invoke 指令调用native方法
func (m *Method) injectCodeAttribute(returnType string) {
	m.maxStack = 4
	m.maxLocals = uint16(m.argSlotCount)
	switch returnType {
	case "V":
		m.code = []byte{instructions.OpInvokeNative, instructions.OpReturn}
	case "D":
		m.code = []byte{instructions.OpInvokeNative, instructions.OpDReturn}
	case "F":
		m.code = []byte{instructions.OpInvokeNative, instructions.OpFReturn}
	case "J":
		m.code = []byte{instructions.OpInvokeNative, instructions.OpLReturn}
	case "L", "[":
		m.code = []byte{instructions.OpInvokeNative, instructions.OpAReturn}
	default:
		m.code = []byte{instructions.OpInvokeNative, instructions.OpIReturn}
	}

}
