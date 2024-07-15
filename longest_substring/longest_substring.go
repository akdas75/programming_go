func lengthOfLongestSubstring(s string) int {

    l := 0
    r := 0
    maxlen := 0

    n := len(s)
    m := make(map[byte]int)

    for (r < n) {
     
          if val,ok := m[s[r]]; ok {
            //fmt.Printf("Map hit %d : %d \n ", r , val)
              
              if l <= val {
                l = val + 1
              }
              m[s[l]] = l
          }

          m[s[r]] = r
          len := r - l + 1
          //fmt.Printf("l %d r %d len %d \n", l, r , len)
          if len > maxlen {
            maxlen = len
          }

          r++
    }

    return maxlen
    
}