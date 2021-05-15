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
  var a []int
  count := 0
  p := StrStdin()
  for _, v := range p[1] {
    n, _ := strconv.Atoi(v)
    a = append(a, n)
  }
  loop:
    for {
      for i, v := range a {
        if (v % 2 == 0) {
          a[i] = v / 2
        } else {
          break loop
        }
      }
      count++
    }
  fmt.Println(count)
}
