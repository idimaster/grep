package main

// Operations generic CLI oprtation
type Operations interface {
    // Validate aruments, returns true if args match command
    Validate([]string) bool
    // Execute actual command
    Execute([]string) (int, string, error)
    // Show usage help for the command
    Help() string
}
