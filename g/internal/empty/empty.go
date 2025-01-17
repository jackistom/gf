// Copyright 2019 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Package empty provides checks for empty variables.
package empty

import (
	"reflect"
)

// IsEmpty checks whether given <value> empty.
// It returns true if <value> is in: 0, nil, false, "", len(slice/map/chan) == 0.
// Or else it returns true.
func IsEmpty(value interface{}) bool {
    if value == nil {
        return true
    }
    // 优先通过断言来进行常用类型判断
    switch value := value.(type) {
        case int:     return value == 0
        case int8:    return value == 0
        case int16:   return value == 0
        case int32:   return value == 0
        case int64:   return value == 0
        case uint:    return value == 0
        case uint8:   return value == 0
        case uint16:  return value == 0
        case uint32:  return value == 0
        case uint64:  return value == 0
        case float32: return value == 0
        case float64: return value == 0
        case bool:    return value == false
        case string:  return value == ""
        case []byte:  return len(value) == 0
        default:
        	// Finally using reflect.
            rv := reflect.ValueOf(value)
            switch rv.Kind() {
                case reflect.Chan,
                     reflect.Map,
	                 reflect.Slice,
	                 reflect.Array:
                    return rv.Len() == 0

	            case reflect.Func,
		             reflect.Ptr,
		             reflect.Interface,
		             reflect.UnsafePointer:
		            if rv.IsNil() {
			            return true
		            }
            }
    }
    return false
}