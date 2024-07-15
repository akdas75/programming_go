
var dp [][]int

/*

func helper(i int, j int, n int, triangle [][]int) int {

    //fmt.Printf("i %d \n", i)

    if i == n-1 {
        return triangle[i][j]
    }

    if dp[i][j]  != -1 {
        return dp[i][j]
    }

    d := triangle[i][j] + helper(i+1, j, n, triangle)
    dg := triangle[i][j] + helper(i+1, j+1, n, triangle)

    dp[i][j] =  int(math.Min(float64(d), float64(dg)))

    return dp[i][j]
}


func minimumTotal(triangle [][]int) int {

    m := len(triangle)
    n := len(triangle[len(triangle) - 1])

    dp = make([][]int, len(triangle))
    for i:=0; i<m; i++ {
        dp[i] = make([]int, n)
    }

    for i:=0; i<m; i++ {
        for j:=0; j<n; j++ {
            dp[i][j] = -1
        }
    }        

    fmt.Printf("length %d \n", len(dp))
    fmt.Printf("length %d \n", len(dp[0]))
    fmt.Printf("length %v \n", dp)
    return helper(0 , 0, len(triangle), triangle)
    
}

*/

/* Tabulation */
func minimumTotal(triangle [][]int) int {

    m := len(triangle)
    n := len(triangle[len(triangle) - 1])

    dp = make([][]int, len(triangle))
    for i:=0; i<m; i++ {
        dp[i] = make([]int, n)
    }

    for i:=0; i<m; i++ {
        for j:=0; j<n; j++ {
            dp[i][j] = -1
        }
    }        

    for j:=0; j<n; j++ { 
        dp[n-1][j] = triangle[n-1][j]
    }   

    for i:=n-2; i>=0; i-- {
       for j:=0; j<=i; j++ { 
        //fmt.Printf("i %d j %d \n",i,j)
            d  := triangle[i][j] + dp[i+1][j]
           dg  := triangle[i][j] + dp[i+1][j+1]
           dp[i][j] =  int(math.Min(float64(d), float64(dg)))

       }     
    }     
    
    //fmt.Printf ("dp %d \n",dp)
    return dp[0][0];
}