// Enter a number and have the program generate the Fibonacci sequence to that number or to the Nth number.
//Fibonacci sequence: a series of numbers in which each number ( Fibonacci number ) is the sum of the two preceding numbers. The simplest is the series 1, 1, 2, 3, 5, 8, etc.
package main

import "fmt"

func main(){
	var n int
	// n is the number the user puts in
	fmt.Println("Generate Fibonacci sequence to what number? ")
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if n < 1 {
		fmt.Println("Error:", n, "is too low - no sequence can be generated.")
		return
	} else if n > int(^uint(0)>>1) {
		fmt.Println("Error:", n, "is too big!  Try a smaller number.")
		return
	}

	fmt.Println("Fibonacci sequence to", n)
	fmt.Print("0, 1, ")

	// Seed with 0, 1
	p := 0
	q := 1
	for p+q <= n {
		if p < q {
			p += q
			fmt.Printf("%d", p)
		} else {
			q += p
			fmt.Printf("%d", q)
		}

		// Leave off the comma on the last entry
		if p+q <= n {
			fmt.Printf(", ")
		} else {
			fmt.Println()
		}
	}
}