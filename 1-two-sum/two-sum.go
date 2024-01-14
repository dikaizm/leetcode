func twoSum(nums []int, target int) []int {
    var cart, result []int
    for i := 0; i < len(nums); i++ {
        cart = []int{nums[i]}
        for j := 0; j < len(nums); j++ {
            if cart[0] + nums[j] == target && i != j {
                if i > j {
                    result = []int{j,i}
                    break
                }
                result = []int{i,j}
                break
            }
        }
        cart = []int{}
    }
    return result
}

/*
- check array length
- pick first index
- use first index to add for the rest of the indices
- loop thru array
- if target has been found, return index i and index j as array
- else pick next index, repeat the same process
*/