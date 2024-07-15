
var dp [][]int

/*
func helper(m int, n int) int {
    if m == 0 && n == 0 {
        return 1
    }

    if m < 0 || n < 0 {
        return 0
    }

    if dp[m][n] != -1 {
         return dp[m][n]
    }     

    //fmt.Printf(" m %d n %d \n",m,n)
    
    up := helper(m-1, n) 
    left := helper(m, n-1)

    dp[m][n] = up + left

    return up + left
}

func uniquePaths(m int, n int) int {
   
    dp = make([][]int, m) 
    for i:=0; i<m; i++ {
        dp[i] = make([]int, n)
    }

    for i:=0; i<m; i++ {
        for j:=0; j<n; j++ {
            dp[i][j] = -1
        }
    }        
    return helper(m-1 , n-1)
    
}
*/

/* Tabulation */
/*
func uniquePaths(m int, n int) int {
    dp = make([][]int, m) 
    for i:=0; i<m; i++ {
        dp[i] = make([]int, n)
    }


    for i := 0; i<m; i++ {
        for j := 0; j<n; j++ {

            if i==0 && j ==0 {
                dp[i][j] = 1
                continue
            }

            left  := 0
            right := 0

            
            if i - 1 >= 0 {
                left  = dp[i-1][j] 
            }    
            if j - 1 >= 0 {
                right = dp[i][j-1] 
            }   

            dp[i][j] = left + right    

        }    
    }

    return dp[m-1][n-1]
} 
*/   

/* space optim */
func uniquePaths(m int, n int) int {
    
    prevRow := make([]int, n) 
 
    for i := 0; i<m; i++ {

        curRow := make([]int, n) 

        for j := 0; j<n; j++ {

            if i==0 && j ==0 {
                curRow[0] = 1
                continue
            }

            left  := 0
            right := 0

            /* left and up */
            if i - 1 >= 0 {
                left  = prevRow[j] 
            }    
            if j - 1 >= 0 {
                right = curRow[j-1] 
            }   

            curRow[j] = left + right
             

        }  

        prevRow = curRow  
    }

    return prevRow[n-1]
}  