package sqrt

import (
  "math"
  "fmt"
)

func Sqrt(number float64) (float64, error) {
  if (number < 0) {
    err := fmt.Errorf("Cannot Sqrt() negative number %d", int(number))
    return 0, err
  }

  sqrt := math.Sqrt(float64(number))
  return sqrt, nil
}

