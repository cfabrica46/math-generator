package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 539

	// obtain unidades, decenas, centenas
	fmt.Println("unidades: ", x%10)
	fmt.Println("decenas: ", (x%100)/10)
	fmt.Println("centenas: ", (x%1000)/100)

	xString := strconv.Itoa(x)
	fmt.Println("numero de cifras: ", len(xString))
}
