package main

import "fmt"

func isMatch(s string, p string) bool {
	pLength := len(p);
	sLength := len(s);
	pIndex := 0;
	sIndex := 0;
	hasStar := false;
	// *是个操作符 具有左结合性 考虑 a** 这里有未定义行为
	for {
		if pIndex >= pLength || sIndex >= sLength{
			if pIndex < pLength && hasStar{
				fmt.Println(22)
				index := 1;
				pIndexCopy := pIndex;
				for { // aaa a*a aacc .*cc
					if pLength - index == pIndexCopy {
						//fmt.Println(string(p[pIndexCopy]),string(s[sLength-index]))
						if p[pIndexCopy] != s[sLength-index] {
							return  false;
						}else{
							pIndex+=1
							break;
						}
					}

					if p[pLength-index] != s[sLength-index] {
						//fmt.Println( string(p[pLength-index]),string(s[sLength-index]))
						return false
					}else{
						pIndex+=1
						index+=1
					}

				}
			}
			break;
		}

		chr := p[pIndex];
		//fmt.Println(string(chr),string(s[sIndex]))
		if chr == s[sIndex] || chr == 46 { // . 或者匹配
			sIndex += 1;
			pIndex += 1;
		}else if chr != 42 { //不匹配
			nextIndex := pIndex+1;
			if nextIndex >= pLength {
				break;
			}
			nextChr := p[nextIndex]
			if nextChr == 42 { // a* 出现0次
				hasStar = true;
				pIndex += 2;
				//sIndex += 1;
			}else{
				return  false;
			}

		}

		if chr == 42 { // *
			hasStar = true;
			lastIndex := pIndex-1;
			if lastIndex < 0 {return  false}
			lastChr := p[lastIndex];
			if lastChr == 46 { // .* 匹配之后的所有
				if pIndex+1 < pLength { // .*匹配全部之后仍有 例如 .*a
					sIndex = sLength;
					pIndex += 1;
				}else{
					return  true;
				}
			}else {
				for{
					if sIndex == sLength {
						break;
					}
					if s[sIndex] != lastChr  {
						break;
					}
					sIndex += 1;
				}
				pIndex += 1;
			}
		}
	}

	if sIndex < sLength  {
		return  false;
	}
	if sIndex == sLength && pIndex == pLength{
		return  true;
	}
	//fmt.Println(pIndex,pLength)
	return  false;
}

func main(){
	res := isMatch("aaa","ab*a*c*a") // aaa a*a  aaa aaaa aacc .*cc
	fmt.Println(res);
}
