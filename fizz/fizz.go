package fizz

import "strconv"

func FizzBuzz(number int) string {
  if EvenlyDivisble(3, number) && EvenlyDivisble(5, number) {
    return "fizzbuzz"
  }
  if EvenlyDivisble(3, number) {
    return "fizz"
  }
  if EvenlyDivisble(5, number) {
    return "buzz"
  }

  return strconv.Itoa(number)
}

func EvenlyDivisble(divisor int, dividend int) bool {
  return (dividend % divisor) == 0
}
