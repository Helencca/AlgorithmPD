package solution1

func longestCommonPrefix(strs []string) string {
	// failed 
	//1.1 变成二维数组遍历
	//1.2 将第一个字符串和第二个字符串对比，取共同部分
	//1.3 再将前两者比较的共同部分与第三者比较，取共同，以此类推
	//1.4 但凡出现没有的就返回空值
	//1.5 不然就继续

	if len(strs) == 0 {
		return " "
}

res := strs[0]

for i:=0; i <len(strs); i++{
	for j:=0; j<len(strs); j++{
		if strs[i][j] != strs[i+1][j]{
			break
		} else{
			//怎么获取共同部分呢
			res = res[0:j] 
		}
	}
}
	

return res
}