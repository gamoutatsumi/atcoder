package main

import (
  "os"
  "bufio"
  "strings"
  "fmt"
)

func StrStdin() (stringInput string) {
  scanner := bufio.NewScanner(os.Stdin)

  scanner.Scan()
  stringInput = scanner.Text()

  stringInput = strings.TrimSpace(stringInput)
  return
}

func main() {
  var N, L, K int
  var A []int

  scanner := bufio.NewScanner(os.Stdin)

  fmt.Scan(&N, &L)
  fmt.Scan(&K)
  for scanner.Scan() {
    line := scanner.Text()
    chars := strings.Split(line, " ")
    for i, ch := range chars
  }
  fmt.Scan(&A)
  fmt.Printf("N: %d\nL: %d\nK: %d\nA:", N, L, K)
  fmt.Println(A)
}
