package main

import (
  "os"
  "bufio"
  "strings"
  "fmt"
  "strconv"
)

func StrStdin() (stringInput [][]string) {
  b := true
  scanner := bufio.NewScanner(os.Stdin)

  for b {
    if (b == false) {
      break
    }
    b = scanner.Scan()
    t := strings.Split(scanner.Text(), " ")
    stringInput = append(stringInput, t)
  }
  return
}

func main() {
  p := StrStdin()
  a, _ := strconv.Atoi(p[0][0])
  b, _ := strconv.Atoi(p[0][1])
  c := a * b
  if (c % 2 == 0) {
    fmt.Printf("Even\n")
  } else {
    fmt.Printf("Odd\n")
  }
}
