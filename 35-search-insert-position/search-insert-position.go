func searchInsert(nums []int, target int) int {
    var left, right, mid int
    var found bool = false
    left = 0
    right = len(nums) - 1

    for left <= right && !found {
         mid = (left+right)/2

        if target < nums[mid] {
            right = mid - 1
        } else if target > nums[mid] {
            left = mid + 1
        } else {
            found = true
        }
    }

    if nums[mid] == target {
        return mid
    } else {
        if target < nums[0] {
            return 0
        }else if target > nums[len(nums) - 1] {
            return len(nums)
        } else {
            return left
        }
    }
}