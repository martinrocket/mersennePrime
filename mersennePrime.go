package main
/* In mathematics, a Mersenne prime is a prime number that is one less than a power of two. 
That is, it is a prime number of the form Mn = 2n − 1 for some integer n. They are named after 
Marin Mersenne, a French Minim friar, who studied them in the early 17th century. If n is a 
composite number then so is 2n − 1. Therefore, an equivalent definition of the Mersenne primes is 
that they are the prime numbers of the form Mp = 2p − 1 for some prime p.*/

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	fmt.Println("\n")
	fmt.Println("Mersenned Prime\n")

	//Create a slice with 10000 integers

	var x, y []int

	//for i := 2; i <= 100; i++ {
	for i := 2; i <= 100000; i++ {
		x = append(x, i)

	}

	//fmt.Println(x)

	for j := len(x); j > 0; j-- {
		for i := 0; i < j; i++ {
			//fmt.Printf("%d / %d = %d %d \n", x[j-1], x[i], x[j-1]/x[i], x[j-1]%x[i])
			if x[j-1]%x[i] == 0 && x[j-1]/x[i] != 1 {
				//fmt.Printf(" break %d \n", x[j-1])
				break
			} else {
				if x[j-1]/x[i] == 1 {
					y = append(y, x[j-1])
					break
				}
			}
		}
	}
	slices.Reverse(y)

	var z []float64
	for i := 1; i < 100; i++ {
		z = append(z, math.Pow(2.0, float64(i))-1)

	}

	for i := range y {
		for j := range z {
			if float64(y[i]) == z[j] {
				fmt.Println(y[i])
			}
		}

	}
}
