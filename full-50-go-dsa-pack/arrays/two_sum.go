
package arrays

func TwoSum(nums []int, target int) []int {
    seen := map[int]int{}

    for i, num := range nums {
        if idx, ok := seen[target-num]; ok {
            return []int{idx, i}
        }

        seen[num] = i
    }

    return nil
}
