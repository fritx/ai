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
	client, err := ai.NewFromEnv()
	if err != nil {
		log.Fatalf("Failed to setup AI client: %v\n", err)
	}
	// 解析flags
	flag.Parse()

	// 解析命令行参数
	args := flag.Args()
	prompt := strings.Join(args, " ")
	if prompt == "" {
		log.Fatalf("prompt variable is not set\n")
	}

	// 发起流式请求
	stream, err := client.ChatStreamOnce(context.Background(), prompt)
	if err != nil {
		log.Fatalf("Error creating chat completion stream: %v\n", err)
	}
	defer stream.Close()

	// 读取流响应
	err = ai.StreamLoop(stream, func(s string) {
		fmt.Print(s)
	})
	fmt.Print("\n")
	if err != nil {
		log.Fatalf("** Error: %v\n", err)
	}
}
