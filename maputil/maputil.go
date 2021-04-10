package maputil

import (
	"reflect"
)

// Transform 生成新的map返回
func Transform(slice, function interface{}) interface{} {
	return transform(slice, function, false)
}

// Transform 替换原有的map
func TransformInPlace(slice, function interface{}) interface{} {
	return transform(slice, function, true)
}

func transform(mapType, function interface{}, inPlace bool) interface{} {
	mapInType := reflect.ValueOf(mapType)
	if mapInType.Kind() != reflect.Map {
		panic("transform: not mapType")
	}

	fn := reflect.ValueOf(function)
	elemType := mapInType.Type().Elem()
	if !verifyFuncSignature(fn, elemType, nil) {
		panic("Transform: function must be of type func(" + mapInType.Type().Elem().String() + ") outputElemType")
	}

	mapOutType := mapInType
	if !inPlace {
		mapOutType = reflect.MakeMapWithSize(reflect.MapOf(fn.Type().In(0), fn.Type().In(1)), mapInType.Len())
	}
	for _, key := range mapInType.MapKeys() {
		newValue := fn.Call([]reflect.Value{key, mapInType.MapIndex(key)})[0]
		mapOutType.SetMapIndex(key, newValue)
	}
	return mapOutType.Interface()
}

// 校验函数
func verifyFuncSignature(fn reflect.Value, types ...reflect.Type) bool {
	// check is a function
	if fn.Kind() != reflect.Func {
		return false
	}
	// check parameter count
	if (fn.Type().NumIn() != 2) || (fn.Type().NumOut() != 1) {
		return false
	}
	// check in parameter type
	for i := 0; i < len(types)-1; i++ {
		if fn.Type().In(i) != types[i] {
			return false
		}
	}
	// check out parameter type
	outType := types[len(types)-1]
	if outType != nil && fn.Type().Out(0) != outType {
		return false
	}
	return true
}
