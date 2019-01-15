package main

type Operations interface {
    Validate([]string) bool
    Execute([]string) (int, string, error)
    Help() string
}
