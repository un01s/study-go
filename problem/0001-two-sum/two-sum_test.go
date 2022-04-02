package twosum

import (
  "fmt"
  "reflect"
  "testing"
)

type test struct {
  input
  output
}

type input struct {
  nums []int
  target int
}

type output struct {
  answer []int
}

func TestTwoSum(t *testing.T) {
  t.Run("1", testTwoSumFunc(input{[]int{3,2,4}, 6}, output{[]int{1,2}}))
  t.Run("2", testTwoSumFunc(input{[]int{3,2,4}, 5}, output{[]int{0,1}}))
  t.Run("3", testTwoSumFunc(input{[]int{0,8,7,3,3,4,2}, 11}, output{[]int{1,3}}))
  t.Run("4", testTwoSumFunc(input{[]int{0,1}, 1}, output{[]int{0,1}}))
  //t.Run("5", testTwoSumFunc(input{[]int{0,3}, 5}, output{[]int{}}))
}

func testTwoSumFunc(i input, o output) func(*testing.T) {
  return func(t *testing.T) {
    actual := twoSum(i.nums, i.target)
    expected := o.answer
    if !reflect.DeepEqual(actual, expected) {
      t.Error(fmt.Printf("Expected %v but get %v!", expected, actual))
    }
  }
}
