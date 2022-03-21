// https://leetcode.com/problems/matrix-diagonal-sum

func diag(mat[][] int) int{
    
    i := 0
    j := 0
    
    n := len(mat)
    ret := 0
    for i < n && j < n {
        
        ret = ret + mat[i][j]
        i++
        j++
    }
    
    return ret
    
}

func antidiag(mat[][] int) int{
    
    i := 0
    j := 0
    
    n := len(mat)
    j = n - 1
    ret := 0
    for i < n && j > -1 {
        
        ret = ret + mat[i][j]
        i++
        j--
    }
    
    return ret
    
}


func diagonalSum(mat [][]int) int {
    n := len(mat)
    return diag(mat) + antidiag(mat) - mat[n / 2][n / 2]
    
}