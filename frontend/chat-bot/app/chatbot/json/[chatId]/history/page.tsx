import { fetchJSONChatHistory } from "@/app/lib/chatService";

interface ChatEntry {
    user_prompt: string;
    bot_response: string;
}

export default async function Page({ params }: { params: { chatId: string }} ) {
    const chatHistory: ChatEntry[] = await fetchJSONChatHistory(params.chatId)

    function renderChatHistory() {
        if (chatHistory) {
            return chatHistory.map((entry, index) => (
                <div key={index}>
                    <p>User: {entry.user_prompt}</p>
                    <p>Bot: {entry.bot_response}</p>
                </div>
            ));
        } else {
            return <p>No chat with this id.</p>
        }
    }

    return (
        <div>
            <h1>Chat History with Feedback for Chat ID: {params.chatId}</h1>
            <div>
                {renderChatHistory()}
            </div>
        </div>
    );
}