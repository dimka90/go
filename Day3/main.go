package main


import (
"fmt"
"consensus/consensusClient"
)


func main (){

fmt.Println(consensusClient.SetSlot(90))
fmt.Println(consensusClient.GetSlot())
}
