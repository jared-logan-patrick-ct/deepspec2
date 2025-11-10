package deepspec

import (
	"context"
	"errors"
)

// ToolFunc is a function that executes a tool with given arguments
type ToolFunc func(ctx context.Context, args map[string]interface{}) (string, error)

// ToolRegistry manages tool functions
type ToolRegistry struct {
	tools map[string]ToolFunc
}

// NewToolRegistry creates a new tool registry
func NewToolRegistry() *ToolRegistry {
	return &ToolRegistry{
		tools: make(map[string]ToolFunc),
	}
}

// Register adds a tool to the registry
func (r *ToolRegistry) Register(name string, fn ToolFunc) {
	r.tools[name] = fn
}

// Get retrieves a tool by name
func (r *ToolRegistry) Get(name string) (ToolFunc, error) {
	fn, ok := r.tools[name]
	if !ok {
		return nil, errors.New("tool not found: " + name)
	}
	return fn, nil
}

// Execute runs a tool by name with the given arguments
func (r *ToolRegistry) Execute(ctx context.Context, name string, args map[string]interface{}) (string, error) {
	fn, err := r.Get(name)
	if err != nil {
		return "", err
	}
	return fn(ctx, args)
}

// Has checks if a tool exists in the registry
func (r *ToolRegistry) Has(name string) bool {
	_, ok := r.tools[name]
	return ok
}
