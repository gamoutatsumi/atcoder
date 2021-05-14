package main

import (
  "os"
  "bufio"
  "strings"
  "fmt"
  "strconv"
)

func IntStdin() (integersInput [][]int) {
  b := true
  scanner := bufio.NewScanner(os.Stdin)

  for b {
    var integers []int
    b = scanner.Scan()
    if !b {
      break
    }
    t := strings.Split(scanner.Text(), " ")
    for _, v := range t {
      value, _ := strconv.Atoi(v)
      integers = append(integers, value)
    }
    integersInput = append(integersInput, integers)
  }
  

  return
}

func main() {
  p := IntStdin()
  p = p[1:]
  for i := range p {
    for j, v := range p[i] {
      p[i][j] += v
    }
  }
  fmt.Println(p)
}
