package main

import (
	"fmt"
	"strings"
)
// 空字符串的边界情况，没说no-empty 行吧
func compareStrings(res *[]string, index int, list *[]string) *[]string {
	cur := (*list)[index]
	for i := 0; i < len(*res); i++ {
		if (*res)[i] == cur {
			// 有重复
			var tem []string
			if i+1 < len(*res) {
				tem = append((*res)[i+1:], cur)
			} else { //bugs here
				tem = make([]string, 1,len(*list))
				tem[0] = cur
			}

			temIndex := index + 1

			if temIndex <= len(*list)-1 {
				temRes := compareStrings(&tem, temIndex, list)
				if len(*temRes) >= len(*res) {
					return temRes
				}
			} else {
				if len(tem)>= len(*res){
					return &tem
				}else{
					//fmt.Println(*res)
					return res
				}
			}

		}
	}
	//add 1 char
	if (index < len(*list)) {
		tem := append(*res, cur)
		if (index+1 < len(*list)) {
			temRes := compareStrings(&tem, index+1, list)
			return  temRes
		} else {
			return &tem
		}
	}
	return res
}
func lengthOfLongestSubstring(s string) int {
	list := strings.Split(s, "")
	//size := len(list)
	if len(list) == 0 {
		return  0
	}
	res := make([]string,1)
	res[0] = list[0]
	i := 0
	result := compareStrings(&res, i, &list)
	return len(*result)
}

func main() {
	num := lengthOfLongestSubstring("abcabcbb")
	fmt.Println(num)
}
