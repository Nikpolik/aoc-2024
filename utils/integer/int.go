package integer

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(fmt.Sprintf("utils/int/SafeAtoi: Cannot convert %s to int, crashing", s))
	}

	return i
}

func Abs(i int) int {
	return int(math.Abs(float64(i)))
}
