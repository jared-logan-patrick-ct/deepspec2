# DeepSpec

A TUI-based AI assistant with Model Context Protocol (MCP) integration and universal code specification system.

## Features

- 🤖 **AI-Powered Chat**: Gemini-powered conversational interface
- 🔌 **MCP Integration**: Tool calling via Model Context Protocol
- 📊 **Code Specifications**: Universal JSON-based code representation system
- 🎨 **Beautiful TUI**: Terminal UI with bubbletea and lipgloss
- ⚡ **Real-time Health Checks**: Automatic MCP server monitoring

## Quick Start

```bash
# Build and run
go run ./cmd

# Or build binary
go build -o deepspec ./cmd
./deepspec
```

## Project Structure

```
deepspec2/
├── cmd/                     # Main application entry point
├── pkg/                     # Core packages
│   ├── api.go              # MCP server API
│   ├── chatmodel.go        # Chat orchestration
│   ├── vertex.go           # Gemini/Vertex AI client
│   ├── client.go           # MCP client with health checks
│   ├── *component.go       # TUI components
│   └── ...
├── schemas/                # JSON Schema definitions
│   ├── code-spec.schema.json  # Universal code spec format (v1.0.0)
│   └── README.md
├── specs/                  # Code specifications (15 files, 1,197 LOC)
│   ├── *.json              # JSON specs for each pkg/ file
│   └── README.md
└── scripts/                # Utility scripts
    ├── validate-specs.sh   # Validate specifications
    └── update-specs.sh     # Update spec format
```

## Code Specification System

DeepSpec includes a **universal code specification format** that represents program structure in a language-agnostic JSON format.

### Key Features

✅ **Language-Agnostic**: Supports Go, Python, JavaScript, TypeScript, Rust, Java, C, C++
✅ **JSON Schema Validated**: All specs conform to `schemas/code-spec.schema.json`
✅ **Detailed AST**: Statement-level breakdown for complex code
✅ **Code Generation Ready**: Regenerate source from specifications
✅ **Simple & Practical**: KISS principle, minimal complexity

### Quick Example

```json
{
  "$schema": "../schemas/code-spec.schema.json",
  "spec_version": "1.0.0",
  "language": "go",
  "module": "deepspec",
  "file": "config.go",
  "constants": [
    {
      "name": "Version",
      "type": "string",
      "value": "v0.1.0",
      "exported": true
    }
  ]
}
```

### Validation

```bash
# Basic validation
./scripts/validate-specs.sh

# Full schema validation (requires ajv-cli)
npm install -g ajv-cli
ajv validate -s schemas/code-spec.schema.json -d "specs/*.json"
```

## Documentation

- 📖 [Code Specifications](specs/README.md) - Detailed spec documentation
- 📐 [Schema Documentation](schemas/README.md) - JSON Schema format
- 📝 [Specifications Overview](SPECIFICATIONS.md) - System overview
- 🔄 [Schema Migration Guide](SCHEMA_MIGRATION.md) - Migration details

## Environment Setup

Required environment variables for Vertex AI:

```bash
export GOOGLE_CLOUD_PROJECT="your-project-id"
export GCP_LOCATION="us-central1"
export GCP_MODEL_NAME="gemini-1.5-flash"  # optional
```

## MCP Server

The MCP server must be running for health checks:

```bash
# Start MCP server (default: http://localhost:8080/mcp)
go run ./cmd server
```

## Architecture

### Components

- **ChatModel**: Main orchestrator integrating UI and AI
- **VertexClient**: Gemini integration with function calling
- **mcpClient**: MCP protocol client with auto-reconnect
- **ViewportComponent**: Scrollable message display with loader
- **InputComponent**: User input with command mode detection
- **HelpComponent**: Connection status and help text

### Message Flow

```
User Input → ChatModel → VertexClient → Gemini API
                ↓              ↓
          ViewportComponent  Tool Calls → MCP Server
```

### Health Checking

- Automatic 3-second health check intervals
- Auto-reconnect on session termination
- Visual connection status indicator

## Development

### Adding New Specifications

1. Create JSON spec following schema format
2. Validate with `./scripts/validate-specs.sh`
3. Include required fields: `spec_version`, `language`, `module`, `file`

### Extending the Schema

The schema is designed for easy extension via optional fields:
- Intent annotations
- Design pattern tags
- Cross-language type mappings
- Execution context metadata

## Dependencies

- **Bubbletea** - TUI framework
- **Lipgloss** - Terminal styling
- **MCP-Go** - Model Context Protocol (v0.43.0)
- **Google Vertex AI** - Gemini model access

## License

[Your License Here]

## Version

Current: **v0.1.0**
Specification Format: **1.0.0**

---

**Built with ❤️ using Go and the power of AI**
