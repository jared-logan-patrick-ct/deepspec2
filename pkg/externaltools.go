package deepspec

import (
	"context"
	"log"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// HealthzTool checks the health status of the server
func HealthzTool(ctx context.Context, args map[string]interface{}) (string, error) {
	log.Println("Healthz check requested")
	return "OK", nil
}

// healthzHandler is the MCP handler for the healthz tool
func healthzHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Convert arguments
	args := make(map[string]interface{})
	if request.Params.Arguments != nil {
		if argsMap, ok := request.Params.Arguments.(map[string]interface{}); ok {
			args = argsMap
		}
	}

	// Execute the tool
	result, err := HealthzTool(ctx, args)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: result,
			},
		},
	}, nil
}

// RegisterExternalTools registers all external tools with the MCP server
func RegisterExternalTools(mcpServer *server.MCPServer) {
	// Register healthz tool
	healthzTool := mcp.NewTool("healthz",
		mcp.WithDescription("Check the health status of the MCP server"),
	)
	mcpServer.AddTool(healthzTool, healthzHandler)
}
