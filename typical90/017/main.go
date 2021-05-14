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
  p := StrStdin()
  fmt.Println(p)
}
