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
    t := strings.Split(scanner.Text(), "")
    stringInput = append(stringInput, t)
  }
  return
}

func main() {
  count := 0
  p := StrStdin()
  a, _ := strconv.Atoi(p[0][0])
  b, _ := strconv.Atoi(p[0][1])
  c, _ := strconv.Atoi(p[0][2])
  if (a == 1) {
    count += 1
  }
  if (b == 1) {
    count += 1
  }
  if (c == 1) {
    count += 1
  }
  fmt.Println(count)
}
