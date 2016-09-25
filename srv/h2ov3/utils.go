/*
  Copyright (C) 2016 H2O.ai, Inc. <http://h2o.ai/>

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU Affero General Public License as
  published by the Free Software Foundation, either version 3 of the
  License, or (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU Affero General Public License for more details.

  You should have received a copy of the GNU Affero General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

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
