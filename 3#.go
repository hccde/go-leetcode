package main

import (
	"fmt"
	"strings"
)
// 空字符串的边界情况，没说no-empty 行吧
func compareStrings(res *[]string, index int, list *[]string) (*[]string,int){
	cur := (*list)[index]
	maxLength := len(*res)
	//fmt.Println(*res,cur)
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
				temRes,maxLength2 := compareStrings(&tem, temIndex, list)
				var longer *[]string
				if len(*temRes) >= len(*res) {
					longer = temRes
				}else{
					longer = res
				}
				return  longer,max(maxLength,maxLength2)
			} else {
				return  &tem,max(maxLength,len(tem));
			}

		}
	}
	//add 1 char
	if (index < len(*list)) {
		tem := append(*res, cur)
		if (index+1 < len(*list)) {
			temRes,maxLength2 := compareStrings(&tem, index+1, list)
			return  temRes,max(maxLength,maxLength2)
		} else {
			return &tem,max(maxLength,len(tem))
		}
	}
	return res,maxLength
}

func max(a int,b int) int{
	if a > b{
		return a
	}else{
		return b
	}
}

func lengthOfLongestSubstring(s string) int {
	list := strings.Split(s, "")
	if len(list) == 0 {
		return  0
	}
	res := make([]string,1)
	res[0] = list[0]
	i := 0
	_,result := compareStrings(&res, i, &list)
	return result;
}

func main() {
	num := lengthOfLongestSubstring("abaaaa")
	fmt.Println(num)
}
