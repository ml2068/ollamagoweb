package main

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
	"syscall"
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
	log.Println("\033[93mGroq go web serve started. Press CTRL+C to quit.\033[0m")

	// Create a new server
	srv := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: r,
	}

	// Start the server in a goroutine
	go func(ipAddress net.IP) {
		log.Printf("URL:%s:%s", ipAddress, os.Getenv("PORT"))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}(ipAddress)

	// Wait for interrupt signal to stop the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Server is shutting down...")

	// Shut down the server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
	log.Println("Server stopped")
}

// index
func index(w http.ResponseWriter, r *http.Request) {
	llm := os.Getenv("llm")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	t, _ := template.ParseFiles("static/index.html")
	err := t.Execute(w, map[string]string{"llm":llm})  
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
	msg := llms.MessageContent{
        	Role:  llms.ChatMessageTypeHuman,
        	Parts: []llms.ContentPart{
            		llms.TextContent{Text: prompt.Input},
        	},
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
	opts := []llms.CallOption{
	        llms.WithMaxTokens(4069),
	        llms.WithTemperature(0.5),
	        llms.WithTopP(0.8),
	        llms.WithTopK(10),
    	}
	resp, err := llm.GenerateContent(context.Background(),[]llms.MessageContent{msg}, opts...)
	if err != nil {
    		http.Error(w, err.Error(), http.StatusInternalServerError)
    		return
    	}
    
	for _, c := range resp.Choices{
		w.Header().Add("mime-type", "text/event-stream")
		f := w.(http.Flusher)
		w.Write([]byte(c.Content))
		f.Flush()
		for key, value := range c.GenerationInfo {
			log.Println(key, ":", value)
		}
	}

}
