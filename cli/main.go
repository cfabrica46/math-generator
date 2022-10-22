package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"strconv"
	"time"
)

type Multiples struct {
	x int
	y int
}

func main() {
	var nOperaciones int
	var multiples []Multiples

	fmt.Print("Ingrese el numero de operaciones que desea realizar: ")
	fmt.Scan(&nOperaciones)

	if nOperaciones <= 0 {
		log.Fatalf("numero invalido: %d", nOperaciones)
	}

	fmt.Println()
	fmt.Println("---")
	fmt.Println()

	// ---

	firstMaxMultiple, nFirstMaxNumber, err := requestMaxNum()
	if err != nil {
		log.Fatal(err)
	}

	// ---

	firstMinMultiple, _, err := requestMinNum(nFirstMaxNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	// ---

	secondMaxMultiple, nSecondMaxNumber, err := requestMaxNum()
	if err != nil {
		log.Fatal(err)
	}

	// ---

	secondMinMultiple, _, err := requestMinNum(nSecondMaxNumber)
	if err != nil {
		log.Fatal(err)
	}

	// ---

	rand.Seed(time.Now().UnixNano())

	fmt.Println()
	fmt.Println("OPERACIONES")
	fmt.Println()

	for i := 0; i < nOperaciones; i++ {

		multipleOne := randIntnLimit(firstMaxMultiple, firstMinMultiple)
		multipleTwo := randIntnLimit(secondMaxMultiple, secondMinMultiple)

		fmt.Printf("%d) %d x %d\n", i+1, multipleOne, multipleTwo)
		fmt.Println("---")

		multiples = append(multiples, Multiples{x: multipleOne, y: multipleTwo})

	}

	fmt.Println()
	fmt.Println("Enter para continuar ...")
	fmt.Scanln()

	fmt.Println()
	fmt.Println("RESULTADOS")
	fmt.Println()

	for i := range multiples {

		fmt.Printf("%d) %d * %d\n", i+1, multiples[i].x, multiples[i].y)

		fmt.Printf("result: %d\n", extendMultiply(multiples[i].x, multiples[i].y))

		fmt.Println("---")
		fmt.Println()
	}
}

func requestMaxNum() (maxNum, nMaxNum int, err error) {
	fmt.Print("Ingrese el numero maximo de cifras del multiplo: ")
	fmt.Scan(&nMaxNum)

	if nMaxNum <= 0 {
		return 0, 0, fmt.Errorf("numero invalido: %d", nMaxNum)
	}

	return int(math.Pow(10, float64(nMaxNum))) - 1, nMaxNum, nil
}

func requestMinNum(nMaxNum int) (maxNum, nMinNum int, err error) {
	fmt.Print("Ingrese el numero minimo de cifras del multiplo: ")
	fmt.Scan(&nMinNum)

	if nMinNum <= 0 {
		return 0, 0, fmt.Errorf("numero invalido: %d", nMinNum)
	}

	if nMinNum > nMaxNum {
		return 0, 0, fmt.Errorf("numero invalido: el valor minimo: %d, no puede ser mayor a: %d", nMinNum, nMaxNum)
	}

	return int(math.Pow(10, float64(nMinNum-1))), nMinNum, nil
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

	for i := range numsY {
		res := (x * numsY[i]) * int(math.Pow(10, float64(i)))
		fmt.Printf("%d * %d = %d\n", x, numsY[i], res)

		result += res
	}

	return result
}

func randIntnLimit(max, min int) (num int) {
	return rand.Intn(max-min+1) + min
}
