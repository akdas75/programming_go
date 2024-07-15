package main

import (
"fmt"
"sort"
)


func threeSum_brute(nums []int) [][]int {

    m := make(map[string]struct{})
    res := make([][]int, 0)

    for i:=0; i<len(nums); i++ {

        for j:=i+1; j<len(nums); j++ {

            for k:=j+1; k<len(nums); k++ {

                //fmt.Printf("%d %d %d \n", nums[i] , nums[j] , nums[k])

                if(nums[i] + nums[j] + nums[k]) == 0 {

                    /* sort */
                    s := []int{nums[i], nums[j],nums[k]}
                    sort.Ints(s)
                    key := fmt.Sprintf("%s_%s_%s",s[0], s[1], s[2])

                    if _, ok := m[key]; !ok {
                        m[key] = struct{}{}

                        r1 := make([]int, 3)
                        r1[0] = nums[i]
                        r1[1] = nums[j]
                        r1[2] = nums[k]

                        res = append(res, r1)
                    }

                    
                }
            }
        }        

    }

    return res
    
}

func threeSum(nums []int) [][]int {
    res := [][]int{} 
    
    sort.Ints(nums) 
    
    for i := 0; i < len(nums)-2; i++ { 
        if(i == 0 || (i > 0 && nums[i] != nums[i-1])) { 
            low := i+1 
            high := len(nums)-1
            sum := 0-nums[i]
            
            for (low < high) { 
                if(nums[low] + nums[high] == sum) { 
                    res = append(res, []int{nums[i], nums[low], nums[high]}) 
                    for (low < high && nums[low] == nums[low+1]) {
                        low++ 
                    }      
                    
                    for(low < high && nums[high] == nums[high-1]) {
                         high-- 
                    }
                    low++
                    high--
                } else if (nums[low] + nums[high] > sum) {
                    high--
                } else {
                    low++
                }
            }
        }    
    }
    
    return res
    
}

func main() {

	nums := []int{-1,0,1,2,-1,-4}


	fmt.Println(threeSum_brute(nums))
	fmt.Println(threeSum(nums))

}	