package classfile

type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}
type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func (ca *CodeAttribute) readInfo(reader *ClassReader) {
	ca.maxStack = reader.readUint16()
	ca.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	ca.code = reader.readBytes(codeLength)
	ca.exceptionTable = readExceptionTable(reader)
	ca.attributes = readAttributes(reader, ca.cp)
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	count := reader.readUint16()
	exceptions := make([]*ExceptionTableEntry, count)
	for i := range exceptions {
		exceptions[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}
	return exceptions
}

func (ca *CodeAttribute) MaxStack() uint16 {
	return ca.maxStack
}
func (ca *CodeAttribute) MaxLocals() uint16 {
	return ca.maxLocals
}

func (ca *CodeAttribute) Code() []byte {
	return ca.code
}

func (ca *CodeAttribute) ConstantPool() ConstantPool {
	return ca.cp
}
func (ca *CodeAttribute) ExceptionTableEntry() []*ExceptionTableEntry {
	return ca.exceptionTable
}

func (e *ExceptionTableEntry) StartPc() uint16 {
	return e.startPc
}
func (e *ExceptionTableEntry) EndPc() uint16 {
	return e.endPc
}
func (e *ExceptionTableEntry) HandlerPc() uint16 {
	return e.handlerPc
}
func (e *ExceptionTableEntry) CatchType() uint16 {
	return e.catchType
}
