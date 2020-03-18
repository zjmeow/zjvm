package main

import (
	"fmt"
	"strings"
)
import "./classpath"

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("v1.0")
	} else if cmd.helpFlag {
		fmt.Println("")
	} else {
		startJVM(cmd)
	}
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
}
