package deepspec

import (
	"context"
	"strings"
	"sync"

	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/client/transport"
	"github.com/mark3labs/mcp-go/mcp"
)

type mcpClient struct {
	mutex  *sync.Mutex
	client *client.Client
}

// NewMCPClient creates a new MCP client instance
func NewMCPClient() *mcpClient {
	return &mcpClient{
		mutex: &sync.Mutex{},
	}
}

// Start initializes and starts the MCP client
func (f *mcpClient) Start(ctx context.Context) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	// Initialize the MCP client
	tp, err := transport.NewStreamableHTTP(MCPServerAddress)
	if err != nil {
		return err
	}
	f.client = client.NewClient(tp)

	// Start the client
	if err := f.client.Start(ctx); err != nil {
		return err
	}

	initRequest := mcp.InitializeRequest{
		Params: mcp.InitializeParams{
			ClientInfo: mcp.Implementation{
				Name:    "deepspec-tui",
				Version: "0.1.0",
			},
			ProtocolVersion: "2024-11-05",
		},
	}

	_, err = f.client.Initialize(ctx, initRequest)
	return err
}

// reconnect attempts to reset and restart the MCP client connection
func (f *mcpClient) reconnect(ctx context.Context) error {
	LogToFile("Session terminated, attempting reconnect")

	// Reset client
	f.mutex.Lock()
	if f.client != nil {
		f.client.Close()
		f.client = nil
	}
	f.mutex.Unlock()

	// Attempt to restart
	if err := f.Start(ctx); err != nil {
		LogToFile("Reconnect failed: %v", err)
		return err
	}

	return nil
}

// HealthCheck pings the MCP server to check its health
func (f *mcpClient) HealthCheck(ctx context.Context) (bool, error) {
	f.mutex.Lock()
	c := f.client
	f.mutex.Unlock()

	if c == nil {
		return false, nil
	}

	request := mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Name: "healthz",
		},
	}
	result, err := c.CallTool(ctx, request)
	if err != nil {
		// Check if error indicates need to re-initialize
		if !f.client.IsInitialized() || strings.Contains(err.Error(), "need to re-initialize") || strings.Contains(err.Error(), "Invalid session ID") {
			if reconnectErr := f.reconnect(ctx); reconnectErr != nil {
				return false, reconnectErr
			}

			// Retry health check after reconnect
			f.mutex.Lock()
			c = f.client
			f.mutex.Unlock()

			if c == nil {
				return false, nil
			}

			result, err = c.CallTool(ctx, request)
			if err != nil {
				return false, err
			}
		} else {
			return false, err
		}
	}

	return !result.IsError, nil
}
