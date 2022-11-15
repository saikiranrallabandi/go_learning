/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target
Example: Given nums = [2, 7, 11, 15], target = 9,
*/

package main
import (
	"fmt"
)

func main() {
	twoSums([1, 3, 5, 9 ])
}

func twoSums(nums []int, target int)  []int {
	 t := make(map[int]int, len(nums))
	 for i, v := range nums{
		      if vv, ok := t[v]; ok {
				        return []int{vv, i}
			  }
			  t[target-v] = i
	 }
	 return nil

}