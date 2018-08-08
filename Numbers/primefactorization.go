//Have the user enter a number and find all Prime Factors (if there are any) and display them.
package main

import "fmt"

func main(){
	var n int

	fmt.Println("Compute prime factors of what number?")
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if n < 2 {
		fmt.Println("Error:", n, "is too small! Try an integer n ≥ 2.")
		return
	}

	
}