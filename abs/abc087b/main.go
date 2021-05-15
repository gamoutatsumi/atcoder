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
  count := 0
  a, _ := strconv.Atoi(p[0][0])
  b, _ := strconv.Atoi(p[1][0])
  c, _ := strconv.Atoi(p[2][0])
  x, _ := strconv.Atoi(p[3][0])
  for i := 0; i < a+1; i++ {
    for j := 0; j < b+1; j++ {
      for k := 0; k < c+1; k++ {
        if ((i * 500 + j * 100 + k * 50) == x) {
          count++
        }
      }
    }
  }
  fmt.Println(count)
}
