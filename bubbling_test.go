package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFirst(t *testing.T)  {
	var array=[]int64{10,100,23,56,48}
	bubbling(array)
	assert:=assert.New(t)
	assert.Equal([]int64{10,23,48,56,100},array,"first pass")
}
