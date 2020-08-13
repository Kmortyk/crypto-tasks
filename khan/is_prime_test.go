package khan

import (
	"testing"
)

func TestIsPrime(t *testing.T) {
	is, steps := isPrime(39712334234567)
	t.Log(is, steps) // 6301770
}

func TestIsPrimeTwoCheck(t *testing.T) {
	is, steps := isPrimeTwoCheck(39712334234567)
	t.Log(is, steps) // 3150886
}

func TestIsPrimeEratosthenes(t *testing.T) {
	primes := isPrimeEratosthenes(39717)
	t.Log(primes)
}
