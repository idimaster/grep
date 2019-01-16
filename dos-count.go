package main

import (
  "io/ioutil"
  "strconv"
)

// CountDOS implement simple counter for DOS line ending
type CountDOS struct {
}

// Validate thta command has correct arguments
func (count CountDOS) Validate(args []string) bool {
  if len(args) == 4 && args[1] == "-UPc" && args[2] == "\\r$" {
    return true
	}
  return false
}

// Execute command
func (count CountDOS) Execute(args []string) (int, string, error) {
  content, err := ioutil.ReadFile(args[3])
	if err != nil {
    return -1, "Error:", err
	}
	c := 0
	for i := 0; i < len(content) - 1; i++ {
		if content[i] == 0xD && content[i + 1] == 0xA {
				//count dos line ending
				c++
		}
	}
  return 1, strconv.Itoa(c), nil
}

// Help string
func (count CountDOS) Help() string {
  return "grep -UPc \"\\r$\" file"
}
