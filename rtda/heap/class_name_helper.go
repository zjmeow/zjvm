package heap

import "strings"

var primitiveTypes = map[string]string{
	"void":    "V",
	"boolean": "Z",
	"byte":    "B",
	"short":   "S",
	"int":     "I",
	"long":    "J",
	"char":    "C",
	"float":   "F",
	"double":  "D",
}

func getArrayClassName(className string) string {
	return "[" + toDescriptor(className)
}

func toDescriptor(className string) string {
	// 多维数组
	if strings.HasPrefix(className, "[") {
		return className
	}
	// 普通数组
	if d, ok := primitiveTypes[className]; ok {
		return d
	}
	// 类数组
	return "L" + className + ";"
}
