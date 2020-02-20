package solution1

func maxArea(height []int) int {
    //1.最暴力的就是循环比较，从下标0开始，
    //如果a[0]<a[0+1],就不比较a[0],从a[0+1]开始
    Max := 0
    sum := 0
    n :=1
    len := len(height)
    for i:=0;i<len-n;i++{
      for n :=1; n <len-i; n++ {  
        if height[i] < height[i+n]{
            sum= height[i]*n
    
        }else{
             sum= height[i+n]*n
        }
        if sum > Max {
            Max = sum
            continue
        }
      }
    }
    return Max
}