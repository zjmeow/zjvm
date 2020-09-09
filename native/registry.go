package native

import "github.com/zjmeow/zjvm/rtda"

type Method func(frame *rtda.Frame)

var registryMap = map[string]Method{}

func Register(className, methodName, descriptor string, method Method) {
	key := className + "-" + methodName + "-" + descriptor
	registryMap[key] = method
}
func FindNativeMethod(className, methodName, descriptor string) Method {
	key := className + "-" + methodName + "-" + descriptor
	if method, ok := registryMap[key]; ok {
		return method
	}
	if descriptor == "()V" && methodName == "registerNatives" {
		return emptyNativeMethod
	}
	//return nil
	panic("method not found")
}

func emptyNativeMethod(frame *rtda.Frame) {

}
