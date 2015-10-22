package wordcount

import (
  "strings"
)

func WordCount(word string) map[string]uint8 {
  result := map[string]uint8{}

  words := strings.Fields(word)
  for _, value := range words {
    result[value] += 1
  }
  return result
}
