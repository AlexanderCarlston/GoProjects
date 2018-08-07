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
/*			∞ n*2^n*(n!)^2
	π + 3 =  Σ ------------
			n=1   (2*n)!
*/
	pi := big.NewRat(0, 1)
	tempInt := big.NewInt(1)
	term := big.NewRat(0.0, 1.0)

	for k := 1; k < 4*n+10; k++ {
		top := big.NewInt(int64(k))
		tempInt = tempInt.Exp(big.NewInt(2), big.NewInt(int64(k)), nil)
		top = top.Mul(top, tempInt)
		tempInt = tempInt.Exp(tempInt.Set(factorialHelper(k)), big.NewInt(2), nil)
		top = top.Mul(top, tempInt)
		bottom := factorialHelper(2 * k)
		term := term.SetFrac(top, bottom)
		pi.Add(pi, term)
	}
	pi = pi.Add(pi, big.NewRat(-3, 1))
	fmt.Println("π to", n, "decimal places: ")
	fmt.Println(pi.FloatString(n))
}

func factorialHelper(a int) *big.Int {
	b := big.NewInt(1)

	for i := a; i > 1; i-- {
		b = b.Mul(b, big.NewInt(int64(i)))
	}

	return b
}