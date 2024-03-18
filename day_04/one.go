package main

import (
  "fmt"
  "bufio"
  "os"
  "regexp"
  "strconv"
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
  re := regexp.MustCompile(`(\d+)`)
  for r.Scan() {
	line := r.Text()
    if len(line) > 0 {
	matches := re.FindAllString(line, -1)
	var nums []int;
	for _, match := range matches {
		num, err := strconv.Atoi(match)
		if err != nil {
			fmt.Println("Error converting string to integer:", err)
			return
		}
		nums = append(nums, num)
	}
	if nums[0] >= nums[2] && nums[1] <= nums[3] || nums[2] >= nums[0] && nums[3] <= nums[1] {
		sum++
	}
    	if err != nil {
      	break
    	}
  	}
}
fmt.Printf("The total sum is: %d\n", sum);
}


