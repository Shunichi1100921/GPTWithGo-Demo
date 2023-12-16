# GPTWithGo-Demo
The demo code to use Open AI API written in Go.

# SetUp
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
go run main.go
```
If you want to run ChatBot which stores to Database and returns JSON file, run
```bash
go run main.go --json
```

## How to open database.
```bash
docker exec -it demo-mysql-1 bash
```

```bash
mysql -p demoSQL
Enter password: 
mysql> 
```

