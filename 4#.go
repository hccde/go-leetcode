package  main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	pointer1 := 0;pointer2 := 0;
	length1 := len(nums1);length2 := len(nums2);
	var longer []int;var shorter []int;
	end := (length1+length2)/2
	isOdd := true //奇数
	if (length1+length2)%2 == 0 {
		isOdd = false
	}
	if length1 > length2 {
		longer = nums1
		shorter = nums2

	}else{
		longer = nums2
		shorter = nums1
	}
	length1 = len(longer)
	length2 = len(shorter)
	if length1 == 0 {
		return  0.0
	} else if length2 == 0 && length1>=4 {
		shorter = []int{math.MaxInt64};
	}

	if length1+length2 == 1 {
		return float64(longer[0])
	}

	if length2+length1 == 2 {
		if len(shorter)>0 {
			return (float64(longer[0])+float64(shorter[0]) )/2
		}else{
			return  (float64(longer[1])+float64(longer[0]) )/2
		}
	}

	if length2+length1 == 3 {
		if len(shorter) == 0 {
			return  float64(longer[1])
		}else if longer[0]>shorter[0]{
			return  float64(longer[0])
		}else if longer[1] > shorter[0]{
			return float64(shorter[0])
		}else if shorter[0] > longer[1] {
			return  float64(longer[1])
		}
	}

	//var cur int;
	j:=0
	for i:=0;i<len(longer);{
		if j == end-1 {
			val1 := math.MaxInt64;val2 := math.MaxInt64;val:=0
			if pointer1 < length1 {
				val1 = longer[pointer1]
			}
			if pointer2 < length2 {
				val2 = shorter[pointer2]
			}
			val = val1
			vall := val2
			if val1 > val2 { val = val2;vall = val1 }
			next1 := math.MaxInt64;next2 := math.MaxInt64;next:=0
			if pointer1+1 < length1 {
				next1 = longer[pointer1+1]
			}
			if pointer2+1 < length2 {
				next2 = shorter[pointer2+1]
			}

			next = next1
			if next1 > next2 {  next = next2 }
			if next > vall {
				next = vall
			}
			fmt.Println(val,next)
			// cur是上一次的值
			// val是当前值
			// next 是下一个值
			if isOdd {
				// 奇数 下一个
				return float64(next)
			}else{
				return (float64(val) + float64(next))/2
			}
		}

		if longer[pointer1] > shorter[pointer2]{
			//cur = shorter[pointer2]
			if pointer2+1 < length2 {pointer2+=1}

		}else{
			//cur = longer[pointer1]
			if pointer1+1 < length1 {pointer1+=1}
			i+=1
		}
		j+=1;
	}
	return  0.0
}

func main() {

	num1 := []int {1}
	num2 :=  []int {2,3,4}
	result := findMedianSortedArrays(num1,num2)
	fmt.Println(result);
}
