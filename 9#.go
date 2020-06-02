package main

import (
	"math"
	//"fmt"
)

func intLength(x int) int{
	remain := x;
	looper := 0;
	for{
		remain /= 10;
		looper += 1
		if remain == 0 {
			break;
		}
	}
	return  looper
}
func isPalindrome(x int) bool {
	if x < 0 {
		return  false;
	}
	var reverseX int;
	length := intLength(x);
	if length == 1 {
		return true
	}
	looper := 0;
	remain := x;
	for{
		curLevel := int(math.Pow10(length-looper-1));
		cur := remain / curLevel;
		remain -= curLevel * cur;

		reverseX += cur * int(math.Pow10(looper));
		if remain == 0 {
			break;
		}
		looper += 1;
	}
	//fmt.Println(reverseX)
	return  x == reverseX;
}

func main(){
	isPalindrome(11133);
}

