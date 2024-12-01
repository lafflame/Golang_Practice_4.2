// Задача 4.4 из учебника "Язык программирования Go - Алан А.А. Донован, Брайан У. Керниган"
// Напишите версию функции rotate, которая работает в один проход

package main

import "fmt"

// Функция rotate с одним проходом
func rotate(nums []int) []int {
	j := len(nums) - 1
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[j] = nums[j], nums[i]
		j--
	}
	return nums
}

func main() {
	testCases := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8},
		{10, 20, 30, 40, 50},
		{7, 2, 9, 1},
	}

	// Вывод тест кейсов
	for _, testCase := range testCases {
		fmt.Println(rotate(testCase))
	}
}
