package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	"gitee.com/fritx/ai"
)

func main() {
	if err := ai.Setup(); err != nil {
		log.Fatalf("Failed to setup AI client: %v\n", err)
	}
	// 解析flags
	flag.Parse()

	// 解析命令行参数
	args := flag.Args()
	prompt := strings.Join(args, " ")
	if prompt == "" {
		log.Fatal("prompt variable is not set")
	}

	// 发起流式请求
	stream, err := ai.ChatStream(context.Background(), prompt)
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
