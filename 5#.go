package  main

import (
	"fmt"
	"strings"
)

func longestPalindrome(s string) string {
	charList := strings.Split(s,"");
	length := len(charList)
	//resultList := make([]string,0,length)

	stack := make([]string,0,length);
	res := ""

	for i:=0;i<length;{
		cur := charList[i]
		//stack info
		topPointer := len(stack)-1
		secondPointer := topPointer -1

		//judge
		if secondPointer >= 0{ //todo 两个字符
			if cur == stack[secondPointer] { // 发现回文转折 递增i
				height := secondPointer
				ii := i
				newRes := stack[topPointer]
				for {
					if ii>=length || height<0  {
						//i = ii
						if height>=0 && ii<length && stack[height] == charList[ii] {
							newRes = charList[ii] + newRes + charList[ii]
							i+=1;
						}
						if len(res)<len(newRes) {
							res = newRes
						}
						//stack = make([]string,0,length)
						break;
					}

					if stack[height] != charList[ii] {
						if len(res)<len(newRes) {
							res = newRes
						}
						//stack = make([]string,0,length)
						break;
					}

					newRes = charList[ii]+newRes+charList[ii]
					height-=1;
					ii+=1;
				}
			}
		}
		stack = append(stack,cur)
		i+=1
		// 针对两个字符特殊处理
		if i>=length && secondPointer<0 {
			if stack[0] == stack[1] {
				return stack[0]+stack[1]
			}
		}
		fmt.Println(secondPointer)
	}
	return  res
}

func main() {
	result := longestPalindrome("aaaa"); // aa aaaaaa
	fmt.Println(result)
}
