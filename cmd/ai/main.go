package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sashabaranov/go-openai"
)

func main() {
	// 加载环境变量
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY environment variable is not set")
	}
	baseUrl := os.Getenv("OPENAI_API_BASE_URL")
	if baseUrl == "" {
		log.Fatal("OPENAI_API_BASE_URL environment variable is not set")
	}
	model := os.Getenv("OPENAI_API_MODEL")
	if model == "" {
		log.Fatal("OPENAI_API_MODEL environment variable is not set")
	}

	// 解析flags
	flag.Parse()

	// 解析命令行参数
	args := flag.Args()
	prompt := strings.Join(args, " ")
	if prompt == "" {
		log.Fatal("prompt variable is not set")
	}

	// 创建一个新的OpenAI客户端
	// client := openai.NewClientWithBaseURL(apiKey, "https://api.siliconflow.cn/v1")
	config := openai.DefaultConfig(apiKey)
	config.BaseURL = baseUrl
	client := openai.NewClientWithConfig(config)

	// 设置请求参数
	req := openai.ChatCompletionRequest{
		Model: model,
		// Messages: []openai.ChatMessage{
		Messages: []openai.ChatCompletionMessage{
			{Role: openai.ChatMessageRoleUser, Content: prompt},
		},
		Stream: true,
	}

	// 发起流式请求
	stream, err := client.CreateChatCompletionStream(context.Background(), req)
	if err != nil {
		log.Fatalf("Error creating chat completion stream: %v\n", err)
	}
	defer stream.Close()

	// 读取流响应
	for {
		response, err := stream.Recv()
		if err != nil {
			break // 结束流
		}
		if response.Choices != nil && len(response.Choices) > 0 {
			delta := response.Choices[0].Delta
			// if delta.Content != nil {
			if delta.Content != "" {
				fmt.Print(delta.Content)
			}
		}
	}
	fmt.Println("")
}
