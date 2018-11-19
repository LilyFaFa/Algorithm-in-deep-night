package main

import (
	"fmt"
	"sort"
)

func main() {
	var food int
	for {
		n, _ := fmt.Scan(&food)
		if n == 0 {
			break
		} else {
			var coupon int
			fmt.Scan(&coupon)

			foods := []int{}
			coupons := []int{}
			conMap := make(map[int]int)

			result := 0
			for i := 0; i < food; i++ {
				var f int
				fmt.Scan(&f)
				result += f
				foods = append(foods, f)
			}

			for i := 0; i < coupon; i++ {
				var c int
				fmt.Scan(&c)
				var d int
				fmt.Scan(&d)
				coupons = append(coupons, c)
				conMap[c] = d
			}

			sort.Ints(foods)
			sort.Ints(coupons)

			find(&result, conMap, coupons, foods, len(foods)-1, len(coupons)-1, result)
			fmt.Println(result)
		}
	}
}

func find(result *int, conMap map[int]int, coupons []int, foods []int, curfood int, curcoupon int, curresult int) {
	if curfood < 0 || curcoupon < 0 {
		if curresult < *result {
			*result = curresult
		}
		return
	}

	i := curfood
	for j := curcoupon; j >= 0; j-- {
		if coupons[j] <= foods[i] {
			cur := curresult - conMap[coupons[j]]
			find(result, conMap, coupons, foods, i-1, j-1, cur)
		} else {
			find(result, conMap, coupons, foods, i-1, j, curresult)
		}
	}

}
