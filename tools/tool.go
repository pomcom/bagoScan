package tools

// Interface for any tool, that can be executed
type Tool interface {
	// Runs the tool with the provided flags and returns jsonOuput or an error
	Execute(flags string) (jsonOutput string, err error)
}
