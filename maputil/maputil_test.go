package maputil

import (
	"reflect"
	"testing"
)

func TestTransform(t *testing.T) {
	list := map[string]string{"test1": "1111", "test2": "2222"}
	expect := map[string]string{"test1": "11111111", "test2": "22222222"}
	result := Transform(list, func(k string, v string) string {
		return k + v
	})
	if !reflect.DeepEqual(expect, result) {
		t.Fatalf("Transform failed: expect %v got %v", expect, result)
	}
}
