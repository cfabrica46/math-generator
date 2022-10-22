package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	x := 163
	y := 354

	// obtain unidades, decenas, centenas
	fmt.Println("unidades: ", x%10)
	fmt.Println("decenas: ", (x%100)/10)
	fmt.Println("centenas: ", (x%1000)/100)

	xString := strconv.Itoa(x)
	fmt.Println("numero de cifras: ", len(xString))

	fmt.Println()
	// ---

	fmt.Println("unidades: ", y%10)
	fmt.Println("decenas: ", (y%100)/10)
	fmt.Println("centenas: ", (y%1000)/100)

	yString := strconv.Itoa(y)
	fmt.Println("numero de cifras: ", len(yString))

	fmt.Println()
	// ---

	/* numsX := descomposeNum(x)
	numsY := descomposeNum(y) */

	res := extendMultiply(x, y)
	fmt.Println("result: ", res)
	fmt.Println("result: ", x*y)
}

func descomposeNum(num int) (nums []int) {
	numString := strconv.Itoa(num)
	nNumber := len(numString)

	for i := 0; i < nNumber; i++ {
		x := int(math.Pow(10, float64(i+1)))
		y := int(math.Pow(10, float64(i)))

		nums = append(nums, (num%x)/y)
	}

	return nums
}

func extendMultiply(x, y int) (result int) {
	numsY := descomposeNum(y)

	fmt.Printf("%d * %d\n", x, y)

	for i := range numsY {
		res := (x * numsY[i]) * int(math.Pow(10, float64(i)))
		fmt.Printf("%d * %d = %d\n", x, numsY[i], res)

		result += res
	}

	return result
}
