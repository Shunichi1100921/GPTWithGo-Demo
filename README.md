# GPTWithGo-Demo
The demo code to use Open AI API written in Go.

# SetUp
1. Build Docker.
```bash
docker-compose up -d
```
2. Set OpenAI API Key as an environment variable.
```bash
OPENAI_API_KEY=YOUR_API_KEY
```
3. Execute the go file. If you want to run streaming ChatBot, run
```bash
go run cmd/stream
```
or if you want to run ChatBot which returns JSON file, run
```bash
go run cmd/json
```
4. Build a go file.
```bash
go build
```
5. Run the app.
```bash
./demo
```
