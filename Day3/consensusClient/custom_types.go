package consensusClient

type Slot uint64

var slot Slot = 12

func SetSlot(value Slot) bool{
 *(&slot) = value
return true
}


func GetSlot() Slot{
return slot
}
