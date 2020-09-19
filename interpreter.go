package main

import (
	"fmt"
	"github.com/zjmeow/zjvm/instructions"
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
	"reflect"
)

func interpret(thread *rtda.Thread) {
	defer catchErr(thread)
	loop(thread)
}

func loop(thread *rtda.Thread) {
	reader := &base.BytecodeReader{}
	for {
		frame := thread.TopFrame()
		pc := frame.NextPc()
		thread.SetPc(pc)
		reader.Reset(frame.Method().Code(), pc)
		opCode := reader.ReadUint8()
		ins := instructions.NewInstruction(opCode)
		fmt.Println(reflect.TypeOf(ins))
		ins.FetchOperands(reader)
		frame.SetNextPc(reader.Pc())
		ins.Execute(frame)
		fmt.Println(frame.OperandStack())
		fmt.Println(frame.LocalVars())
		if thread.StackIsEmpty() {
			break
		}
	}

}
func catchErr(thread *rtda.Thread) {
	if r := recover(); r != nil {
		logFrame(thread)
		panic(r)
	}
}
func logFrame(thread *rtda.Thread) {
	for !thread.StackIsEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class()
		fmt.Printf(">> pc: %4d %v %v %v \n", frame.NextPc(), className, method.Name(), method.Descriptor())
	}
}
