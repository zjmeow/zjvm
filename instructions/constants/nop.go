package constants

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
)

type Nop struct {
	base.NoOperandsInstruction
}

func (n *Nop) Execute(frame *rtda.Frame) {

}
