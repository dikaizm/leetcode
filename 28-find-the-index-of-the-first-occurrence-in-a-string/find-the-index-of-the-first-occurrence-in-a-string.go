func strStr(haystack string, needle string) int {
    var n, m, i int

    n = len(haystack)
    m = len(needle)

    for i = 0; i <= n-m; i++ {
        if haystack[i:i+m] == needle {
            return i
        }
    }

    return -1
}