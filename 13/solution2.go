package solution

func romanToInt(s string) int {

	//1.先将字符对应的值以及特殊情况的字符记录到字典中
	//2.判断当前字符和后一个字符是否是特殊对
	//3.如果是特殊对，直接从字典中获取值
	//4.如果不是，获取当前值，直到取到最后一个字符为止
	romanMap :=map[string]int {"I":1 , "V":5, "X":10, "L":50, "C":100, "D":500, "M":1000, "IV":4, "IX":9, "XL" :40, "XC":90, "CD":400, "CM":900}
	
	 res :=0
	len := len(s)
	for i:=0; i <len; i++{
	   
		if len ==1 {
		return romanMap[string(s[i])]
		break
		}
	
		if i< len-1{
		curString :=string(s[i])+string(s[i+1])
		if _, ok := romanMap[curString]; ok{
			res += romanMap[curString]
			i++
			continue
		}
		}
		
		 if _, ok := romanMap[string(s[i])]; ok{
			res +=romanMap[string(s[i])]
		}
		
		
	}
	return res
	}