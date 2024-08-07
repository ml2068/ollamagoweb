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

    //list models and decode json data
    models_list,err:=client.List(ctx)
    if err != nil {
        log.Fatal(err)
    }


 	b, err := json.MarshalIndent(models_list, "   ", "   ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nmodels_list_all: \n")
	fmt.Println(string(b))


    rawdata, err := json.Marshal(models_list)
	var m map[string]interface{}
	var data =[]byte(rawdata)
	err = json.Unmarshal(data, &m)
	if err != nil {
		log.Fatal(err)
	}

	models := m["models"].([]interface{})

	fmt.Println("\nmodels_name_list: \n")
	
	for i, model := range models {
		modelMap := model.(map[string]interface{})
		name := modelMap["name"].(string)
		digest := modelMap["digest"].(string)
		fmt.Println(i+1,name+"/---/"+digest)
	}
}
