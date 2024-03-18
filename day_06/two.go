package main

import (
  "fmt"
  "io/ioutil"
  "os"
)

func main() {
  fileName := "input.txt"
  //Section 1
  if len(os.Args) > 1 {
    fileName = os.Args[1] + ".txt"
  }
  fmt.Println("Using file ", fileName)
  c, err := ioutil.ReadFile(fileName)
  if err != nil {
    panic("fail")
  }
  result := 0

  for i := 0; i < len(c); i++ {
    if charsAreUnique(c[i : i+14]) {
      result = i + 14
      break
    }
  }
  
fmt.Printf("The answer is: %d\n", result)
}

func charsAreUnique(chars []byte) bool {
  var bitmap uint32 = 0
  for _, b := range chars {
    pos := b - 'a'
    if bitmap & (1<<pos) > 0 {
	    return false
    }
    bitmap |= (1 << pos)
  }
  return true
}


