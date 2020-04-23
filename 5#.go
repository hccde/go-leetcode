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
		if secondPointer >= 0{ //todo 两个字符
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
				height := topPointer
				ii := i
				newRes := ""
				g1 := false
				g2 := false
				for{
					if ii < length {
						if charList[ii] == cur {
							newRes = newRes+cur
							ii+=1
						}else{
							g2=true
						}
					}else{
						g2=true
					}
					if height>=0{
						if stack[height] == cur {
							newRes = cur + newRes
							height-=1;
						}else{
							g1=true
						}
					}else {
						g1 = true
					}


					if g1 && g2  {
						if len(res)<len(newRes) {
							res = newRes
						}
						break;
					}

					if height>=0 && ii<length && stack[height] != cur && charList[ii] != cur {
						if len(res)<len(newRes) {
							res = newRes
						}
						break
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
	result := longestPalindrome("cabba"); // cabba
	fmt.Println(result)
}
