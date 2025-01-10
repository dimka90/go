package main


import "fmt"

func main(){
// Declare a vairable 
// var name = "Dimka"

nums := make([]int,3)
nums[0] =4
//fmt.Println(names)
printArray(nums)
}

func printArray(nums []int){
for i :=0; i<cap(nums); i++{
fmt.Print(nums[i], "==>"  + " ")
}
}
