package main

import (
  "fmt"
  "bufio"
  "os"
)

func main() {
  //Section 1
file, err := os.Open("input.txt")
defer file.Close()
if err != nil {
  panic("fail")
}
r := bufio.NewReader(file)

sum := 0
// Section 2
for {
  line, _, err := r.ReadLine()

  if len(line) > 0 {
    score := 0
    theirs := line[0]
    mine := line[2]

    if mine == 'Y' {
      if theirs == 'A' {
        mine = 'X'
      } else if theirs == 'B' {
        mine = 'Y'
      } else {
        mine = 'Z'
      }
    } else if mine == 'X' {
      if theirs == 'A' {
        mine = 'Z'
      } else if theirs == 'B' {
        mine = 'X'
      } else {
        mine = 'Y'
      }
    } else {
      if theirs == 'A' {
        mine = 'Y'
      } else if theirs == 'B' {
        mine = 'Z'
      } else {
        mine = 'X'
      }
    }
    if mine == 'Y' {
      score += 2
    } else if mine == 'X' {
      score += 1
    } else if mine == 'Z' {
      score += 3
    }

    if mine == 'Y' && theirs == 'B' ||
      mine == 'X' && theirs == 'A' ||
      mine == 'Z' && theirs == 'C' {
      score += 3
    } else if mine == 'Y' && theirs == 'A' ||
      mine == 'X' && theirs == 'C' ||
      mine == 'Z' && theirs == 'B' {
        score += 6
      }
    if err != nil {
      panic("bad number")
    }
    sum += score
  }
  if err != nil {
    break
  }
 }
 fmt.Printf("And we are done and the total score is: %d\n", sum)
}

