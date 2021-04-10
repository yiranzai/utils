package structutil

import "reflect"

// 结构体字段
type Field struct {
	Name, JsonTag, NameTag string
}

// 获取结构体的名称
func GetFields(s interface{}) []Field {
	var results []Field
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Struct {
		t := reflect.TypeOf(s)
		for i := 0; i < v.NumField(); i++ {
			results = append(results, Field{
				Name:    t.Field(i).Name,
				JsonTag: t.Field(i).Tag.Get("json"),
				NameTag: t.Field(i).Tag.Get("name"),
			})
		}
	}
	return results
}

// GetEntityColumns 获取实体字段(s)
func GetEntityColumns(columns []string, i interface{}) []Field {
	s := reflect.TypeOf(i)
	var results []Field
	for i := 0; i < s.NumField(); i++ {
		if len(columns) >= 1 && columns[0] != "" {
			for _, column := range columns {
				if column == s.Field(i).Tag.Get("json") {
					results = append(results, Field{
						Name:    s.Field(i).Name,
						JsonTag: s.Field(i).Tag.Get("json"),
						NameTag: s.Field(i).Tag.Get("name"),
					})
				}
			}
		} else {
			results = append(results, Field{
				Name:    s.Field(i).Name,
				JsonTag: s.Field(i).Tag.Get("json"),
				NameTag: s.Field(i).Tag.Get("name"),
			})
		}
	}
	return results
}

// CheckIsSlice 判断是否为slice数据
func CheckIsSlice(arg interface{}) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)
	if val.Kind() == reflect.Slice {
		ok = true
	}
	return
}

// InterfaceToArraySlice interface{} 转为数组切片
func InterfaceToArraySlice(resource interface{}) []interface{} {
	data, ok := CheckIsSlice(resource)
	if !ok {
		return nil
	}
	length := data.Len()
	result := make([]interface{}, length)

	for i := 0; i < length; i++ {
		result[i] = data.Index(i).Interface()
	}
	return result
}

//Interface2 interface转化
func Interface2(value interface{}) interface{} {
	switch value.(type) {
	case bool:
		return value.(bool)
	case uint8:
		return value.(uint8)
	case uint16:
		return value.(uint16)
	case uint32:
		return value.(uint32)
	case uint64:
		return value.(uint64)
	case int8:
		return value.(int8)
	case int16:
		return value.(int16)
	case int32:
		return value.(int32)
	case int64:
		return value.(int64)
	case float32:
		return value.(float32)
	case float64:
		return value.(float64)
	case complex64:
		return value.(complex64)
	case complex128:
		return value.(complex128)
	case string:
		return value.(string)
	case int:
		return value.(int)
	case uint:
		return value.(uint)
	case uintptr:
		return value.(uintptr)
	}
	return nil
}
