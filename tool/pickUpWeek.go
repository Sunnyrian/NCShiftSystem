package tool

// PickupWeek 遍历Slice取出其中的周的字符串
func PickupWeek(s []string, count int) []string{
	var weekSlice []string
	for count>0{
		count --
		weekSlice = append(weekSlice, s[3])
		//把前面五个元素舍弃
		s = s[5:]
	}
	return weekSlice
}
