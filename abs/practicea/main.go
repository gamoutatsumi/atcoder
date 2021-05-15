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
  b, _ := strconv.Atoi(p[1][0])
  c, _ := strconv.Atoi(p[1][1])
  s := p[2][0]
  fmt.Printf("%d %s\n", a + b + c, s)
}
