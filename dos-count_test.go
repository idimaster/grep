package main

import "testing"

func TestCountValidate(t *testing.T) {
  count := CountDOS{}
  if count.Validate([]string{"test", "test"}) {
    t.Error("Incorrect validation")
  }
  if !count.Validate([]string{"grep", "-UPc", "\\r$", "file"}) {
    t.Error("Incorrect possitive validation")
  }
}

func TestCountExecute(t *testing.T) {
  count := CountDOS{}
  code, message, error := count.Execute([]string{"grep", "-UPc", "\\r$", "grep.go"})
  if code != 1 {
    t.Error("Incorrect exit code")
  }
  if message != "0" {
    t.Error("Incorrect result")
  }
  if error != nil {
    t.Error("Error during execution")
  }
}
