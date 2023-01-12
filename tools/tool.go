package tools

// Interface for any tool, that can be executed
type Tool interface {
	// Runs the tool with the provided flags and returns the output or an error
	Execute(flags string) (output string, err error)
	Name() string
}
