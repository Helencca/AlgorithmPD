package solution2

func twoSum(nums []int, target int) []int{
temp := make(map[int]int)
for i:=0; i<len(nums);i++{
	 difference := target - nums[i]
	if v, ok := temp[difference]; ok {
		return []int{v,i}
	} else {
		temp[nums[i]]=i
	}
}
return []int{}
}
