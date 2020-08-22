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
	Offset int
}

type Index8Instruction struct {
	Index uint
}
type Index16Instruction struct {
	Index uint
}

func (np *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {

}

func (bi *BranchInstruction) Execute(frame *rtda.Frame) {

}
func (bi *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	bi.Offset = int(reader.ReadInt16())
}

func (i *Index8Instruction) Execute(frame *rtda.Frame) {

}
func (i *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	i.Index = uint(reader.ReadUint8())
}
func (i *Index16Instruction) Execute(frame *rtda.Frame) {

}
func (i *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	i.Index = uint(reader.ReadUint16())
}
