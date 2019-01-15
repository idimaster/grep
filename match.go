package main

import (
  "bufio"
  "os"
  "strings"
  "fmt"
  "io"
)

type Match struct {
}

func (match Match) Validate(args []string) bool {
  if len(args) == 3 && args[1] == "-i" && args[2] == "error\\|msg" {
    return true
	}
  return false
}

func (match Match) Execute(args []string) (int, string, error) {
  reader := bufio.NewReader(os.Stdin)
  for {
    text, err := reader.ReadString('\n')
    if err != nil {
      if err == io.EOF {
        return 0, "", nil
      }
      return -1, "Error:", err
    }
    lower := strings.ToLower(text)
    if strings.Contains(lower, "error") || strings.Contains(lower, "msg") {
      fmt.Print(text)
    }
  }
  return 0, "", nil
}

func (match Match) Help() string {
  return "grep -i \"error\\|msg\""
}
