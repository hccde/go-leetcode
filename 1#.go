package  main

import "fmt"
func twoSum(nums []int, target int) []int {
	size:=len(nums);
	copy := make([]int,size)
	res := make([]int,2)
	for i:=0;i<size;i++{
		copy[i] = target - nums[i]
		for j:=0;j<i;j++ {
			if copy[j] == nums[i]{
				res := []int{j,i}
				return res
			}
		}
	}

	return res;
}

func main() {
	var input = make([]int,5)
	input = []int{2,7,11,15}
	res := twoSum(input,9)
	fmt.Println(res);
}
