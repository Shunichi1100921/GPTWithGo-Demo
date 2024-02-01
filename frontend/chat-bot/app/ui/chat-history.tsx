import React from "react";
import {fetchJSONChatHistory, fetchStreamChatHistory} from "@/app/lib/chatService";

interface ChatHistory {
    user_prompt: string;
    bot_response: string;
}

export async function JSONChatHistory({ chatId }: {chatId: string}) {
    const history = await fetchJSONChatHistory(chatId);
    return <ChatHistory history={history}/>;
}

export async function StreamChatHistory({ chatId }: {chatId: string}) {
    const history = await fetchStreamChatHistory(chatId);
    return <ChatHistory history={history}/>;
}


export async function ChatHistory({ history }: { history: ChatHistory[] }) {

    function renderChatHistory() {
        if (history) {
            return history.map((entry: ChatHistory, index: number) => (
                <div key={index}>
                    <p>User: {entry.user_prompt}</p>
                    <p>Bot: {entry.bot_response}</p>
                </div>
            ));
        } else {
            return <p>No chat with this id.</p>;
        }
    }

    return (
        <div>
            {renderChatHistory()}
        </div>
    )
}