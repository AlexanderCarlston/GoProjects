// Instructions: find pi to the Nth digit.

package main

import (
	"fmt"
	"math/big"
)

func main(){
	var n int
	// n is the NTH digit the user inputs
	fmt.Printf("Compute π to how many decimal places? (≤200) ")
	_, err := fmt.Scanf("%d", &n)
	//Error handling
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if n > 200 {
		fmt.Println("Error:", n, "> 200 - too many decimal places!")
		return
	} else if n < 0 {
		fmt.Println("Error:", n, "< 0")
		return
	}

	

}

func factorialHelper(a int) *big.Int {
	b := big.NewInt(1)

	for i := a; i > 1; i-- {
		b = b.Mul(b, big.NewInt(int64(i)))
	}

	return b
}