package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"
)

var (
	firstMaxMultiple  int
	firstMinMultiple  int
	secondMaxMultiple int
	secondMinMultiple int
)

func main() {
	var nOperaciones, nFirstMaxNumber, nFirstMinNumber, nSecondMaxNumber, nSecondMinNumber int
	var results []int

	fmt.Print("Ingrese el numero de operaciones que desea realizar: ")
	fmt.Scan(&nOperaciones)

	if nOperaciones <= 0 {
		log.Fatalf("numero invalido: %d", nOperaciones)
	}

	fmt.Println()
	fmt.Println("---")
	fmt.Println()

	// ---

	fmt.Print("Ingrese el numero maximo de cifras de el primer multiplo: ")
	fmt.Scan(&nFirstMaxNumber)

	if nFirstMaxNumber <= 0 {
		log.Fatalf("numero invalido: %d", nFirstMaxNumber)
	}

	firstMaxMultiple = int(math.Pow(10, float64(nFirstMaxNumber))) - 1

	// ---

	fmt.Print("Ingrese el numero minimo de cifras de el primer multiplo: ")
	fmt.Scan(&nFirstMinNumber)

	if nFirstMinNumber <= 0 {
		log.Fatalf("numero invalido: %d", nFirstMinNumber)
	}

	if nFirstMinNumber > nFirstMaxNumber {
		log.Fatalf("numero invalido: el valor minimo: %d, no puede ser mayor a: %d", nFirstMinNumber, nFirstMaxNumber)
	}

	firstMinMultiple = int(math.Pow(10, float64(nFirstMinNumber-1)))

	fmt.Println()
	// ---

	fmt.Print("Ingrese el numero maximo de cifras de el segundo multiplo: ")
	fmt.Scan(&nSecondMaxNumber)

	if nSecondMaxNumber <= 0 {
		log.Fatalf("numero invalido: %d", nSecondMaxNumber)
	}

	secondMaxMultiple = int(math.Pow(10, float64(nSecondMaxNumber))) - 1

	// ---

	fmt.Print("Ingrese el numero minimo de cifras de el primer multiplo: ")
	fmt.Scan(&nSecondMinNumber)

	if nSecondMinNumber <= 0 {
		log.Fatalf("numero invalido: %d", nFirstMaxNumber)
	}

	if nSecondMinNumber > nSecondMaxNumber {
		log.Fatalf("numero invalido: el valor minimo: %d, no puede ser mayor a: %d", nFirstMinNumber, nFirstMaxNumber)
	}

	secondMinMultiple = int(math.Pow(10, float64(nSecondMinNumber-1)))

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

		results = append(results, multipleOne*multipleTwo)
	}

	fmt.Println()
	fmt.Println("RESULTADOS")
	fmt.Println()

	for i := range results {
		fmt.Printf("%d) %d\n", i+1, results[i])
		fmt.Println("---")
	}
}

func randIntnLimit(max, min int) (num int) {
	return rand.Intn(max-min+1) + min
}
