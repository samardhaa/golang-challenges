package main

import "fmt"

func main() {
	fmt.Println("Two sum array index problem")

	numbers := []int{1, 2, 3, 4, 5}
	target := 7

	result := twoSum(numbers, target)

	fmt.Println("Two sum of indexes in array that adds upto target :", result)

}

func twoSum(numb []int, target int) []int {

	num := make(map[int]int)

	for index, value := range numb {
		complement := target - value

		if i, ok := num[complement]; ok {
			return []int{i, index}
		}

		num[value] = index
		fmt.Printf("Values captured in map in iteration %d : %d\n", index, num)
	}

	return []int{}
}
