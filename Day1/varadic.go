package main

import "fmt"

func main() {

user := make(map[string]int)

user["dimka"] = 10
_ = user
num := []int{20,2,3,4,5}
variadic(num...)
}

func  variadic(nums...int){
total := 0
for _, num := range nums{
total+= num
}
fmt.Println(total)
}

func  variadicForMap(names...string) {
fmt.Println(names)
}
