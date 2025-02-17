package ai

import (
	"errors"
	"io"

	"github.com/sashabaranov/go-openai"
)

func StreamLoop(stream *openai.ChatCompletionStream, onRecv func(string)) error {
	if stream == nil {
		return errors.New("stream is nil")
	}
	// 读取流响应
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if response.Choices != nil && len(response.Choices) > 0 {
			delta := response.Choices[0].Delta
			onRecv(delta.Content)
		}
	}
}
