package main

import (
    
    "context"
    "fmt"
    "os"

    "github.com/joho/godotenv"
    "github.com/tmc/langchaingo/llms"
    "github.com/tmc/langchaingo/llms/openai"
)


// initialise to load environment variable from .env file
func init() {
    err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file")
        return
    }
}

func main() {
    llm, err := openai.New(
        openai.WithModel(os.Getenv("llm")),
        openai.WithBaseURL(os.Getenv("baseUrl")),
        openai.WithToken(os.Getenv("apiKey")),
    )
    if err != nil {
        fmt.Println("Cannot create local LLM:", err.Error())
        return
    }

    msg := llms.MessageContent{
        Role:  llms.ChatMessageTypeHuman,
        Parts: []llms.ContentPart{
            llms.TextContent{Text: "how many planets are there?"},
        },
    }

    opts := []llms.CallOption{
        llms.WithMaxTokens(1024),
        llms.WithTemperature(0.5),
    }

    resp, err := llm.GenerateContent(context.Background(), []llms.MessageContent{msg}, opts...)
    if err != nil {
        fmt.Println(err)
        return
    }

    if resp != nil {
        choices := resp.Choices
        for _, c := range choices {
            fmt.Println(c.Content)
            fmt.Println(c.GenerationInfo)
        }
    }
}
