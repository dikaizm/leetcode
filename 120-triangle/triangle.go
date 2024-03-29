func minimumTotal(triangle [][]int) int {
    n := len(triangle)

    for i := n - 2; i >= 0; i-- {
        m := len(triangle[i])
        for j := 0; j < m; j++ {
            triangle[i][j] += min(triangle[i+1][j], triangle[i+1][j+1])
        }
    }

    return triangle[0][0]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
