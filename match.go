package main

import (
  "bufio"
  "os"
  "strings"
  "fmt"
  "io"
)

// Match implements simple case insnsetive mathc for two hardcoded messages
type Match struct {
}

// Validate thta command has correct arguments
func (match Match) Validate(args []string) bool {
  if len(args) == 3 && args[1] == "-i" && args[2] == "error\\|msg" {
    return true
	}
  return false
}

// Execute command
func (match Match) Execute(args []string) (int, string, error) {
  m := 1
  reader := bufio.NewReader(os.Stdin)
  for {
    text, err := reader.ReadString('\n')
    if err != nil {
      if err == io.EOF {
        return m, "", nil
      }
      return -1, "Error:", err
    }
    lower := strings.ToLower(text)
    if strings.Contains(lower, "error") || strings.Contains(lower, "msg") {
      fmt.Print(text)
      m = 0
    }
  }
  return m, "", nil
}

// Help string
func (match Match) Help() string {
  return "grep -i \"error\\|msg\""
}
