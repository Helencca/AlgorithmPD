package solution2
import "sort"

func threeSum(nums []int) [][]int {
    //1.如若len(nums) 小于三，return
    if nums ==nil || len(nums) < 3 {
        return nil
    }

    //2.对数组进行排序
    sort.Ints(nums)

    //3.使用左右指针
    temp := make([][]int, 0,1)
    for i :=0; i<len(nums);i++{
        
        if nums[i] >0 {
            break
        }
        
        //去重i
        if i !=0 && nums[i] == nums[i-1]{
            continue
        }
        l := i+1
        r := len(nums)-1
    
        for l < r {
            sum := nums[i]+nums[l]+nums[r]
            if sum == 0{
                temp = append(temp, []int{nums[i],nums[l],nums[r]})
                //对 l,r去重
                    for l < r && nums[l] == nums[l+1]{
                            l++
                        } 
                    for l < r && nums[r] == nums[r-1]{
                            r--
                    } 
                
            l++
            r--
            } else if sum < 0{
                    l++
            } else{
                r--
            }
        } 
        
    }

    return temp
}