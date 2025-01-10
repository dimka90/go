package main


import "fmt"

func main(){
// Declare a slice
nums := make([]int,3)
nums[0] =4

nums = append(nums, 20)
printArray(nums)
}

func printArray(nums []int){
for i :=0; i<len(nums); i++{
fmt.Print(nums[i], "==>"  + " ")
}
}
