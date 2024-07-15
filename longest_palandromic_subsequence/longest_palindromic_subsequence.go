
func max(a int, b int) int {
    
    if a > b{
        return a
    }
    
    return b
}

/*

func longestPalindrome(s string) string {
    
    count := 0
    
    b1 := []byte(s)
    b2 := []byte(s)
    
    table := make([][]bool, len(s))
    for i:=0; i<len(s); i++{
        table[i] = make([]bool, len(s))
    }
  
    l := 0
    x := 0
    y := 0
    for g:=0 ; g<len(s); g++ {
        
        for i ,j := 0 ,g; j<len(s); i,j = i+1 ,j + 1 {
            
            if g == 0 {
                table[i][j] = true
            } else if g == 1 {
                // 2 characters only 
                if b1[i] == b2[j] {
                    table[i][j] = true
                } else {
                     table[i][j] = false
                }
            } else {
                if b1[i] == b2[j]  && table[i+1][j-1]{
                    table[i][j] = true
                } else {
                     table[i][j] = false
                }
            }
            
            if table[i][j] == true {
                count++
                if l < j-i+1 {
                   l = max(l, j-i+1)
                   x = i
                   y = j 
                }    
                
            }
            
        }
        
    }
    
    fmt.Printf ("Number of palindromes %d \n",count)
    fmt.Printf ("longest palindrome length %d \n",l)
    fmt.Printf ("longest palindrome x %d  y %d : %s\n",x,y, string(b1[x:y+1]))
    return string(b1[x:y+1])
}

*/

func checkIfPalindrome(i int, j int , s string) bool{

    for i <= j {
        if s[i] != s[j] {
            break
        }
        i++
        j--
    }

    if i>j {
        return true
    }
    return false
}

/*
func longestPalindrome(s string) string {

    l := len(s)
    maxi := 0
    startIndex := 0
    endIndex := 0
    for i:=0; i<l; i++ {
        for j:=i; j<l; j++ {

           
           isPalindrome :=  checkIfPalindrome(i,j, s)
           //fmt.Printf("i %d j %d isPalindrome %v \n", i, j, isPalindrome) 

           if isPalindrome {
                len := j-i
                if len > maxi {
                    maxi = len
                    startIndex = i
                    endIndex = j
               }
           }   
        }    

    }

    //fmt.Printf("startIndex %d endIndex %d \n",startIndex, endIndex)

    return s[startIndex : endIndex+1]

}  

*/

func solve(i int , j int , s string) int {

    if i > j {
        return 1
    }

    if dp[i][j] != -1 {
        return dp[i][j]
    }

    if s[i] != s[j] {
        return 0
    }

    dp[i][j] = solve(i+1, j-1, s)
    
    return dp[i][j]
}

var dp [][]int

func longestPalindrome(s string) string {

    l := len(s)

    dp = make([][]int, l)
    for i:=0; i<l; i++ {
        dp[i] = make([]int, l)
    } 

    for i:=0; i<l; i++ {
        for j:=0; j<l; j++ {
            dp[i][j] = -1
        }
    }     

    maxi := 0
    startIndex := 0
    endIndex := 0
    for i:=0; i<l; i++ {
        for j:=i; j<l; j++ {
                       
           isPalindrome :=  solve(i,j, s)
           dp[i][j] = isPalindrome
           //fmt.Printf("i %d j %d isPalindrome %v \n", i, j, isPalindrome) 

           if isPalindrome == 1 {
                len := j-i
                if len > maxi {
                    maxi = len
                    startIndex = i
                    endIndex = j
               }
           }   
        }    

    }

    //fmt.Printf("startIndex %d endIndex %d \n",startIndex, endIndex)

    return s[startIndex : endIndex+1]

}    