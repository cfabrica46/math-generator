package utils

import (
	"fmt"
	"math"
	"math/rand"
)

func RandIntnLimit(max, min int) (num int) {
	//nolint:gosec
	return rand.Intn(max-min+1) + min
}

func GetMaxNum(maxCipher int) (maxNum int) {
	//nolint:gomnd
	return int(math.Pow(10, float64(maxCipher))) - 1
}

func GetMinNum(minCipher int) (minNum int) {
	//nolint:gomnd
	return int(math.Pow(10, float64(minCipher-1)))
}

func CheckMinCipher(max, min int) (err error) {
	if min > max {
		//nolint:goerr113
		return fmt.Errorf("min can't be mayor than max")
	}

	return nil
}
