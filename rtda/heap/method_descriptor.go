package heap

import (
	"strings"
)

type MethodDescriptor struct {
	parameterTypes []string
	returnType     string
}

// todo 补全解析
func parseMethodDescriptor(descriptor string) MethodDescriptor {
	res := MethodDescriptor{}
	res.returnType = strings.Split(descriptor, ")")[1]
	// 剥离特殊符号
	idx := strings.Index(descriptor, ")")
	descriptor = descriptor[1:idx]
	if descriptor == "" {
		return res
	}
	var parameterTypes []string

	for parameter, descriptor := parseNext(descriptor); parameter != ""; parameter, descriptor = parseNext(descriptor) {
		parameterTypes = append(parameterTypes, parameter)
	}
	res.parameterTypes = parameterTypes
	return res
}

func parseNext(descriptor string) (string, string) {
	if descriptor == "" {
		return "", ""
	}
	firstChar := string(descriptor[0])
	switch firstChar {
	case "B", "C", "D", "F", "I", "J", "S", "Z":
		return firstChar, descriptor[1:]
	case "L":
		endIndex := strings.Index(descriptor, ";")
		return descriptor[1:endIndex], descriptor[endIndex:]
	case "[":
		res, sub := parseNext(descriptor[1:])
		return "[" + res, sub
	}
	return "", ""
}
