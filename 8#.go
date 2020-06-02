package main

import (
	"fmt"
	"math"
	"unicode"
)
const (
	SPACE = "space"
	SYMBOL = "symbol"
	NUMBER = "number"
	ZEROBEGIN = "zeroBegin"
)
type preState struct {
	Chr rune
	TypeString string //定义成常量
}
func myAtoi(str string) int {
	lastState := preState{Chr:rune(32),TypeString:SPACE};
	var valueString string;
	var value int; //valueString 倒序乘以个十百
	//得到valueString 以及符号 可选的符号默认为+
	for _, chr := range str {
		// chr 是否是数字 不是看last last为数字则返回 而last为空格 符号的情况  last是空格 看chr是否为空格 last符号 chr退出
		// 是数字的话看last 需要为+-符号 或者空格 或者数字
		// if chr == rune(' ') && last == rune(' ') {
		// 	continue;
		// }
		// if last == rune(' ') && chr not number { // 不是数字和符号 - +
		// 	return 0
		// }
		// if last == rune(' ') &&
		if unicode.IsNumber(chr) {
			if lastState.TypeString == SYMBOL || lastState.TypeString == SPACE || lastState.TypeString == NUMBER ||
				lastState.TypeString == ZEROBEGIN {
					if chr == rune(48) && (lastState.TypeString == SYMBOL || lastState.TypeString == SPACE || lastState.TypeString == ZEROBEGIN) {
						lastState.TypeString = ZEROBEGIN;
						lastState.Chr = chr;
						continue;
					}
					valueString += string(chr);
					lastState.TypeString = NUMBER;
					lastState.Chr = chr;
			}else{
				break;
			}
		}else{
			if lastState.TypeString == NUMBER {
				break;
			}
			if lastState.TypeString == SYMBOL {
				break;
			}
			if lastState.TypeString == ZEROBEGIN {
				break;
			}

			if lastState.TypeString == SPACE {
				if unicode.IsSpace(chr) {
					continue;
				}
				if chr == rune(43) || chr == rune(45) {
					lastState.TypeString = SYMBOL;
					valueString += string(chr);
					lastState.Chr = chr;
				}else{
					break;
				}

			}

		}
	}
	if valueString == "" {
		return 0;
	}
	//fmt.Println(valueString)
	if unicode.IsNumber(rune(valueString[0])){
		valueString = "+"+valueString;
	}
	isNegative := false;
	for index,chr := range valueString {
		if unicode.IsNumber(chr) {
			value += (int(chr)-48)* int(math.Pow10( (len(valueString)-1)-index))
			//fmt.Println(value)
		}else{
			if chr == rune(45) {
				isNegative = true;
			}
		}
	}
	if isNegative {value = -value;}
	//超出go int 内部溢出
	if len(valueString)>11{
		if isNegative {
			return  math.MinInt32
		}else{
			return  math.MaxInt32
		}
	}
	if value > math.MaxInt32 {
		return math.MaxInt32
	}
	if value < math.MinInt32 {
		return  math.MinInt32
	}
	return value;
}

func main() {
	value := myAtoi("  ++") // " 123" "+123" "+123+4" "-123" "t123" " t123" "+++" "+0+01"
	fmt.Println(value)
}