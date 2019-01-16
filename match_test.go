package main

import "testing"

func TestMatchValidate(t *testing.T) {
  match := Match{}
  if match.Validate([]string{"test", "test"}) {
    t.Error("Incorrect validation")
  }
  if !match.Validate([]string{"grep", "-i", "error\\|msg"}) {
    t.Error("Incorrect possitive validation")
  }
}

func TestMatchExecute(t *testing.T) {
  match := Match{}
  code, message, error := match.Execute([]string{"grep", "-i", "error\\|msg"})
  if code != 1 {
    t.Error("Incorrect exit code")
  }
  if message != "" {
    t.Error("Incorrect result")
  }
  if error != nil {
    t.Error("Error during execution")
  }
}
