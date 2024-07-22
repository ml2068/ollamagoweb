<img src="static/logo.png" alt="logo" width="150px"/>

ollamagoweb is a simple ChatGPT clone built with Go. It uses Llama-compatible LLMs, through ollama.

Here are some screenshots.

<img src="assets/screenshot01.png" alt="ollamagoweb main page" width="60%"/>

Starting screen.

<img src="assets/screenshot02.png" alt="What is Golang?" width="60%"/>

Answering the question "What is the Golang?"

# Models

You can download various LLM models from https://www.ollama.com

# edit .env 
You can setting the port and llm by edit the .env file

`PORT=1102`

`llm=codellama:latest`

# run the servcie
`cd ollamagoweb`

`go run main.go`

# build the .exe

`cd ollamagoweb`

`go build main.go`

`./main`
# Groq support
<img src="GroqGoWeb/groqfast.png" alt="Groq go fast" width="60%"/>

#1、replace file from GroqGoWeb as below:

`./main.go`

`./.env`

`./static/logo.png`

#2、edit the .env file:

`PORT=1102`

`llm=llama3-8b-8192`

`baseUrl=https://api.groq.com/openai/v1`

`apiKey=gsk_rmxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx0H`

#3、and run 

`go run main.go`
# Groq
you can get Groq free api-key from https://www.groq.com
