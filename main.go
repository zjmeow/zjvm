package main

import (
	"fmt"
	"github.com/zjmeow/zjvm/classfile"
	"github.com/zjmeow/zjvm/classpath"
	"io/ioutil"
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
	cmd.class = "java.lang.Object"
	startJVM(cmd)
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.CpOption)
	fmt.Println(cmd)
	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}
	fmt.Println(classData)
	bytes, err := ioutil.ReadFile("C:/Users/Administrator/Desktop/GuassTest.class")
	cf, err := classfile.Parse(bytes)
	if err != nil {
		panic(err)
	}
	mainMethod := getMainMethod(cf)
	interpret(mainMethod)
	fmt.Println(cf.ConstantPool())
	fmt.Println(cf.AccessFlags())
	fmt.Println(cf.ThisClass())
	fmt.Println(cf.SuperClass())
}

func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo {
	for _, m := range cf.Methods() {
		if m.Name() == "main" && m.Descriptor() == "([Ljava/lang/String;)V" {
			return m
		}
	}
	return nil
}
