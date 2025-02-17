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

**客户端初始化：ai.Setup**

```go
import "gitee.com/fritx/ai"

if err := ai.Setup(); err != nil {
	log.Fatalf("Failed to setup AI client: %v\n", err)
}
// ...
```

**普通调用：ai.Chat**

```go
prompt := "中国大模型行业2025年将会迎来哪些机遇和挑战"

response, err := ai.Chat(context.TODO(), prompt)
if err != nil {
	log.Fatalf("Error creating chat completion stream: %v\n", err)
}
fmt.Println(response.Choices[0].Message)
```

**流式读取：ai.ChatStream**

```go
prompt := "中国大模型行业2025年将会迎来哪些机遇和挑战"

stream, err := ai.ChatStream(context.TODO(), prompt)
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
```
