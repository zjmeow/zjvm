package main

import (
	"fmt"
	"github.com/zjmeow/zjvm/classpath"
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
	"github.com/zjmeow/zjvm/rtda/heap"
	"strings"
)

type Jvm struct {
	cmd         *Cmd
	classLoader *heap.ClassLoader
	mainThread  *rtda.Thread
}

func newJvm(cmd *Cmd) *Jvm {
	cp := classpath.Parse(cmd.XjreOption, cmd.CpOption)
	classLoader := heap.NewClassLoader(cp)
	return &Jvm{
		cmd:         cmd,
		classLoader: classLoader,
		mainThread:  rtda.NewThread(),
	}
}

func (j *Jvm) start() {
	vmClass := j.classLoader.LoadClass("sun/misc/VM")
	base.InitClass(j.mainThread, vmClass)
	interpret(j.mainThread)
}

func (j *Jvm) execMain() {
	className := strings.Replace(j.cmd.class, ".", "/", -1)
	mainClass := j.classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod == nil {
		fmt.Printf("Main method not found in class %s\n", j.cmd.class)
		return
	}
	argsArr := j.createArgsArray()
	frame := j.mainThread.NewFrame(mainMethod)
	frame.LocalVars().SetRef(0, argsArr)
	main()
	j.mainThread.PushFrame(frame)
	interpret(j.mainThread)
}

func (j *Jvm) createArgsArray() *heap.Object {
	stringClass := j.classLoader.LoadClass("java/lang/String")
	argsLen := uint(len(j.cmd.args))
	argsArr := stringClass.ArrayClass().NewArray(argsLen)
	jArgs := argsArr.Refs()
	for i, arg := range j.cmd.args {
		jArgs[i] = heap.JString(j.classLoader, arg)
	}
	return argsArr
}
