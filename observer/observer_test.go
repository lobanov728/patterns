package main

import (
	"reflect"
	"testing"
)

func TestEmpty(t *testing.T) {
	assert := []int{1,2,3}
	result := []int{2,3,1}
	if !reflect.DeepEqual(assert, result) {
		t.Error(assert, " not equal ", result)
	}
}