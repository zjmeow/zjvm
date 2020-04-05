package base

import (
	"github.com/zjmeow/zjvm/rtda"
)

type Instruction interface {
	Execute(frame *rtda.Frame)
	FetchOperands(reader *BytecodeReader)
}

type NoOperandsInstruction struct {
}
type BranchInstruction struct {
	offset int
}

type Index8Instruction struct {
	index uint
}
type Index16Instruction struct {
	index uint
}

func (np *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {

}

func (bi *BranchInstruction) Execute(frame *rtda.Frame) {

}
func (bi *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	bi.offset = int(reader.ReadInt16())
}

func (i *Index8Instruction) Execute(frame *rtda.Frame) {

}
func (i *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	i.index = uint(reader.ReadUint8())
}
func (i *Index16Instruction) Execute(frame *rtda.Frame) {

}
func (i *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	i.index = uint(reader.ReadUint16())
}
