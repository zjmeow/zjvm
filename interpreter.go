package main

import (
	"fmt"
	"github.com/zjmeow/zjvm/instructions"
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
	"github.com/zjmeow/zjvm/rtda/heap"
	"reflect"
)

func interpret(method *heap.Method) {
	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)
	loop(thread, method.Code())
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
		fmt.Println(reflect.TypeOf(ins))
		ins.FetchOperands(reader)
		frame.SetNextPc(reader.Pc())
		ins.Execute(frame)
		fmt.Println(frame.OperandStack())
		fmt.Println(frame.LocalVars())
	}

}
