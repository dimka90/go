package main


import "fmt"

func main(){
// Declare a vairable 
// var name = "Dimka"

var nums [5]int
nums[0] =4
//fmt.Println(names)
printArray(nums)
}

func printArray(nums [5]int){
for i :=0; i<len(nums); i++{
fmt.Print(nums[i], "==>"  + " ")
}
}
