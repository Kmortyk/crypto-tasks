package khan

import (
	"math"
)

func isPrime(num int) (bool, int) {
	steps := 0

	sqrt := math.Sqrt(float64(num))
	for i := 2; i < int(sqrt); i++ {
		steps++
		if num%i == 0 {
			return false, steps
		}
	}
	return true, steps
}

func isPrimeTwoCheck(num int) (bool, int) {
	steps := 1

	if num%2 == 0 {
		// if 2 return true (prime) otherwise return false
		return num == 2, steps
	}

	sqrt := math.Sqrt(float64(num))
	for i := 3; i < int(sqrt); i += 2 {
		steps++
		if num%i == 0 {
			return false, steps
		}
	}
	return true, steps
}

func isPrimeSieve(num int) (bool, int) {
	steps := 1
	isComposite := make([]int, num+1)

	if num < 2 {
		return false, 0
	}

	for m := 2; m*m <= num; m++ {
		steps++
		if isComposite[m] != 1 {
			// check if m divides upperBound
			if num%m == 0 {
				// found prime divisor, therfore composite!
				return false, steps
			}
			// now mark off all composites starting at m^2
			// which are multiples of m
			for z := m * m; z*z <= num; z += m {
				steps++
				// mark position z as composite (1)
				isComposite[z] = 1
			}
		}
	}
	return true, steps
}

func sieveEratosthenes(upperBound int) []int {
	if upperBound < 2 {
		return nil
	}

	ubf := float64(upperBound + 10000)
	primes := make([]int, int(math.Ceil(ubf/math.Log(ubf)))) // ???
	isComposite := make([]int, upperBound)

	isComposite[0] = 1
	isComposite[1] = 1

	// loop from 2 to sqrt(upperBound)
	for m := 2; m*m < upperBound; m++ {
		// if prime (ie !== 1)
		if isComposite[m] != 1 {
			// now mark off all multiples starting at m^2
			for z := m * m; z < upperBound; z += m {
				// mark position z as composite (ie. === 1)
				isComposite[z] = 1
			}
		}
	}

	p := 0
	// print all primes by scanning array
	for h := 0; h < upperBound; h++ {
		// when you find a unmarked number
		if isComposite[h] != 1 {
			// put it in the prime array
			primes[p] = h
			// increment to next cell in array
			p++
		}
	}
	return primes
}
