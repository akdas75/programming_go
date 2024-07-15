
var dp [][][]int

/*
func helper (i int , j1 int,  j2 int ,maxRows int, maxColumns int, grid [][]int) int {

    if j1 < 0  || j1 > maxColumns || j2 < 0  || j2 > maxColumns {
        return -100000000
    }

    if i == maxRows {
           
           if j1 != j2 {
                 return grid[i][j1] + grid[i][j2]
           }

           if j1 == j2 {
                return grid[i][j1]
           }
    }

    maxi := 0

    if dp[i][j1][j2] != -1 {
        return dp[i][j1][j2]
    }

    for dj1 :=-1; dj1<=1; dj1++ {

        for dj2 :=-1; dj2<=1; dj2++ {

            if j1 != j2 {
                maxi = int(math.Max(float64(maxi) ,                 
                            float64 (grid[i][j1] + 
                         grid[i][j2] + 
                         helper(i+1 , j1+dj1, j2+dj2, maxRows, maxColumns, grid))))
            }

            if j1 == j2 {                
                  maxi = int(math.Max(float64(maxi) ,                 
                            float64 (grid[i][j1] +                     
                         helper(i+1 , j1+dj1, j2+dj2, maxRows, maxColumns, grid))))
            }
        }    
    }
     
    dp[i][j1][j2]  = maxi
    return maxi
}
*/
/*
Robot1 i1 , j1
Robot2 i2 ,  j2
*/
/*
func cherryPickup(grid [][]int) int {

    maxRows := len(grid)
    maxColumns := len(grid[0])

    dp = make([][][]int, maxRows)
    for i:=0; i<maxRows; i++ {
        dp[i] = make([][]int, maxColumns)
        for j:=0; j<maxColumns; j++ {
           dp[i][j] = make([]int, maxColumns)
        }
    }

    for i:=0; i<maxRows; i++ {
        for j:=0; j<maxColumns; j++ {
            for k:=0; k<maxColumns; k++ {        
                dp[i][j][k] = -1
            }    
        }
    }    

    return helper(0, 0 , maxColumns - 1, maxRows -1, maxColumns -1, grid) 
    
}
*/

/* Tabulation */
func cherryPickup(grid [][]int) int {

    maxRows := len(grid)
    maxColumns := len(grid[0])

    dp = make([][][]int, maxRows)
    for i:=0; i<maxRows; i++ {
        dp[i] = make([][]int, maxColumns)
        for j:=0; j<maxColumns; j++ {
           dp[i][j] = make([]int, maxColumns)
        }
    }

    for i:=0; i<maxRows; i++ {
        for j:=0; j<maxColumns; j++ {
            for k:=0; k<maxColumns; k++ {        
                dp[i][j][k] = -1
            }    
        }
    }    

    /* base case */
    for j1:=0; j1<maxColumns; j1++ {
        for j2:=0; j2<maxColumns; j2++ {
            if j1 == j2 {
                dp[maxRows -1][j1][j2] = grid[maxRows-1][j1]
            } else {
                 dp[maxRows -1][j1][j2] = grid[maxRows-1][j1] + grid[maxRows-1][j2]
            }
        }
    }   

    for i:=maxRows-2; i>=0; i-- {
        for j1:=0; j1<maxColumns; j1++ {
            for j2:=0; j2<maxColumns; j2++ {
                
                maxi := -200000
                for dj1 :=-1; dj1<=1; dj1++ {
                    for dj2 :=-1; dj2<=1; dj2++ {
                        
                        value := 0
                        if j1 != j2 {
                            value = grid[i][j1] + grid[i][j2]
                        } else  {
                            value = grid[i][j1]
                        }   
                        
                        if j1+dj1 >=0 && j1+dj1 < maxColumns && j2+dj2 >=0 && j2+dj2 < maxColumns {
                            value += dp[i+1][j1+dj1][j2+dj2]
                        } else {
                            value += -200000
                        }

                        maxi = int(math.Max(float64(value), float64(maxi)))   
                            
                    }    
                }
                dp[i][j1][j2] = maxi
            }
        } 
    } 

    //fmt.Printf("%v \n",dp)

    return dp[0][0][maxColumns-1]      
    
}