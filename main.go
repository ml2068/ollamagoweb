package main

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"net/http"
	"os"
	"text/template"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

// initialise to load environment variable from .env file
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	//get VPS ip adress 
	conn, error := net.Dial("udp", "8.8.8.8:80")  
    	if error != nil {  
    	log.Fatal("error")
    	}	  
    	defer conn.Close()  
    	ipAddress:=conn.LocalAddr().(*net.UDPAddr).IP 
	
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Handle("/static/*", http.StripPrefix("/static",
		http.FileServer(http.Dir("./static"))))
	r.Get("/", index)
	r.Post("/run", run)
	log.Println("\033[93mGroq go web started. Press CTRL+C to quit.\033[0m")
	log.Printf("URL:%s:"+os.Getenv("PORT"),ipAddress)
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

	llm, err := ollama.New(ollama.WithModel(os.Getenv("llm")))
	if err != nil {
		log.Println("Cannot create local LLM:", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("mime-type", "text/event-stream")
	f := w.(http.Flusher)
	llm.Call(context.Background(), prompt.Input,
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			w.Write(chunk)
			f.Flush()
			return nil
		}), llms.WithMaxTokens(4096), llms.WithTemperature(0.5))
}
