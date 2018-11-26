package main

import "fmt"

func bubbling(n []int64)  {
	for i:=0;i<len(n)-1 ;i++  {
		for j:=i+1;j<len(n);j++{
			if n[i]>n[j]{
				n[i],n[j]=n[j],n[i]
			}
		}
	}
}
func main() {
	var x []int64 =[]int64{10,20,100,3,65}
	bubbling(x)
	fmt.Println(x)
}