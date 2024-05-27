package main

import (
	"fmt"
)


func main() {
	 primes := []int{2, 3, 5, 7, 11, 13}
	 primes = append(primes, 17, 19, 23)

	 //Slice the slice from index 2 to index 5.
	 fmt.Println(primes[2:5])

     //Print out the slice.
	fmt.Println(primes)

	primeMaker := make([]int, 0, 5)
	for i := 0; i < len(primes); i++ {
		primeMaker = append(primeMaker, primes[i])
    }
	fmt.Println(primeMaker)
	
	
}