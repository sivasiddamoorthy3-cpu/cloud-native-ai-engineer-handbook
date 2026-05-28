
package arrays

func TwoSum(nums []int, target int) []int {
    seen := map[int]int{}

    for i, n := range nums {
        if idx, ok := seen[target-n]; ok {
            return []int{idx, i}
        }
        seen[n] = i
    }

    return nil
}
