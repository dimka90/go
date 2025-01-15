package main

import (
"operation/math"
"operation/english"
"fmt"
)
func init() {
fmt.Println("I am the starting Point")

}
func main(){
fmt.Println(math.Add(9,0))
num :=[]int{2,9,0,9, 90, 90, 39, 58, 59}
fmt.Println(math.AddmultipleInt(num...))
fmt.Println(math.AddmultipleFloat(9.0,11.0,90.9,9.4))
fmt.Println(english.GetName("dimka"))
fmt.Println(math.Modtwo(45))
}
