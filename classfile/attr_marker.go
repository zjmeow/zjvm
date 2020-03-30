package classfile

type MarkerAttribute struct {
}

type DeprecatedAttribute struct {
	MarkerAttribute
}

type SyntheticAttribute struct {
	MarkerAttribute
}

func (ma *MarkerAttribute) readInfo(reader *ClassReader) {
}
