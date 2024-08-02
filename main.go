package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"text/template"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
	"github.com/ollama/ollama/api"
)

// initialise to load environment variable from .env file
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
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
	client, err := api.ClientFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}
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
	
	client, err := api.ClientFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	w.Header().Add("mime-type", "text/event-stream")
	f := w.(http.Flusher)
	
	req := &api.GenerateRequest{
		Model:  os.Getenv("llm"),
		Prompt:prompt.Input,
		Options: map[string]interface{}{
			"Seed ": 5,
			"Temperature":0.1,
			//other options please check ollama doc
		},
	}

	respFunc := func(resp api.GenerateResponse) error {
		c := []byte(resp.Response)
		w.Write(c)
		f.Flush()
		return nil
	}
	err = client.Generate(ctx, req, respFunc)
	if err != nil {
		log.Fatal(err)
	}

}
