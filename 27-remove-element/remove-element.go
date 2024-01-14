func removeElement(nums []int, val int) int {
    i := 0
    last := len(nums) - 1

    for i <= last {
        if nums[i] == val {
            nums[i], nums[last] = nums[last], nums[i]
            last--
        } else {
            i++
        }
    }

    return last + 1
}