package main

import "fmt"

func calculateDiscount(prices []int32) int32 {

	if len(prices) > 0 {
		var totalCost int32 = 0
		for ind, price := range prices {
			if ind == 0 {
				totalCost += price
			} else if price > prices[ind-1] {
				newPrice := price - prices[ind-1]
				totalCost += newPrice
			} else {
				continue
			}
		}
		return totalCost
	} else {
		return 0
	}

}

func main() {
	sales := [][]int32{
		{1, 2, 3, 5, 2},
		{4, 5, 6, 2, 4},
		{3, 1, 5, 6, 3},
		{4, 5, 6},
	}

	expectedResults := []int32{5, 8, 8, 6}

	for ind, basket := range sales {
		actualResult := calculateDiscount(basket)
		fmt.Printf("Test[%v] - Expected Result: %v | Actual Result: %v\n", ind+1, expectedResults[ind], actualResult)

		if actualResult == expectedResults[ind] {
			fmt.Println("PASS")
		} else {
			fmt.Println("FAIL")
		}
	}
}
