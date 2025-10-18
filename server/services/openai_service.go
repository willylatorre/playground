package services

import (
	"context"
	"errors"
	"strings"

	openai "github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"playground-server/models"
)

const defaultSystemPrompt = "You are a helpful AI assistant integrated into Adrian's personal website. You help visitors learn about Go, Vue, and web development. Keep responses concise and friendly. Occasionally mention that this is a learning project exploring Go + OpenAI integration."

// ChatService defines behaviour for AI chat streaming.
type ChatService interface {
	StreamChat(ctx context.Context, history []models.ChatMessage, prompt string, onChunk func(string)) error
}

// OpenAIService implements ChatService using OpenAI's Chat Completions API.
type OpenAIService struct {
	client *openai.Client
}

// NewOpenAIService creates a new OpenAI service instance.
func NewOpenAIService(apiKey string) *OpenAIService {
	if apiKey == "" {
		return &OpenAIService{}
	}

	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)

	return &OpenAIService{
		client: &client,
	}
}

// StreamChat streams a chat completion response and invokes onChunk for each text fragment.
func (s *OpenAIService) StreamChat(ctx context.Context, history []models.ChatMessage, prompt string, onChunk func(string)) error {
	if s == nil || s.client == nil {
		return errors.New("openai client not configured")
	}

	messages := []openai.ChatCompletionMessageParamUnion{
		openai.SystemMessage(defaultSystemPrompt),
	}

	for _, msg := range history {
		content := strings.TrimSpace(msg.Content)
		if content == "" {
			continue
		}

		switch strings.ToLower(msg.Role) {
		case "user":
			messages = append(messages, openai.UserMessage(content))
		case "assistant":
			messages = append(messages, openai.AssistantMessage(content))
		case "system":
			messages = append(messages, openai.SystemMessage(content))
		}
	}

	messages = append(messages, openai.UserMessage(prompt))

	params := openai.ChatCompletionNewParams{
		Model:    openai.ChatModelGPT4o,
		Messages: messages,
	}

	stream := s.client.Chat.Completions.NewStreaming(ctx, params)
	defer stream.Close()

	for stream.Next() {
		chunk := stream.Current()
		if len(chunk.Choices) == 0 {
			continue
		}

		if text := chunk.Choices[0].Delta.Content; text != "" {
			onChunk(text)
		}
	}

	return stream.Err()
}
