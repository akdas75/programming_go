/*

var mem [][]bool;

func helper(stones []int,  lastIndex int, curIndex int,) bool {

    if (curIndex == len(stones) - 1) {
        return true
    }

    // fmt.Printf(" %d \n",len(mem))
    //fmt.Printf(" %d \n",len(mem[0]))

    fmt.Printf("lastIndex %d curIndex %d \n",lastIndex,  curIndex)
    if mem[lastIndex][curIndex] == true {
        return false
    }

    lastJump := stones[curIndex] - stones[lastIndex]
    nextIndex := curIndex + 1

    for (nextIndex < len(stones) && (stones[nextIndex] <= stones[curIndex] + lastJump + 1)) {

        nextJump := stones[nextIndex] - stones[curIndex]
        jump  :=   nextJump - lastJump

        if (jump >= -1 && jump <= 1) {
            if (helper(stones , curIndex, nextIndex)) == true {
                return true
            }
        }
        if jump > 1 {
            break
        }
        nextIndex++
    }

    mem[lastIndex][curIndex] = true

    return false

}


func canCross(stones []int) bool {

    n := len(stones)
    
    mem = make([][]bool, n+1)
    for i := range mem {
       mem[i] = make([]bool, n+1)
    }

    //fmt.Printf(" %d \n",len(mem))
    //fmt.Printf(" %d \n",len(mem[0]))

    if stones[1] != 1 {
        return false
    }

    return helper(stones , 0, 1)   
    
}

*/

func canCross(stones []int) bool {

    n := len(stones)

    sln := make([][]int, n+1)
    for i := range sln {
        if i == 0 {
           sln[i] = make([]int, 1)
        } else {
            sln[i] = make([]int, 0)
        }
          
    }

    fmt.Printf("sln %v \n", sln)

    for i := 1; i<len(stones); i++ {
        for j := 0; j<i; j++ {
            for _ , pos := range sln[j] {
                //fmt.Printf("i %d j %d pos %d \n", i, j, pos)

                if ((stones[j] + pos == stones[i])     ||
                   (stones[j] + pos - 1 == stones[i])  ||
                   (stones[j] + pos + 1 == stones[i])) {

                    //fmt.Printf("coming here")

                    sln[i] = append(sln[i], stones[i] - stones[j])
                    break;

                } 

            } 

        }
    }

    return len (sln[len(stones) - 1]) > 0

}    