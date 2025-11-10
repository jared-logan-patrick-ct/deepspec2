package deepspec

import (
	"context"
	"errors"
)

// Internal tool registry - isolated from external tools for security
var internalTools = NewToolRegistry()

func init() {
	// Register all internal tools
	internalTools.Register("echo", EchoTool)
}

// GetInternalTool retrieves an internal tool by name
func GetInternalTool(name string) (ToolFunc, error) {
	return internalTools.Get(name)
}

// EchoTool echoes back the exact text provided
func EchoTool(ctx context.Context, args map[string]interface{}) (string, error) {
	text, ok := args["text"].(string)
	if !ok {
		return "", errors.New("text parameter is required and must be a string")
	}
	return text, nil
}
