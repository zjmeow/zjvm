package heap

import "unicode/utf16"

var internedStrings = map[string]*Object{}

// 转成java 的 string 类
func JString(loader *ClassLoader, goStr string) *Object {
	if str, ok := internedStrings[goStr]; ok {
		return str
	}
	chars := stringToUtf16(goStr)
	jChars := NewDefaultObject(loader.LoadClass("[C"), chars)
	jStr := loader.LoadClass("java/lang/String").NewObject()
	jStr.SetRefVar("value", "[C", jChars)
	internedStrings[goStr] = jStr
	return jStr
}

func stringToUtf16(s string) []uint16 {
	runes := []rune(s) // utf32
	return utf16.Encode(runes)
}
func InternString(jStr *Object) *Object {
	goStr := GoString(jStr)
	if internedStr, ok := internedStrings[goStr]; ok {
		return internedStr
	}
	internedStrings[goStr] = jStr
	return jStr
}
func GoString(jStr *Object) string {
	charArr := jStr.GetRefVar("value", "[C")
	return string(utf16.Decode(charArr.Chars()))
}
