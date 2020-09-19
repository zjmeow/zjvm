package io

import (
	"github.com/zjmeow/zjvm/native"
	"github.com/zjmeow/zjvm/rtda"
	"os"
	"unsafe"
)

func init() {
	native.Register("java/io/FileOutputStream",
		"writeBytes",
		"(LBIIZ)V", writeBytes)
}
func writeBytes(frame *rtda.Frame) {
	vars := frame.LocalVars()
	//this := vars.GetRef(0)
	b := vars.GetRef(1)
	off := vars.GetInt(2)
	len := vars.GetInt(3)
	//append := vars.GetBoolean(4)
	jBytes := b.Data().([]int8)
	goBytes := castInt8sToUint8s(jBytes)
	goBytes = goBytes[off : off+len]
	os.Stdout.Write(goBytes)
}
func castInt8sToUint8s(jBytes []int8) []byte {
	ptr := unsafe.Pointer(&jBytes)
	goBytes := *((*[]byte)(ptr))
	return goBytes
}
