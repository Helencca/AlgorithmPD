package solution1

func removeElement(nums []int, val int) int {
    if len(nums) == 0{
        return 0
    }
    int :=0
     for i:=0; i < len(nums); i++{
         //转换思路，把不等于val的重新排队赋值
     if  nums[i] != val {
         nums[int] = nums[i]
         int++
     } 
     }
return int
     }