package main
import (
	"fmt"
)
type hotdog int
var t hotdog = 101

func main() {
	x:= 10
	fmt.Printf("%v\n", x)
	x = int(t)
	fmt.Printf("%v\n", x)
}
/*
- Conversão de tipos é o que soa.
- Em Go não se diz casting, se diz conversion.
- a = int(b)
- ref/spec#Conversions
*/