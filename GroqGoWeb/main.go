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
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
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
	log.Println("\033[93mOllamagoweb started. Press CTRL+C to quit.\033[0m")
	log.Println("Local URL: http://localhost:"+os.Getenv("PORT"))
	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}

// index
func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("static/index.html")
	t.Execute(w, nil)
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

	llm, err := openai.New(
		openai.WithModel(os.Getenv("llm")),
		openai.WithBaseURL(os.Getenv("baseUrl")),
		openai.WithToken(os.Getenv("apiKey")),
	)
	if err != nil {
		log.Println("Cannot create local LLM:", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("mime-type", "text/event-stream")
	f := w.(http.Flusher)
	llms.GenerateFromSinglePrompt(context.Background(),llm, prompt.Input,
		llms.WithMaxTokens(8129), llms.WithTemperature(0.5),
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			w.Write(chunk)
			f.Flush()
			return nil
		}), )
}
