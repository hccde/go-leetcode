package main
type preState struct {
	Chr rune,
	TypeString string //定义成常量
}
func myAtoi(str string) int {
	var lastState preState;
	var valueString string;
	var value int; //valueString 倒序乘以个十百
	isNegative := false; // todo 没有遇到负号直接加入＋号 统一处理
	//得到valueString 以及符号 可选的符号默认为+
	for index, chr range str {
		// chr 是否是数字 不是看last last为数字则返回 而last为空格 符号的情况  last是空格 看chr是否为空格 last符号 char退出
		// 是数字的话看last 需要为+-符号 或者空格 或者数字
		// if chr == rune(' ') && last == rune(' ') {
		// 	continue;
		// }
		// if last == rune(' ') && chr not number { // 不是数字和符号 - +
		// 	return 0
		// }
		// if last == rune(' ') && 
	}
	return 0;
}

func main() {
	myAtoi("123") // " 123" "+123" "+123+4" "-123" "t123" " t123" "+++"
}