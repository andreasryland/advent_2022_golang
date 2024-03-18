package main

import (
  "fmt"
  "bufio"
  "os"
  "math"
  "regexp"
  "strconv"
)

type Stack struct {
	items []rune
}

// Push
func (s *Stack) Push(i rune) {
	s.items = append(s.items, i)
}

// Pop
func (s *Stack) Pop() rune {
	last := len(s.items) - 1
	result := s.items[last]
	s.items = s.items[:last]
	return result
}

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
  result := ""
  setupDone := false
  re := regexp.MustCompile(`(\d+)`)
  var setupLines [][]rune
  var stacks []Stack
  // Section 2
  for r.Scan() {
	line := r.Text()
    if len(line) > 0 {
	    if !setupDone {
	    	fmt.Printf("Line length is %d\n", len(line))
		runeSlice := []rune(line)
		setupLines = append(setupLines, runeSlice)
    } else { 
	match := re.FindAllStringSubmatch(line, -1)
	fmt.Println(line)
	fmt.Println(match)
	numItems, _ := strconv.Atoi(match[0][0])
	from, _ := strconv.Atoi(match[1][0])
	to, _ := strconv.Atoi(match[2][0])  
	fmt.Printf("Move %d items from %d to %d\n", numItems, from, to)
	for i := 0; i < numItems; i++ {
		itm := stacks[from - 1].Pop()
		stacks[to - 1].Push(itm)
	}
    }
  	} else if !setupDone {
		setupDone = true
		fmt.Println("Setup is now done")
		fmt.Println(setupLines)
		stacks = make([]Stack, int(math.Ceil(float64(len(setupLines[1])) / 4)))
		for i := len(setupLines) - 2; i >= 0; i-- {
			for e, n := 1, 0; e < len(setupLines[i]); e, n = e + 4, n + 1 {
				if setupLines[i][e] != ' ' {
					stacks[n].Push(setupLines[i][e])
				fmt.Printf("should add this %c to stack %d\n", setupLines[i][e], n)
			}
			}
		}
	}
}
fmt.Println(stacks)
for i := 0; i < len(stacks); i++ {
result += fmt.Sprintf("%c", stacks[i].Pop())
}
fmt.Printf("The result is: %s\n", result);
}


