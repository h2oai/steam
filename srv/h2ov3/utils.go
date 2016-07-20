package h2ov3

import (
	"bytes"
	"fmt"
	"reflect"
)

func keyArrayToStringArray(keys interface{}) []string {
	// XXX verify inherits from Key
	arr := reflect.ValueOf(keys)
	strgs := make([]string, arr.Len())

	for i := 0; i < arr.Len(); i++ {
		key := reflect.Indirect(arr.Index(i))
		strgs[i] = key.FieldByName("Name").Interface().(string)
	}

	return strgs
}

func DoubleStringArraysToSingle(array [][]string) []string {
	ret := make([]string, len(array))
	for i, nest := range array {
		var buf bytes.Buffer
		buf.WriteString("[")
		for j, val := range nest {
			if j > 0 {
				buf.WriteString(", ")
			}
			buf.WriteString(fmt.Sprintf("%q", val))
		}
		buf.WriteString("]")
		ret[i] = buf.String()
	}

	return ret
}
