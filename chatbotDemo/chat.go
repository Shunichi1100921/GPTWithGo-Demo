package chatbotDemo

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"io"
	"os"
)

type Content struct {
	Answer   string
	Feedback string
}

// promptUser gets user input from the console.
func promptUser(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

// createOPENAIClient creates and returns an OpenAI client.
func createOpenAIClient(apiKey string) *openai.Client {
	return openai.NewClient(apiKey)
}

// createInitialMessage create initial chat messages.
func createInitialMessage() []openai.ChatCompletionMessage {
	return []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "Answer my question.",
		},
	}
}

// createInitialMessageForJSON create initial chat messages which make ChatBot return in JSON format.
func createInitialMessageForJSON() []openai.ChatCompletionMessage {
	return []openai.ChatCompletionMessage{
		{
			Role: openai.ChatMessageRoleUser,
			Content: "You're the boss of me who answers my questions.  You gotta give me the answers and feedback " +
				"about the way I asked. You should return in JSON format, which has two keys.  One is the answer and " +
				"the other is feedback.",
		},
	}
}

// mainChatLoop handles the main logic of the streaming chat loop.
func mainStreamChatLoop(ctx context.Context, c *openai.Client, messages []openai.ChatCompletionMessage) {
	for {
		userInput := promptUser("You: ")
		if userInput == "exit" {
			return
		}

		// Manage the conversation.
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: userInput,
		})

		//Send the request to OpenAI.
		req := openai.ChatCompletionRequest{
			Model:     openai.GPT3Dot5Turbo,
			MaxTokens: 1000,
			Messages:  messages,
			Stream:    true,
		}
		stream, err := c.CreateChatCompletionStream(ctx, req)
		if err != nil {
			fmt.Printf("ChatCompletionStream error: %v\n", err)
			return
		}
		defer stream.Close()

		// Display the bot's response.
		fmt.Print("Chatbot: ")
		for {
			response, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				break
			}

			fmt.Print(response.Choices[0].Delta.Content)
		}
		fmt.Println("")
	}
}

// mainChatJSONLoop handles the main logic of the chat loop which returns in JSON format.
func mainChatJSONLoop(ctx context.Context, c *openai.Client, messages []openai.ChatCompletionMessage) {
	for {
		userInput := promptUser("You: ")
		if userInput == "exit" {
			return
		}

		// Manage the conversation.
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: userInput,
		})

		//Send the request to OpenAI.
		req := openai.ChatCompletionRequest{
			Model:          openai.GPT3Dot5Turbo1106,
			MaxTokens:      1000,
			Messages:       messages,
			ResponseFormat: &openai.ChatCompletionResponseFormat{Type: openai.ChatCompletionResponseFormatTypeJSONObject},
		}

		response, err := c.CreateChatCompletion(ctx, req)
		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			return
		}

		msg := response.Choices[0].Message.Content
		msgByte := []byte(msg)

		var content Content
		err = json.Unmarshal(msgByte, &content)
		if err != nil {
			fmt.Printf("JSON Unmarshal error: %v\n", err)
		}

		SaveContent(content)

		// Display the ChatBot's response.
		fmt.Println("ChatBot: " + content.Answer)
		fmt.Println("Feedback: " + content.Feedback)
	}

}

func ChatStream() {
	// Gain OpenAI API key from environment variable.
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("OpenAI API is not set.")
		return
	}

	client := createOpenAIClient(apiKey)
	ctx := context.Background()
	messages := createInitialMessage()

	mainStreamChatLoop(ctx, client, messages)

}

// ChatJSON start ChatBot which returns answers in JSON format.
func ChatJSON() {

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("OpenAI API is not set.")
		return
	}

	client := createOpenAIClient(apiKey)
	ctx := context.Background()
	messages := createInitialMessageForJSON()

	mainChatJSONLoop(ctx, client, messages)

}
