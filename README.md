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
3. Execute the go file.
  - If you want to use a Streaming response, code `func main` in `main.go` as below.
```main.go
func main() {
    chatbotDemo.ChatStream()
    // chatbotDemo.ChatJSON()
}
```
  - If you want to use a JSON response, code `func main` in `main.go` as below.
```main.go
func main() {
    // chatbotDemo.ChatStream()
    chatbotDemo.ChatJSON()
}
```
5. Build a go file.
```bash
go build
```
6. Run the app.
```bash
./demo
```
