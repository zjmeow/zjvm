package main

import (
	"fmt"
	"github.com/zjmeow/zjvm/classpath"
	"github.com/zjmeow/zjvm/rtda/heap"
	"strings"
)

func main() {
	//cmd := parseCmd()
	//if cmd.versionFlag {
	//	fmt.Println("v1.0")
	//} else if cmd.helpFlag {
	//	fmt.Println("")
	//} else {
	//	startJVM(cmd)
	//}
	cmd := &Cmd{}
	cmd.class = "C:/Users/Administrator/Desktop/GuassTest.class"
	startJVM(cmd)
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.CpOption)
	fmt.Println(cmd)
	classLoader := heap.NewClassLoader(cp)
	className := strings.Replace(cmd.class, ".", "/", -1)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	interpret(mainMethod)
}
