func maxArea(height []int) int {
    var max, curr, w int

    i, j := 0, len(height)-1

    for i < j {
        w = j - i

        if height[i] < height[j] {
            curr = height[i] * w
            i++
        } else {
            curr = height[j] * w
            j--
        }

        if curr > max {
            max = curr
        }
    }

    return max
}