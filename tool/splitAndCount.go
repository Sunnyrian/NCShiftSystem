package tool

import (
	xlsConst "NCShiftSystem/const"
	"strings"
)


// SplitAndCount 判断有多少行，推断出有多少门课在这个Cell中,并将Cell中元素切片
func SplitAndCount(s string)(count int, cellSlice []string){
	cellSlice = strings.Split(s, "\n")
	count = len(cellSlice)/xlsConst.PerCourseLine
	return count,cellSlice
}