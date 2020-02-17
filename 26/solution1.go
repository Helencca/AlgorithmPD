package solution1

func removeDuplicates(nums []int) int {
	if len(nums) < 2{
		return len(nums)
	}  
  
	int := 0
	for j:= 1; j< len(nums); j ++{
		if nums[int] != nums[j] {
			int++
			nums[int] = nums[j]
		}
	}
  return int +1
  }