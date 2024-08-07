package main

import (
    "github.com/ollama/ollama/api"
    "encoding/json"
    "fmt"
    "log"
    "context"
)

type Data struct {
    ModelInfo map[string]interface{} `json:"model_info"`
    Details   map[string]interface{} `json:"details"`
    //other data can diret show up 
}

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
    
    //show license
    fmt.Println("\nlicense: \n ", models_show.License)
	
    //show model info
    var data Data
    err = json.Unmarshal(b, &data)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("\n Model Info: \n")
    for key, value := range data.ModelInfo {
        fmt.Printf("%s: %v\n", key, value)
    }
    //show model detail
    fmt.Println("\nModel Detail:\n")
    for key, value := range data.Details {
        fmt.Printf("%s: %v\n", key, value)
    }
}
