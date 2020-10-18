package main

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	assert := []int{1,2,3}
	result := QuickSort([]int{2,3,1})
	if !reflect.DeepEqual(assert, result) {
		t.Error(assert, " not equal ", result)
	}
}