package lang

import (
	"github.com/zjmeow/zjvm/native"
	"github.com/zjmeow/zjvm/rtda"
	"math"
)

func init() {
	native.Register("java/lang/Double", "doubleToRawLongBits", "(D)J", doubleToRawLongBits)
}
func doubleToRawLongBits(frame *rtda.Frame) {
	value := frame.LocalVars().GetDouble(0)
	bits := math.Float64bits(value)
	frame.OperandStack().PushLong(int64(bits))
}
