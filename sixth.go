package main

import (
	"strings"
	"strconv"
	"fmt"
	"sort"
)

func main()  {
	minStorage := 28625
	volumes := strings.Split(i, " ")
	hardDisks := make([]int, len(volumes))
	for i, v := range volumes {
		volume, err := strconv.Atoi(v)
		if err!=nil {
			fmt.Println("Error")
		}
		hardDisks[i]= volume
	}
	sort.Sort(sort.Reverse(sort.IntSlice(hardDisks)))
	for i := range hardDisks {
		if sum(hardDisks[:i]...) > minStorage {
			fmt.Println(i)
			break
		}
	}
}

func sum(nums ...int) int  {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

var i = `34 132 181 232 593 413 862 887 808 18 35 89 356 640 339 280 975 82 345 398 948 372 91 755 75 153 948 603 35 694 722 293 363 884 264 813 175 169 646 138 449 488 828 417 134 84 763 288 845 801 556 972 332 564 934 699 842 942 644 203 406 140 37 9 423 546 675 491 113 587`