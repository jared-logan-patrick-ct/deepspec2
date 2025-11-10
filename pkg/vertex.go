package deepspec

import (
	"context"
	"os"

	genai "google.golang.org/genai"
)

// Environment variables
var (
	gcpProjectID = os.Getenv("GOOGLE_CLOUD_PROJECT")
	gcpLocation  = os.Getenv("GCP_LOCATION")
	gcpModelName = os.Getenv("GCP_MODEL_NAME")
)

const DefaultModelName = "gemini-2.5-flash-lite"

const SystemPrompt = `You are an AI assistant with access to specific function tools. You MUST ONLY provide responses that can be handled by the available user-provided functions.

CRITICAL RULES:
1. Before responding, verify that an appropriate function exists to handle your response
2. If no suitable function is available for the user's request, politely explain that you cannot perform that action
3. Do not make assumptions about available capabilities beyond the provided functions
4. Always structure your responses to align with the function schemas provided
5. If asked to perform an action without a corresponding function, decline and suggest what you CAN do instead

Remember: You can only take actions through the functions provided to you. Do not promise or attempt actions for which no function exists.`

// VertexClient wraps the GCP Vertex AI client
type VertexClient struct {
	client    *genai.Client
	ctx       context.Context
	projectID string
	location  string
	modelName string
}

// NewVertexClient creates a new Vertex AI client using environment variables
// Requires:
//   - GOOGLE_CLOUD_PROJECT: Google Cloud project ID
//   - GCP_LOCATION: Google Cloud location (e.g., us-central1)
//   - GCP_MODEL_NAME: Model name (optional, defaults to gemini-1.5-flash-latest)
func NewVertexClient(ctx context.Context) (*VertexClient, error) {
	modelName := gcpModelName
	if modelName == "" {
		modelName = DefaultModelName
	}

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		Project:  gcpProjectID,
		Location: gcpLocation,
		Backend:  genai.BackendVertexAI,
	})
	if err != nil {
		return nil, err
	}
	return &VertexClient{
		client:    client,
		ctx:       ctx,
		projectID: gcpProjectID,
		location:  gcpLocation,
		modelName: modelName,
	}, nil
}

// StartChat starts a new chat session with tool support
func (v *VertexClient) StartChat() (*genai.Chat, error) {
	// Define the echo tool
	echoTool := &genai.Tool{
		FunctionDeclarations: []*genai.FunctionDeclaration{
			{
				Name:        "echo",
				Description: "Echo back the exact text provided. Used for testing.",
				Parameters: &genai.Schema{
					Type: genai.TypeObject,
					Properties: map[string]*genai.Schema{
						"text": {
							Type:        genai.TypeString,
							Description: "The text to echo back",
						},
					},
					Required: []string{"text"},
				},
			},
		},
	}

	// Create config with tools
	config := &genai.GenerateContentConfig{
		Tools: []*genai.Tool{echoTool},
	}

	chat, err := v.client.Chats.Create(v.ctx, v.modelName, config, nil)
	if err != nil {
		return nil, err
	}
	return chat, nil
}

// SendMessage sends a message to the chat session and returns the response
// The system prompt is prepended to every message to ensure consistent behavior
// This method handles tool calls recursively until a final text response is received
func (v *VertexClient) SendMessage(chat *genai.Chat, message string) (string, error) {
	// Prepend system instructions to maintain context throughout the conversation
	fullMessage := SystemPrompt + "\n\n" + message

	result, err := chat.SendMessage(v.ctx, genai.Part{Text: fullMessage})
	if err != nil {
		return "", err
	}

	// Handle tool calls recursively
	return v.handleResponse(chat, result)
}

// handleResponse processes the LLM response, executing tool calls if needed
func (v *VertexClient) handleResponse(chat *genai.Chat, response *genai.GenerateContentResponse) (string, error) {
	// Check for function calls
	functionCalls := response.FunctionCalls()
	if len(functionCalls) == 0 {
		// No tool calls, return the text response
		return response.Text(), nil
	}

	// Execute each function call and collect responses
	var functionResponses []genai.Part
	for _, fc := range functionCalls {
		result := v.executeFunction(fc)
		functionResponses = append(functionResponses, *genai.NewPartFromFunctionResponse(fc.Name, result))
	}

	// Send function responses back to the model
	result, err := chat.SendMessage(v.ctx, functionResponses...)
	if err != nil {
		return "", err
	}

	// Recurse to handle potential additional tool calls
	return v.handleResponse(chat, result)
}

// executeFunction executes a local function and returns the result
func (v *VertexClient) executeFunction(fc *genai.FunctionCall) map[string]any {
	// Try to get the tool from internal registry
	toolFn, err := GetInternalTool(fc.Name)
	if err != nil {
		return map[string]any{"error": "unknown function: " + fc.Name}
	}

	// Execute the tool
	result, err := toolFn(v.ctx, fc.Args)
	if err != nil {
		return map[string]any{"error": err.Error()}
	}

	return map[string]any{"output": result}
}
