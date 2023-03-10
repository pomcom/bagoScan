package tools

// Every tool must implement this interface
type Tool interface {
	// The implementation gets executed in the `testrunner.go`
	Execute(target string) (output string, err error)
}
