package handlers

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"playground-server/models"
)

// ChatHandler handles AI chat requests
type ChatHandler struct {
	apiKey string
}

// NewChatHandler creates a new chat handler
func NewChatHandler(apiKey string) *ChatHandler {
	return &ChatHandler{apiKey: apiKey}
}

// OpenAIMessage represents a message in the OpenAI format
type OpenAIMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// OpenAIRequest represents the request to OpenAI API
type OpenAIRequest struct {
	Model    string          `json:"model"`
	Messages []OpenAIMessage `json:"messages"`
	Stream   bool            `json:"stream"`
}

// SendMessage handles streaming chat messages from OpenAI
func (h *ChatHandler) SendMessage(c *gin.Context) {
	// Check if API key is configured
	if h.apiKey == "" {
		log.Println("ERROR: OpenAI API key not configured")
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"error": "AI service not configured. Please set OPENAI_API_KEY environment variable.",
		})
		return
	}

	var req models.ChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("ERROR: Invalid request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request format",
		})
		return
	}

	// Easter egg detection
	easterEgg := checkForEasterEgg(req.Message)
	if easterEgg != "" {
		c.Header("Content-Type", "text/event-stream")
		c.Header("Cache-Control", "no-cache")
		c.Header("Connection", "keep-alive")
		
		// Stream the easter egg response
		streamText(c.Writer, easterEgg)
		return
	}

	// Prepare OpenAI request
	openAIReq := OpenAIRequest{
		Model: "gpt-3.5-turbo",
		Messages: []OpenAIMessage{
			{
				Role:    "system",
				Content: "You are a helpful AI assistant integrated into Adrian's personal website. You help visitors learn about Go, Vue, and web development. Keep responses concise and friendly. Occasionally mention that this is a learning project exploring Go + OpenAI integration.",
			},
			{
				Role:    "user",
				Content: req.Message,
			},
		},
		Stream: true,
	}

	// Make request to OpenAI
	jsonData, err := json.Marshal(openAIReq)
	if err != nil {
		log.Printf("ERROR: Failed to marshal request: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to process request",
		})
		return
	}

	httpReq, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("ERROR: Failed to create request: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create request",
		})
		return
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+h.apiKey)

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		log.Printf("ERROR: Failed to call OpenAI: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to connect to AI service",
		})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("ERROR: OpenAI API error: %s", string(body))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "AI service returned an error",
		})
		return
	}

	// Set headers for SSE
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")

	// Stream the response
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "data: ") {
			data := strings.TrimPrefix(line, "data: ")
			if data == "[DONE]" {
				break
			}

			// Parse the chunk
			var chunk struct {
				Choices []struct {
					Delta struct {
						Content string `json:"content"`
					} `json:"delta"`
				} `json:"choices"`
			}

			if err := json.Unmarshal([]byte(data), &chunk); err != nil {
				continue
			}

			if len(chunk.Choices) > 0 && chunk.Choices[0].Delta.Content != "" {
				fmt.Fprintf(c.Writer, "data: %s\n\n", chunk.Choices[0].Delta.Content)
				c.Writer.Flush()
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("ERROR: Stream error: %v", err)
	}
}

// checkForEasterEgg checks if the message contains an easter egg trigger
func checkForEasterEgg(message string) string {
	lower := strings.ToLower(message)
	
	if strings.Contains(lower, "coffee") && (strings.Contains(lower, "how many") || strings.Contains(lower, "count")) {
		return "🎉 Easter Egg Found! The coffee counter on the dashboard shows how many cups Adrian has had. It's a fun way to track caffeine consumption while learning Go + SQLite!"
	}
	
	if strings.Contains(lower, "secret") || strings.Contains(lower, "easter egg") {
		return "🥚 You found the easter egg system! Try asking about coffee, or explore the dashboard. There might be more hidden surprises..."
	}
	
	if strings.Contains(lower, "adrian") && strings.Contains(lower, "who") {
		return "👨‍💻 Adrian is a developer exploring the intersection of Go backends and modern Vue frontends. This entire site is a learning playground for experimenting with clean architecture, type safety, and AI integration!"
	}
	
	return ""
}

// streamText is a helper to stream text as SSE
func streamText(w gin.ResponseWriter, text string) {
	words := strings.Fields(text)
	for _, word := range words {
		fmt.Fprintf(w, "data: %s \n\n", word)
		w.Flush()
	}
}
