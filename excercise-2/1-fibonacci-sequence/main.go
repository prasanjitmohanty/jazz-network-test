package main

import (
	"fmt"
	"math/big"
)

func main() {
	f := fibonacci()
	totalNum := 0
	var sum  = big.NewInt(0)
	for totalNum < 100 {
		num := f()
		if num.Uint64() %2 == 0 {
			fmt.Println(num)
			sum.Add(sum,num)
			totalNum++
		}
	}

	fmt.Println("SUM of first 100 even fibonacci numbers:", sum)
}

func fibonacci() func() *big.Int {
	var a, b = big.NewInt(1), big.NewInt(1)
	return func() *big.Int {
		a, b = b, b.Add(a,b)
		return a
	}
}
