package main

import (
	"fmt"
	"math"
	"strconv"
)

func reverse(x int) int {
	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0;
	}
	isNegative := false;
	if x < 0 {
		isNegative = true;
		x = -x;
	}
	numString := strconv.Itoa(x);
	res := "";
	begin := true;
	for _,chr := range numString {
		if begin && chr != rune(0) {
			begin = false;
			res = string(chr)+res;
			continue;
		}
		if !begin {
			res = string(chr)+res;
		}
	}
	if isNegative {
		res = "-"+res;
	}
	//fmt.Println(res)
	value,_ := strconv.Atoi(res);
	if value > math.MaxInt32 || value < math.MinInt32 {
		return 0;
	}
	return  value;
}

func main() {
	res := reverse(1230);
	fmt.Println(res);
}