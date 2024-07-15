func reverse(x int) int {

    

  
    n := 0
    for (x != 0) {     
         
        rem := x % 10
        fmt.Printf("%d \n",rem)
        x = x /10
        n = n * 10 + rem
          
    }

    if n > math.MaxInt32 {
        n = 0
    }

    if n <  math.MinInt32 {
        n = 0
    }

    return n
    
}