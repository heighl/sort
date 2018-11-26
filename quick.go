package main


//func main() {
//	var sortArray  =[]int{41, 24, 76, 11, 45, 64, 21, 69, 19, 36}
//	Quicksort(sortArray,0,len(sortArray)-1)
//	for _,v:=range sortArray{
//		fmt.Print(v,"  ")
//	}
//}
func Quicksort(array []int,left,right int)  {
	if left<right{
		key:=partition(array,left,right)
		Quicksort(array,left,key-1)
		Quicksort(array,key+1,right)
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
		swap(array,i,j)
		for i<j&&array[i]<=array[j] {
			i++
		}
		swap(array,i,j)
	}
	return i
}
func swap(array []int,a,b int)  {
	array[a],array[b]=array[b],array[a]
}

