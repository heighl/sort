package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuicksort(t *testing.T)  {
	var s =[]int{3,1,4}
	Quicksort(s,0,len(s)-1)
	assert.Equal(t,[]int{1,3,4},s)
}
func TestM(t *testing.T)  {
	var sortArray  =[]int{41, 24, 76, 11, 45, 64, 21, 69, 19, 36}
	Quicksort(sortArray,0,len(sortArray)-1)
	assert.Equal(t,[]int{11,19, 21, 24, 36 ,41 ,45 ,64 ,69, 76},sortArray)
}

func TestThird(t *testing.T)  {
	var array  =[]int{3,6,55,8,20}
	Quicksort(array,0,len(array)-1)
	assert.Equal(t,[]int{3,6,8,20,55},array)
}