// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

package main

import "fmt"

func twoSum(input []int, target int)[]int{
	var result []int
	for i:=0; i<len(input); i++{
		for j:=i+1; j<len(input); j++{
			if input[i]+input[j] == target{
				result = append(result, i)
				result = append(result, j)
				return result
			}
		}
	}
	return result
}

func main(){
	input := []int{2,7,11,15}
	fmt.Println(twoSum(input, 9))
}