package chatBotAPI

type ChatInput struct {
	ChatID  int    `json:"chat_id"`
	Message string `json:"message"`
}

type ChatHistory struct {
	UserPrompt  string `json:"user_prompt"`
	BotResponse string `json:"bot_response"`
}

type JSONChatResponse struct {
	Answer   string `json:"answer"`
	Feedback string `json:"feedback"`
}
