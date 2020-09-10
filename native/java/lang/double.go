package lang

import (
	"github.com/zjmeow/zjvm/native"
	"github.com/zjmeow/zjvm/rtda"
	"math"
)

func init() {
	native.Register("java/lang/Double", "doubleToRawIntBits", "(F)I", doubleToRawIntBits)
}
func doubleToRawIntBits(frame *rtda.Frame) {
	value := frame.LocalVars().GetDouble(0)
	bits := math.Float64bits(value)
	frame.OperandStack().PushLong(int64(bits))
}
