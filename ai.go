package ai

import (
	"context"
	"errors"
	"os"

	"github.com/sashabaranov/go-openai"
)

type Client struct {
	client *openai.Client
	model  string
}

func New(apiKey, baseUrl, model string) *Client {
	// 创建一个新的OpenAI客户端
	config := openai.DefaultConfig(apiKey)
	config.BaseURL = baseUrl
	client := openai.NewClientWithConfig(config)
	c := &Client{client, model}
	return c
}

func NewFromEnv() (*Client, error) {
	// 加载环境变量
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return nil, errors.New("OPENAI_API_KEY environment variable is not set")
	}
	baseUrl := os.Getenv("OPENAI_API_BASE_URL")
	if baseUrl == "" {
		return nil, errors.New("OPENAI_API_BASE_URL environment variable is not set")
	}
	model := os.Getenv("OPENAI_API_MODEL")
	if model == "" {
		return nil, errors.New("OPENAI_API_MODEL environment variable is not set")
	}
	return New(apiKey, baseUrl, model), nil
}

func (c *Client) Chat(ctx context.Context, messages []openai.ChatCompletionMessage) (openai.ChatCompletionResponse, error) {
	return c.client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:    c.model,
		Messages: messages,
	})
}

func (c *Client) ChatStream(ctx context.Context, messages []openai.ChatCompletionMessage) (*openai.ChatCompletionStream, error) {
	return c.client.CreateChatCompletionStream(ctx, openai.ChatCompletionRequest{
		Model:    c.model,
		Messages: messages,
	})
}

func (c *Client) ChatOnce(ctx context.Context, prompt string) (openai.ChatCompletionResponse, error) {
	return c.Chat(ctx, []openai.ChatCompletionMessage{
		{Role: openai.ChatMessageRoleUser, Content: prompt},
	})
}

func (c *Client) ChatStreamOnce(ctx context.Context, prompt string) (*openai.ChatCompletionStream, error) {
	return c.ChatStream(ctx, []openai.ChatCompletionMessage{
		{Role: openai.ChatMessageRoleUser, Content: prompt},
	})
}
