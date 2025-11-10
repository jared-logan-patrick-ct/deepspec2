package deepspec

import (
	"log"

	"github.com/mark3labs/mcp-go/server"
)

// Server represents our MCP server instance
type Server struct {
	mcpServer *server.MCPServer
}

// NewServer creates and initializes a new MCP server
func NewServer() *Server {
	// Create a new MCP server with tool capabilities
	mcpServer := server.NewMCPServer(
		"deepspec-server",
		"1.0.0",
		server.WithToolCapabilities(true),
	)

	srv := &Server{mcpServer: mcpServer}

	// Register all external tools
	RegisterExternalTools(mcpServer)

	return srv
}

// Start starts the MCP server on HTTP port 8080
func (s *Server) Start() error {
	log.Println("Starting MCP server on :8080...")
	httpServer := server.NewStreamableHTTPServer(s.mcpServer)
	return httpServer.Start(":8080")
}

// GetMCPServer returns the underlying MCP server for advanced usage
func (s *Server) GetMCPServer() *server.MCPServer {
	return s.mcpServer
}
