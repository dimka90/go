package main

import "fmt"


func main(){

num1,_ := getPoints()

fmt.Println(num1)
}


func getPoints() (int, int) {

return 2, 3
}
