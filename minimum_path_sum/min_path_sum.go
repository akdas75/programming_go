

var dp [][]int
func helper(m int, n int, grid [][]int) int {

    if m < 0 || n < 0 {
        return math.MaxInt32
    }
    if m == 0 && n == 0 {
        return  grid[0][0]
    }   

    if dp[m][n] != -1 {
         return dp[m][n]
    }     

    //fmt.Printf(" m %d n %d \n",m,n)
    
    up := grid[m][n] + helper(m-1, n, grid) 
    left := grid[m][n] + helper(m, n-1, grid)

    dp[m][n] = int(math.Min (float64(up), float64(left)))

    return dp[m][n]
}


func minPathSum(grid [][]int) int {
     m := len(grid)
    n := len(grid[0])

    fmt.Printf("m %d n %d \n",m,n);

    dp = make([][]int, m) 
    for i:=0; i<m; i++ {
        dp[i] = make([]int, n)
    }

    for i:=0; i<m; i++ {
        for j:=0; j<n; j++ {
            dp[i][j] = -1
        }
    } 

     return helper(m-1 , n-1, grid)          
    
}