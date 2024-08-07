package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"github.com/ollama/ollama/api"
)



func main() {
	//set new <scheme>://<host>:<port>
	base, err := url.Parse("http://localhost:11434")
	if err != nil {
		log.Fatal(err)
	}
	//creat new client
    client := api.NewClient(base, &http.Client{})


	req := &api.GenerateRequest{
		Model:  "llama3.1:8b",
		Prompt: "how many planets are there?",
	}

	ctx := context.Background()
	respFunc := func(resp api.GenerateResponse) error {
		// Only print the response here; GenerateResponse has a number of other
		// interesting fields you want to examine.
		fmt.Print(resp.Response)
		return nil
	}

	err = client.Generate(ctx, req, respFunc)
	if err != nil {
		log.Fatal(err)
	}
}