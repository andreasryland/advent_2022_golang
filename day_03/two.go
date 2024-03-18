package main

import (
  "fmt"
  "bufio"
  "os"
  "math/bits"
//  "strconv"
)

func main() {
  fileName := "input.txt"
  //Section 1
  if len(os.Args) > 1 {
    fileName = os.Args[1] + ".txt"
  }
  fmt.Println("Using file ", fileName)
  file, err := os.Open(fileName)
  defer file.Close()
  if err != nil {
    panic("fail")
  }
  r := bufio.NewScanner(file)

  // Problem part
  //var deer []int
  sum := 0
  // Section 2
  groupCount := 1
  var first, second, third uint64
  for r.Scan() {
	line := r.Bytes()
    if len(line) > 0 {
	var currentMap *uint64
	switch groupCount {
		case 1:
			currentMap = &first
		case 2:
			currentMap = &second
		case 3:
			currentMap = &third
	}

      for i := 0; i < len(line); i++ {
	      mask := uint64(1) << getPriority(line[i])
	      *currentMap |= mask
      }
      if groupCount == 3 {
	      common := first & second
	common = common & third
	sum += bits.TrailingZeros64(common)
	first, second, third = 0, 0, 0
	groupCount = 1
      } else {
      	groupCount++
      }
    	if err != nil {
      	break
    	}
  	}
}
fmt.Printf("The total sum is: %d\n", sum);
}

func getPriority(letter byte) int {
	if letter > 'Z' {
		return int(letter) - 96
	} else {
	return int(letter) - 38
}
}

