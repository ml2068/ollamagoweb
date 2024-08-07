package main

import (
    "github.com/ollama/ollama/api"
    "encoding/json"
    "fmt"
    "log"
    "context"
)

func main() {
	//load client
    client, err := api.ClientFromEnvironment()
    if err != nil {
        log.Fatal(err)
    }

    ctx := context.Background()

    sq := &api.ShowRequest{
		Model: "llama3.1:8b",
	}

    models_show,err:=client.Show(ctx, sq )
    if err != nil {
        log.Fatal(err)
    }

    b, err := json.MarshalIndent(models_show, "   ", "   ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nmodel_show: \n")
	fmt.Println(string(b))

}