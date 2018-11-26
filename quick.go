package main

import "fmt"

var sortArray  =[]int{41, 24, 76, 11, 45, 64, 21, 69, 19, 36}
func main() {
	sort(sortArray,0,len(sortArray)-1)
	for _,v:=range sortArray{
		fmt.Print(v,"  ")
	}
}
func sort(array []int,left,right int)  {
	if left<right{
		key:=partition(array,left,right)
		sort(array,left,key-1)
		sort(array,key+1,right)
	}
}
func partition(array []int,left,rigth int) int {
	//pivotkey:=array[left]
	i:=left
	j:=rigth
	for i<j{
		for i<j&&array[j]>=array[i] {
			j--
		}
		swap(i,j)
		for i<j&&array[i]<=array[j] {
			i++
		}
		swap(i,j)
	}
	return i
}
func swap(a,b int)  {
	sortArray[a],sortArray[b]=sortArray[b],sortArray[a]
}