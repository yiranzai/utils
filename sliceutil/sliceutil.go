package sliceutil

import (
	"errors"
	"reflect"
)

// Filter
func Filter(slice, function interface{}) interface{} {
	result, _ := filter(slice, function, false)
	return result
}

// FilterInPlace
func FilterInPlace(slicePtr, function interface{}) {
	in := reflect.ValueOf(slicePtr)
	if in.Kind() != reflect.Ptr {
		panic("FilterInPlace: not a pointer to slice")
	}
	_, n := filter(in.Elem().Interface(), function, true)
	in.Elem().SetLen(n)
}

func filter(slice, function interface{}, inPlace bool) (interface{}, int) {
	sliceInType := reflect.ValueOf(slice)
	if sliceInType.Kind() != reflect.Slice {
		panic("filter: not slice")
	}

	fn := reflect.ValueOf(function)
	elemType := sliceInType.Type().Elem()
	if !verifyFuncSignature(fn, elemType, nil) {
		panic("filter: function must be of type func(" + sliceInType.Type().Elem().String() + ") outputElemType")
	}

	var count []int
	for i := 0; i < sliceInType.Len(); i++ {
		if fn.Call([]reflect.Value{sliceInType.Index(i)})[0].Bool() {
			count = append(count, i)
		}
	}

	out := sliceInType
	if !inPlace {
		out = reflect.MakeSlice(sliceInType.Type(), len(count), len(count))
	}
	for i := range count {
		out.Index(i).Set(sliceInType.Index(count[i]))
	}

	return out.Interface(), len(count)
}

// Transform 生成新的slice返回
func Transform(slice, function interface{}) interface{} {
	return transform(slice, function, false)
}

// Transform 替换原有的slice
func TransformInPlace(slice, function interface{}) interface{} {
	return transform(slice, function, true)
}

func transform(slice, function interface{}, inPlace bool) interface{} {
	sliceInType := reflect.ValueOf(slice)
	if sliceInType.Kind() != reflect.Slice {
		panic("transform: not slice")
	}

	fn := reflect.ValueOf(function)
	elemType := sliceInType.Type().Elem()
	if !verifyFuncSignature(fn, elemType, nil) {
		panic("Transform: function must be of type func(" + sliceInType.Type().Elem().String() + ") outputElemType")
	}

	sliceOutType := sliceInType
	if !inPlace {
		sliceOutType = reflect.MakeSlice(reflect.SliceOf(fn.Type().Out(0)), sliceInType.Len(), sliceInType.Len())
	}
	for i := 0; i < sliceInType.Len(); i++ {
		sliceOutType.Index(i).Set(fn.Call([]reflect.Value{sliceInType.Index(i)})[0])
	}
	return sliceOutType.Interface()
}

// 聚合的Item
type GroupItems struct {
	Group interface{}   `json:"group"`
	Count int           `json:"count"`
	Items []interface{} `json:"items"`
}

// GroupBy 把slice按回调函数的值分组
func GroupBy(slice, function interface{}) interface{} {
	sliceInType := reflect.ValueOf(slice)
	if sliceInType.Kind() != reflect.Slice {
		panic("GroupBy: not slice")
	}

	fn := reflect.ValueOf(function)
	elemType := sliceInType.Type().Elem()
	if !verifyFuncSignature(fn, elemType, nil) {
		panic("GroupBy: function must be of type func(" + sliceInType.Type().Elem().String() + ") outputElemType")
	}

	// group by callback function
	result := make([]*GroupItems, 0)
	flag := make(map[interface{}]*GroupItems)
	for i := 0; i < sliceInType.Len(); i++ {
		group := fn.Call([]reflect.Value{sliceInType.Index(i)})[0].Interface()
		if node, exist := flag[group]; !exist {
			node = &GroupItems{Group: group}
			result = append(result, node)
			flag[group] = node
		}
		flag[group].Items = append(flag[group].Items, sliceInType.Index(i).Interface())
		flag[group].Count++
	}
	return result
}

// 校验函数
func verifyFuncSignature(fn reflect.Value, types ...reflect.Type) bool {
	// check is a function
	if fn.Kind() != reflect.Func {
		return false
	}
	// check parameter count
	if (fn.Type().NumIn() != len(types)-1) || (fn.Type().NumOut() != 1) {
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

// 求最小值
func Min(slice []int) (int, error) {
	if len(slice) == 0 {
		return 0, errors.New("empty slice")
	}

	min := slice[0]
	for _, v := range slice {
		if v < min {
			min = v
		}
	}

	return min, nil
}

// InArray 是否存在于数组
func InArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}
	return
}
