var dp [][]int

func helper(m int, n int, obstacleGrid [][]int) int {

    if m < 0 || n < 0 {
        return 0
    }


    if obstacleGrid [m][n] == 1 {
        return 0
    }

    if m == 0 && n == 0 {
        return 1
    }

    
    

    if dp[m][n] != -1 {
         return dp[m][n]
    }     

    //fmt.Printf(" m %d n %d \n",m,n)
    
    up := helper(m-1, n, obstacleGrid) 
    left := helper(m, n-1, obstacleGrid)

    dp[m][n] = up + left

    return up + left
}


func uniquePathsWithObstacles(obstacleGrid [][]int) int {

    m := len(obstacleGrid)
    n := len(obstacleGrid[0])

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
    return helper(m-1 , n-1, obstacleGrid)   
    
}


