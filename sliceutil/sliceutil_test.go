package sliceutil

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFilterInPlace(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expect := []int{1, 3, 5, 7, 9}
	FilterInPlace(&list, func(i int) bool { return i%2 == 1 })
	if !reflect.DeepEqual(expect, list) {
		t.Fatalf("FilterInPlace failed: expect %v got %v", expect, list)
	}
}

func TestTransform(t *testing.T) {
	list := []string{"1", "3", "5"}
	expect := []string{"111", "333", "555"}
	result := Transform(list, func(s string) string { return s + s + s })
	if !reflect.DeepEqual(expect, result) {
		t.Fatalf("Transform failed: expect %v got %v", expect, result)
	}
}

type CompanyCategory struct {
	Category    string
	CompanyName string
}

func TestGroupBy(t *testing.T) {
	list := []CompanyCategory{
		{"赛道1", "公司1"},
		{"赛道1", "公司2"},
		{"赛道2", "公司3"},
		{"赛道2", "公司4"},
		{"赛道2", "公司5"},
		{"赛道3", "公司6"},
	}
	expect := [][]CompanyCategory{
		{{"赛道1", "公司1"}, {"赛道1", "公司2"}},
		{{"赛道2", "公司3"}, {"赛道2", "公司4"}, {"赛道2", "公司5"}},
		{{"赛道3", "公司6"}},
	}

	result := GroupBy(list, func(elem CompanyCategory) string {
		return elem.Category
	})

	if fmt.Sprintf("%v", expect) != fmt.Sprintf("%v", result) {
		t.Fatalf("GroupBy failed: expect %v got %v", expect, result)
	}
}
