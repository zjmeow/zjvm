package main

import (
	"fmt"
	"github.com/zjmeow/zjvm/classfile"
	"github.com/zjmeow/zjvm/instructions"
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
)

func interpret(methodInfo *classfile.MemberInfo) {
	codeAttr := methodInfo.CodeAttribute()
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	code := codeAttr.Code()
	thread := rtda.NewThread()
	frame := thread.NewFrame(maxLocals, maxStack)
	thread.
}

func loop(thread *rtda.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}
	for {
		pc := frame.NextPc()
		thread.SetPc(pc)
		reader.Reset(bytecode, pc)
		opCode := reader.ReadUint8()
		ins := instructions.NewInstruction(opCode)
		ins.FetchOperands(reader)
		frame.SetNextPc(reader.Pc())
		fmt.Println()
		ins.Execute(frame)
	}

}
