var mem []int

func helper(n int, nums []int) int {

    if n == 0 {
        return nums[0]
    }

    if n < 0 {
        return 0
    }

    if mem[n] != -1 {
        return mem[n]
    }

    pick := nums[n] + helper(n-2, nums)
    noPick := 0 + helper(n-1, nums)

    res :=  math.Max (float64(pick), float64(noPick))

    mem [n] = int (res) 
    return int(res)

}


func rob(nums []int) int {

    n := len(nums)

    mem = make([]int, n)

    for i:=0; i<n; i++ {
        mem[i] = -1
    }
    
    return helper(n-1, nums)    
}

func rob(nums []int) int {

    /* Tabulation */  

    fmt.Println(nums)
    
    n := len(nums)
    mem = make([]int, n)

    mem[0] = nums[0]

    for i := 1; i<n; i++ {

        pick := 0
        noPick := 0
           
        if (i > 1) {
            pick = nums[i] + mem[i-2]
        } else {
            pick = nums[i]
        }
           
        noPick = 0 + mem[i-1]

        res :=  math.Max (float64(pick), float64(noPick))
        mem [i] = int (res) 
    }

    fmt.Println(mem)

    return mem[n-1]
}