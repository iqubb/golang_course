//https://leetcode.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    m[nums[0]] = 0
    for i := 1; i < len(nums); i++ {
        if value, ok := m[target-nums[i]]; ok {
            return []int {i, value}
        } else {
            m[nums[i]] = i
        }
    }
    return []int {0, 0}
}
