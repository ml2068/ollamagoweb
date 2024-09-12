<h3 align="center">
  <img
    src="https://github.com/ml2068/ollamagoweb/blob/main/static/logo.png"
    height="200"
  >
</h3>
<div align="center">
<a href="https://github.com/ml2068/ollamagoweb/blob/main/LICENSE">
  <img src="https://img.shields.io/github/license/ml2068/ollamagoweb" alt="License">
</a>
<a href="https://github.com/ml2068/ollamagoweb/stargazers">
  <img src="https://img.shields.io/github/stars/ml2068/ollamagoweb.svg" alt="GitHub Contributors">
</a>
<a href="https://github.com/ml2068/ollamagoweb/forks">
  <img src="https://img.shields.io/github/forks/ml2068/ollamagoweb.svg" alt="GitHub Contributors">
</a>
<a href="https://github.com/ml2068/ollamagoweb/graphs/contributors">
  <img src="https://img.shields.io/github/contributors/ml2068/ollamagoweb.svg" alt="GitHub Contributors">
</a>
<a href="https://github.com/ml2068/ollamagoweb">
  <img src="https://badgen.net/badge/Open%20Source%20%3F/Yes%21/blue?icon=github" alt="Open Source">
</a>
</div>

# Introduction

ollamagoweb, a simple ChatGPT clone built in Go that leverages Llama-compatible LLM via the ollama service. This innovative tool provides a seamless conversation experience and features:

**Simple interface**: 

The main page displays the Ollama version, LLM tag, and context length, providing essential information for a productive conversation.

**Answering Questions**: 

Contextual Discussion ollamagoweb effortlessly responds to questions, such as "What is TOTP?", and continues the discussion in context, allowing for a natural and engaging exchange.

**Conversation Management**: 

For each round of dialogue, users can easily delete the dialogue by clicking the button in the upper right corner to ensure that llm's ideas are logically coherent during the dialogue.

**Conversation saved**: 

Ollamagoweb allows users to save conversations as HTML documents for later reference, providing a convenient way to review and analyze previous discussions.

**Generate Response log**: 

The backend server efficiently displays and calculates the session's token and speed, providing a robust foundation for the application.

Here are some screenshots.

<img src="assets/screenshot01.png" alt="ollamagoweb main page" width="60%"/>

<img src="assets/screenshot02.png" alt="What is TOTP?" width="60%"/>

<img src="assets/screenshot03.png" alt="backend log" width="60%"/>

# Models

Download various LLM models from https://www.ollama.com

# Git Clone
```
git clone https://github.com/ml2068/ollamagoweb.git
cd ollamagoweb
go mod tidy
```
# edit .env 

Setting the port and llm tag by edit the .env file
```
vim .env
PORT=1102
llm=llama3.1:8b
```
# run the servcie

`go run main.go`

# build the .exe
```
cd ollamagoweb
go build main.go
./main
```
# Groq support

<img src="GroqGoWeb/groqfast.png" alt="Groq go fast" width="60%"/>

1. replace file from GroqGoWeb as below:
```
./main.go
./.env
./static/logo.png
./static/favicon.ico
```
2. edit the .env file:
```
PORT=1102
llm=llama3-8b-8192
baseUrl=https://api.groq.com/openai/v1
apiKey=gsk_rmxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx0H
```
3. go run 

`go run main.go`

4. api-key

Get Groq free api-key from https://www.groq.com

# deepinfra  support

Get deepinfra api-key from https://deepinfra.com/

.env format:
```
PORT=1102
llm=meta-llama/Meta-Llama-3.1-405B-Instruct
baseUrl=https://api.deepinfra.com/v1/openai
apiKey=JZxxxxxxxxxxxxxxxxxxxxxwHrK
```
## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=ml2068/ollamagoweb&type=Date)](https://star-history.com/#ml2068/ollamagoweb&Date)
