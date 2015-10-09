package fizz

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNoFizzBuzz(t *testing.T) {
	assertFizzBuzzIs(1, "1", t)
	assertFizzBuzzIs(2, "2", t)
	assertFizzBuzzIs(4, "4", t)
	assertFizzBuzzIs(7, "7", t)
	assertFizzBuzzIs(8, "8", t)
	assertFizzBuzzIs(11, "11", t)
}

func TestFizzNoBuzz(t *testing.T) {
	assertFizzBuzzIs(3, "fizz", t)
	assertFizzBuzzIs(6, "fizz", t)
	assertFizzBuzzIs(9, "fizz", t)
	assertFizzBuzzIs(12, "fizz", t)
	assertFizzBuzzIs(18, "fizz", t)
}

func TestBuzzNoFizz(t *testing.T) {
	assertFizzBuzzIs(5, "buzz", t)
	assertFizzBuzzIs(20, "buzz", t)
	assertFizzBuzzIs(25, "buzz", t)
}

func TestFizzBuzz(t *testing.T) {
  assertFizzBuzzIs(15, "fizzbuzz", t)
}

func assertFizzBuzzIs(number int, expected string, t *testing.T) {
	assert.Equal(t, FizzBuzz(number), expected)
}
