package classfile

type ExceptionAttribute struct {
	exceptionIndexTable []uint16
}

func (ea *ExceptionAttribute) readInfo(reader *ClassReader) {
	ea.exceptionIndexTable = reader.readUint16s()
}

func (ea *ExceptionAttribute) ExceptionIndexTable() []uint16 {
	return ea.exceptionIndexTable
}
