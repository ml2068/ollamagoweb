package main

import (
	"context"
	"fmt"
	"log"
	"encoding/json"
	"github.com/ollama/ollama/api"
)

func main() {
	client, err := api.ClientFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	messages := []api.Message{
		api.Message{
			Role:    "system",
			Content: "Provide very brief, concise responses",
		},
		api.Message{
			Role:    "user",
			Content: "Name some unusual animals",
		},
		api.Message{
			Role:    "assistant",
			Content: "Monotreme, platypus, echidna",
		},
		api.Message{
			Role:    "user",
			Content: "which of these is the most dangerous?",
		},
	}

	ctx := context.Background()

	req := &api.ChatRequest{
		Model:    "llama3.1:8b",
		Messages: messages,
	}
	messagesJSON, _ := json.MarshalIndent(req.Messages, " ", "   ")
	fmt.Println(string(messagesJSON))

	respFunc := func(resp api.ChatResponse) error {
		fmt.Print(resp.Message.Content)
		if resp.Done{
			fmt.Println("\nMetrics:")
			fmt.Println("TotalDuration:", resp.Metrics.TotalDuration)
			fmt.Println("LoadDuration:", resp.Metrics.LoadDuration)
			fmt.Println("PromptEvalCount:", resp.Metrics.PromptEvalCount)
			fmt.Println("PromptEvalDuration:", resp.Metrics.PromptEvalDuration)
			fmt.Println("EvalCount:", resp.Metrics.EvalCount)
			fmt.Println("EvalDuration:", resp.Metrics.EvalDuration)
			fmt.Printf("Speed:%.2f tokens/s\n", float64(resp.Metrics.EvalCount)/resp.Metrics.EvalDuration.Seconds())
		}
	 
		return nil
	}

	err = client.Chat(ctx, req, respFunc)
	if err != nil {
		log.Fatal(err)
	}
}
