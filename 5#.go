package  main

import (
	"fmt"
	"strings"
)

func longestPalindrome(s string) string {
	charList := strings.Split(s,"");
	length := len(charList)
	stack := make([]string,0,length);  //这个数组可以不要
	res := ""
	if length == 0 {
		return ""
	}
	for i:=0;i<length;{
		cur := charList[i]
		//stack info
		topPointer := len(stack)-1
		secondPointer := topPointer -1

		//judge
		if secondPointer >= 0{
			if cur == stack[secondPointer] { // 发现回文转折 递增i
				height := secondPointer
				ii := i
				newRes := stack[topPointer]
				for {
					if ii>=length || height<0  {
						if len(res)<len(newRes) {
							res = newRes
						}
						break;
					}

					if stack[height] != charList[ii] {
						if len(res)<len(newRes) {
							res = newRes
						}
						break;
					}

					newRes = charList[ii]+newRes+charList[ii]
					height-=1;
					ii+=1;
				}
			}
			//相同串
			if cur == stack[topPointer] {
				height := topPointer-1
				ii := i+1
				newRes := cur+stack[topPointer]
				g1 := false
				g2 := false
				for{
					//fmt.Println(ii,height)
					if ii < length && height>=0 {
						if charList[ii] == stack[height] {
							newRes = stack[height] + newRes+charList[ii]
							ii+=1
							height-=1;
						}else{
							g2=true
							g1=true
						}
					}else{
						if ii >= length {
							g2 = true
						}else{
							ii+=1
						}
						if height < 0 {
							g1 = true
						}else{
							height-=1
						}
					}

					if g1 && g2  {
						if len(res)<len(newRes) {
							res = newRes
						}
						break;
					}

				}
			}
		}
		stack = append(stack,cur)
		i+=1

		if secondPointer<0 && length>=2 {
			if charList[0] == charList[1]  && res=="" {
				res = charList[0]+charList[1]
			}

		}
	}
	// 没有返回第一个字符 不是很合理
	if res == "" {
		return charList[0]
	}
	return  res
}

func main() {
	result := longestPalindrome("abaabaa"); // cabba
	fmt.Println(result)
}
