package solution

func maxSubArray(nums []int) int {
	if len(nums) == 1{
		return nums[0] 
	}
	
	 max := nums[0]
	 sum := 0 
	 for i:=0;i< len(nums);i++{
		 //sum>0继续加，始终记录最大值
		if sum >=0{
			sum +=nums[i]
			if sum > max{
			max = sum
			} 
		}else{
			//如果sum<0,就丢弃之前的数
			sum =nums[i]
		if sum > max{
			max = sum
		} 
	 } 
	 }
	 return max
	}