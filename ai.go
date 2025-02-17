package ai

import (
	"context"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
)

var client *openai.Client
var model string

func Setup() error {
	// 加载环境变量
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("OPENAI_API_KEY environment variable is not set")
	}
	baseUrl := os.Getenv("OPENAI_API_BASE_URL")
	if baseUrl == "" {
		return fmt.Errorf("OPENAI_API_BASE_URL environment variable is not set")
	}
	model = os.Getenv("OPENAI_API_MODEL")
	if model == "" {
		return fmt.Errorf("OPENAI_API_MODEL environment variable is not set")
	}
	// 创建一个新的OpenAI客户端
	// client := openai.NewClientWithBaseURL(apiKey, "https://api.siliconflow.cn/v1")
	config := openai.DefaultConfig(apiKey)
	config.BaseURL = baseUrl
	client = openai.NewClientWithConfig(config)
	return nil
}

func Chat(ctx context.Context, prompt string) (openai.ChatCompletionResponse, error) {
	if client == nil {
		return openai.ChatCompletionResponse{}, fmt.Errorf("client is nil, please call ai.Setup first")
	}
	// 发起流式请求
	return client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: model,
		// Messages: []openai.ChatMessage{
		Messages: []openai.ChatCompletionMessage{
			{Role: openai.ChatMessageRoleUser, Content: prompt},
		},
		Stream: false,
	})
}

func ChatStream(ctx context.Context, prompt string) (*openai.ChatCompletionStream, error) {
	if client == nil {
		return nil, fmt.Errorf("client is nil, please call ai.Setup first")
	}
	// 发起流式请求
	return client.CreateChatCompletionStream(ctx, openai.ChatCompletionRequest{
		Model: model,
		// Messages: []openai.ChatMessage{
		Messages: []openai.ChatCompletionMessage{
			{Role: openai.ChatMessageRoleUser, Content: prompt},
		},
		Stream: true,
	})
}
