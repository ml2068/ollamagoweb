# Introduction

ollamagoweb, a simple ChatGPT clone built with Go, utilizing Llama-compatible LLMs through ollama. This innovative tool offers a seamless conversational experience, as showcased in the following screenshots:

**Starting Screen**: 

The main page displays the Ollama version, LLM tag, and context length, providing essential information for a productive conversation.

**Answering Questions**: 

Contextual Discussion ollamagoweb effortlessly responds to questions, such as "What is TOTP?", and continues the discussion in context, allowing for a natural and engaging exchange.

**Conversation Management**: 

For each round of dialogue, users can conveniently delete the conversation by clicking the button in the upper-right corner, ensuring a clutter-free interface.

**Backend Server**: 

The backend server efficiently displays and calculates the session's token and speed, providing a robust foundation for the application.

Here are some screenshots.

<img src="assets/screenshot01.png" alt="ollamagoweb main page" width="60%"/>

<img src="assets/screenshot02.png" alt="What is TOTP?" width="60%"/>

<img src="assets/screenshot03.png" alt="backend log" width="60%"/>

# Models

Download various LLM models from https://www.ollama.com

# edit .env 

Setting the port and llm tag by edit the .env file

`PORT=1102`

`llm=codellama:latest`

# run the servcie

`git clone https://github.com/ml2068/ollamagoweb.git`

`cd ollamagoweb`

`vim .env`

`go mod tidy`

`go run main.go`

# build the .exe

`cd ollamagoweb`

`go build main.go`

`./main`

# Groq support

<img src="GroqGoWeb/groqfast.png" alt="Groq go fast" width="60%"/>

1. replace file from GroqGoWeb as below:

`./main.go`

`./.env`

`./static/logo.png`

2. edit the .env file:

`PORT=1102`

`llm=llama3-8b-8192`

`baseUrl=https://api.groq.com/openai/v1`

`apiKey=gsk_rmxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx0H`

3. go run 

`go run main.go`

# Groq

Get Groq free api-key from https://www.groq.com
