package math

func Add(a int,b int)int{

return a + b
}
// function to add multiple numbers using variadic function format
func AddmultipleInt(num...int) int {
total  := 0
for _,number := range num {
total += number
};
return total
}

func AddmultipleFloat(num...float64) float64{
total := 0.0
for _,number := range num {
total += number
}

return total
}
