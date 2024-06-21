package main

import (
"fmt"
)

func twoSum(nums []int, target int) []int {

    m := make(map[int]int)
    res := make([]int, 2)

    for i := 0; i<len(nums); i++ {
     
          diff := nums[i] - target
          diff = diff * (-1)        

          if val,ok := m[diff]; ok {
               res[0] = val
               res[1] =  i
               break
          } else {
            m[nums[i]]=i
          }

    }
    return res    
}

func main() {

	nums := []int{2,7,5,1}
	target := 9

	fmt.Println(twoSum(nums, target))

}	