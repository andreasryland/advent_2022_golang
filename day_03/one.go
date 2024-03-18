package main

import (
  "fmt"
  "bufio"
  "os"
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
  backno := 1
  for r.Scan() {
	line := r.Bytes()
    if len(line) > 0 {
      mid := len(line) / 2
      first := line[:mid]
      second := line[mid:]
      var bitmap uint64


      for i := 0; i < len(first); i++ {
	      mask := uint64(1) << getPriority(first[i])
	      bitmap |= mask
      }
      for i := 0; i < len(second); i++ {
	priority := getPriority(second[i])
	mask := uint64(1) << priority
	if bitmap & mask == mask {
		fmt.Printf("Found duplicate letter %c at index %d in backpack %d with priority %d\n", second[i], i, backno, priority)
		sum += priority
		break
	}
      }
      backno++
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

