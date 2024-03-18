package main

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
)

func main() {
  //Section 1
file, err := os.Open("input.txt")
defer file.Close()
if err != nil {
  panic("fail")
}
r := bufio.NewReader(file)

var deer []int
sum := 0
// Section 2
for {
  line, _, err := r.ReadLine()
  
  if len(line) > 0 {
    lineStr := string(line)
    number, err := strconv.Atoi(lineStr)
    if err != nil {
      panic("bad number")
    }
    sum += number
  } else if (sum > 0) {
    deer = append(deer, sum)
    sum = 0
  }
  if err != nil {
    break
  }
 }
 if (sum > 0) {
   deer = append(deer, sum)
 }
  maxA := 0
  maxB := 0
  maxC := 0
 for _, num := range deer {
   fmt.Printf("Now considering sum of %d against %d, %d, %d\n", num, maxA, maxB, maxC)
    updateNumbers(num, &maxA, &maxB, &maxC)
 }
 total := maxA + maxB + maxC
 fmt.Printf("And we are done and the max I found was for the top three was: %d, the sum of %d %d %d\n", total, maxA, maxB, maxC)
}

func updateNumbers(num int, a *int, b *int, c *int) {
  tmp := 0
  if num > *a {
    tmp = *a
    *a = num
  } else if num > *b {
    tmp = *b
    *b = num
  } else if num > *c {
    tmp = *c
    *c = num
  }
  if (tmp > 0) {
    updateNumbers(tmp, a, b, c)
  }
}
