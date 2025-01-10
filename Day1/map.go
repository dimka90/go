package main


import "fmt"

func main(){

user := createnewMap()
user1 := updateMap(user, "dimka", 10)
user2 := updateMap(user1, "james",20)
fmt.Println(user1)
fmt.Println(user)
fmt.Println(user2)
}

func  createnewMap() map[string]int {
return make(map[string]int)
}

func updateMap(data map[string]int,key string, value int) map[string]int {
// update map with the specific key
data[key] = value
return data
}

func deleteValue(data map[string]int, key string) {
delete(data,key)

}

//func clearMapingField(data map[string]int) {
//clear(data)
//}
