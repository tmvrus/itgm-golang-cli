package main

import (
	"testing"
)

func TestPrime(t *testing.T) {
	prime := []int{3, 7, 11, 19}
	notPrime := []int{1, 4, 6, 12, 20}

	for i := range prime {
		if !isPrime(prime[i]) {
			t.Errorf("%d expexted to be prime", prime[i])
		}
	}

	for i := range notPrime {
		if isPrime(notPrime[i]) {
			t.Errorf("%d not expexted to be prime", notPrime[i])
		}
	}
}
