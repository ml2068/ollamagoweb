package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"math"
	"text/template"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
	"github.com/ollama/ollama/api"
)

var client *api.Client

// initialise to load environment variable from .env file
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	client, err = api.ClientFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Handle("/static/*", http.StripPrefix("/static",
		http.FileServer(http.Dir("./static"))))
	r.Get("/", index)
	r.Post("/run", run)
	log.Println("\033[93mOllama go web serve started. Press CTRL+C to quit.\033[0m")
	log.Println("Local URL: http://localhost:"+os.Getenv("PORT"))
	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}

// index.html
func index(w http.ResponseWriter, r *http.Request) {
	llm := os.Getenv("llm")
	ollamaversion, _ :=client.Version(context.Background())
	t, _ := template.ParseFiles("static/index.html")
	data := map[string]interface{}{
		"llm":    llm,
		"Ollav": ollamaversion,
	}
	err := t.Execute(w, data)  
    	if err != nil {
        	log.Println(err)
    	}
}

// get LLM Context Length
func getContextLength() int {
    ctx := context.Background()
    sq := &api.ShowRequest{
        Model: os.Getenv("llm"),
    }
    model, err := client.Show(ctx, sq)
    if err != nil {
        log.Println(err)
    }
    var num = model.ModelInfo["llama.context_length"].(float64)
    clen := math.Min(8192, math.Max(num, 0))
    return int(clen)
}

//get Option Setting function
func GetOptionSetting(client *api.Client) (map[string]interface{}, error) {
	clen :=getContextLength()
	options_setting := map[string]interface{}{
		"Runner": map[string]interface{}{
			"NumCtx":    clen,
			"NumBatch":  16,
			"NumGPU":    1,
			"MainGPU":   0,
			"LowVRAM":   false,
			"F16KV":     true,
			"LogitsAll": true,
			"VocabOnly": false,
			"NumThread": 2,
		},
		"NumKeep":         256,
		"Seed":            42,
		"NumPredict":      10,
		"TopK":            5,
		"TopP":            0.8,
		"MinP":            0.1,
		"TFSZ":            0.5,
		"TypicalP":        0.9,
		"RepeatLastN":     4,
		"Temperature":     1.2,
		"RepeatPenalty":   0.6,
		"PresencePenalty": 0.7,
		"FrequencyPenalty": 0.8,
		"Mirostat":        1,
		"MirostatTau":     5.0,
		"MirostatEta":     0.4,
		"PenalizeNewline": false,
		"Stop":            []string{"\n"},
	}
	return options_setting, nil
}


// call the LLM and return the response
func run(w http.ResponseWriter, r *http.Request) {
	prompt := struct {
		Input string `json:"input"`
	}{}
	// decode JSON from client
	err := json.NewDecoder(r.Body).Decode(&prompt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	ctx := context.Background()
	
	w.Header().Add("mime-type", "text/event-stream")
	f := w.(http.Flusher)
	
	options_setting, err := GetOptionSetting(client)
	if err != nil {
		log.Fatal(err)
	}
	
	req := &api.GenerateRequest{
		Model:  os.Getenv("llm"),
		Prompt:prompt.Input,
		Options: options_setting,
	}

	respFunc := func(resp api.GenerateResponse) error {
		c := []byte(resp.Response)
		w.Write(c)
		f.Flush()
		if resp.Done{
			log.Println("\nMetrics:")
			log.Println("TotalDuration:", resp.Metrics.TotalDuration)
			log.Println("LoadDuration:", resp.Metrics.LoadDuration)
			log.Println("PromptEvalCount:", resp.Metrics.PromptEvalCount)
			log.Println("PromptEvalDuration:", resp.Metrics.PromptEvalDuration)
			log.Println("EvalCount:", resp.Metrics.EvalCount)
			log.Println("EvalDuration:", resp.Metrics.EvalDuration)
			log.Printf("Speed:%.2f tokens/s\n", float64(resp.Metrics.EvalCount)/resp.Metrics.EvalDuration.Seconds())
		}
		return nil
	}
	err = client.Generate(ctx, req, respFunc)
	if err != nil {
		log.Fatal(err)
	}

}
