package main

import (
		"strconv"
	"fmt"
)

func main() {
	// Launch
	num := "1" // First  term
	for i:=0; i<41; i++ {
		num = lookAndSay(num)
	}
	fmt.Println(len(num)) // length of 42nd term

	//Test
	//fmt.Println(lookAndSay("1"), "11")
	//fmt.Println(lookAndSay("11"), "21")
	//fmt.Println(lookAndSay("21"), "1211")
	//fmt.Println(lookAndSay("1211"), "111211")
	//fmt.Println(lookAndSay("111211"), "311221")
}

func lookAndSay(look string) string{
	var say = ""
	var prev= string(look[0])
	var count=1
	for _, char := range look[1:] {
		if string(char) ==prev {
			count++
			continue
		}
		say += strconv.Itoa(count) + string(prev)
		prev = string(char)
		count = 1
	}
	return say + strconv.Itoa(count) + prev
}


//func main(){
//	nums := []int{1,2,3,4,5,6,7,8,9}
//	fmt.Println(sum(nums...))
//}

//func sum(nums ...int) int  {
//	total := 0
//	for _, num := range nums {
//		total += num
//	}
//	return total
//}

//func main(){
//	// Closures
//	nextInt := intSeq()
//	fmt.Println(nextInt())
//	fmt.Println(nextInt())
//	fmt.Println(nextInt())
//	fmt.Println(nextInt())
//}
//
//func intSeq() func() int {
//	i := 0
//	return func() int {
//		i++
//		return i
//	}
//}

//func zeroVal(val int)  {
//	val = 0
//}
//
//func zeroPtr(val *int){
//	*val = 0
//}
//
//type Person2 struct {
//	name string
//	age int
//}
//
//func main(){
//	val := 10
//	zeroVal(val)
//	fmt.Println("Zero", val)
//	zeroPtr(&val)
//	fmt.Println("Ptr", val)
//
//	azimjon := Person2{name: "Azimjon", age: 20}
//	fmt.Println(azimjon.age)
//	azimjon.age = 19
//	fmt.Println(azimjon.age)
//	azimjonPtr := &azimjon
//	azimjonPtr.age = 18
//	fmt.Println(azimjon.age)
//}
