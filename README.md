# GPTWithGo-Demo
The demo code to use Open AI API written in Go.

# SetUp
1. Setup MySQL with docker using docker-compose.yml.
```bash
docker-compose up -d
```
2. Set OpenAI API Key as environment variable.
```bash
OPENAI_API_KEY=YOUR_API_KEY
```
3. Execute go file.
  - If you want to use Streaming response, code `func main` in `main.go` as below.
```main.go
func main() {
    chatbotDemo.ChatStream()
    // chatbotDemo.ChatJSON()
}
```
  - If you want to use JSON response, code `func main` in `main.go` as below.
```main.go
func main() {
    // chatbotDemo.ChatStream()
    chatbotDemo.ChatJSON()
}
```
5. Build go file.
```bash
go build main.go
```
6. Run app.
```bash
./demo
```
