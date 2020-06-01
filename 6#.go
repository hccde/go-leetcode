package main

import (
	"fmt"
)

func printString(array *[][]rune,numRows int,colCounter int,stringLen int) string{
	res := make([]byte,stringLen);
	counter := 0;
	for i:=0;i<numRows;i++{
		for j:=0;j<=colCounter;j++{
			//fmt.Println(string((*array)[j][i]));
			if (*array)[j][i] != rune(0) {
				res[counter] = byte((*array)[j][i])
				counter += 1;
			}
		}
	}
	return  string(res);
}

func convert(s string, numRows int) string {
	l := len(s);
	if numRows == 1 {
		return s;
	}
	n := l / (numRows+numRows-2);
	numCols := (n+1)*(numRows-1);
	//row := make([]rune,numRows)
	array := make([][]rune,numCols)
	colCounter := -1;
	for index,chr := range s {
		remain := index % (numRows+numRows-2)
		if remain < numRows {
			if remain == 0 {
				colCounter+=1;
				row := make([]rune,numRows)
				row[0] = rune(chr);
				array[colCounter] = row;
			}else{
				array[colCounter][remain] = rune(chr);
			}
		}else{
			colCounter+=1;
			row := make([]rune,numRows)
			row[numRows-1 - (remain+1-numRows)]=rune(chr);
			array[colCounter] = row;
		}
	}
	res := printString(&array,numRows,colCounter,l)
	return  res;
}

func main() {
	result := convert("PAYPALISHIRING",2);
	//result := convert("test",3) //s为空
	fmt.Println(result)
}