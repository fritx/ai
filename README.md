# fritx/ai

### 命令行用法：[cmd/ai](./cmd/ai/README.md)

```sh
go install gitee.com/fritx/ai/cmd/ai@latest

# 添加对应的环境变量
# - 如 硅基流动: https://docs.siliconflow.cn/quickstart
export OPENAI_API_KEY="sk-xxxxxxxxxxxxxxxxxxxxxxxxxxxx"
export OPENAI_API_BASE_URL="https://api.siliconflow.cn/v1"
export OPENAI_API_MODEL="Qwen/Qwen2.5-7B-Instruct"  # 免费

ai "你是什么模型"
>> 我是Qwen模型，由阿里巴巴云开发。作为一个预训练语言模型，我能够生成与给定词语相关的文本，帮助回答问题、撰写文章等多种自然语言处理任务。如果您有任何问题或需要帮助，请随时告诉我！
```

### API用法：

**客户端创建：ai.New / ai.NewFromEnv**

```go
import "gitee.com/fritx/ai"

client, err := ai.New(apiKey, baseUrl, model)

// or 从环境变量读取配置
client, err := ai.NewFromEnv()
```

**长上下文对话：ai.Chat / ai.ChatStream**

```go
import "github.com/sashabaranov/go-openai"

response, err := ai.Chat(ctx, []openai.ChatCompletionMessage{
	{Role: openai.ChatMessageRoleSystem, Content: "..."},
	{Role: openai.ChatMessageRoleUser, Content: "..."},
	{Role: openai.ChatMessageRoleAssistant, Content: "..."},
	{Role: openai.ChatMessageRoleUser, Content: "..."},
})

// or 发起流式请求
stream, err := ai.ChatStream(ctx, []openai.ChatCompletionMessage{
	{Role: openai.ChatMessageRoleAssistant, Content: "..."},
	{Role: openai.ChatMessageRoleUser, Content: "..."},
})
defer stream.Close()
````

**一次性对话：ai.ChatOnce / ai.ChatStreamOnce**

```go
prompt := "中国大模型行业2025年将会迎来哪些机遇和挑战"

response, err := ai.ChatOnce(ctx, prompt)
if err != nil {
	// ...
}
fmt.Println(response.Choices[0].Message)

// or 发起流式请求
stream, err := ai.ChatStreamOnce(ctx, prompt)
if err != nil {
	// ...
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
```
