package main

import (
	"fmt"
	"math"
)

func main(){
	//var result uint64 =  1
	//for i:=0;i<42;i++ {
	//	result *= 2
	//}
	//fmt.Println(result)
	// OR
	var another uint64 = 1<<42
	fmt.Println(another)
	fmt.Println(math.Pow(2, 42))
}
