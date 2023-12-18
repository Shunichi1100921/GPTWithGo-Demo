# GPTWithGo-Demo
The demo code to use Open AI API written in Go.
The normal ChatBot response with Stream by GPT-3 and the ChatBot which returns answer and feedback in JSON format are implemented.

## Console version.

### SetUp
1. Build Docker.
```bash
docker-compose down
docker-compose up -d
```
2. Set OpenAI API Key as an environment variable.
```bash
export OPENAI_API_KEY=YOUR_API_KEY
```
3. Execute the go file. 
```bash
go run cmd/console/main.go
```
If you want to run ChatBot which stores to Database and returns JSON file, run
```bash
go run main.go --json
```

## WebAPI Version.
### SetUp
1. Build Docker.
```bash
docker-compose down
docker-compose up -d
```
2. Set OpenAI API Key as an environment variable.
```bash
export OPENAI_API_KEY=YOUR_API_KEY
```
3. Execute the go file.
```bash
go run cmd/webApp/main.go
```

### Endpoints
1. `/chat/stream`
    - POST: Send 'chat_id', 'message' to the chatbot with JSON body and return the response by SSE event.
2. `/chat/strea/history?id={chat_id}`
    - GET: Return the chat history of the chat_id.
2. `/chat/json`
    - POST: Send 'chat_id', 'message' to the chatbot with JSON body and return the response and feedback with JSON Format. 
3. `/chat/json/history?id={chat_id}`
    - GET: Return the chat history of the chat_id.
## How to open database.
1. Open bash in the container.
```bash
docker exec -it demo-mysql-1 bash
```
2. Login to MySQL.
```bash
mysql -p demoSQL
Enter password: 
mysql> 
```

