package twosum

// switch key and value in a new map
func twoSum(nums []int, target int) []int {
  m := make(map[int]int)
  for k, v := range nums {
    if idx, ok := m[target-v]; ok {
      return []int{idx, k}
    }
    m[v] = k
  }
  return nil
}

// brutal force
func two_sum(nums []int, target int) ([2]int) {
  var r [2]int

  for i, n := range nums {
    for j := i+1; j < (len(nums)-1); j++ {
      if (n + nums[j] == target) {
        r[0] = i
        r[1] = j
        return r
      }
    }
  }

  return r
}
