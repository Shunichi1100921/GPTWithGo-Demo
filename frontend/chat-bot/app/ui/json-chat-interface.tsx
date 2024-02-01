'use client';
import ChatInput  from "@/app/ui/chat-input";
import React, {useState} from "react";
import {sendMessageJSONChat} from "@/app/lib/chatService";
import {revalidatePath} from "next/cache";
import {redirect} from "next/navigation";


export interface ChatMessage {
    sender: 'user' | 'bot';
    message: string;
}

export default function JSONChatInterface({ chatId }: { chatId: string }) {

    const [chatLog, setChatLog] = useState<ChatMessage[]>([]);
    const [currentMessage, setCurrentMessage] = useState<string>("");

    const handleSendMessage = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        if (currentMessage === "") {
            return;
        }

        console.log("Sending message: ", currentMessage);
        setChatLog([...chatLog, {sender: 'user', message: currentMessage}])
        try {
            const response = await sendMessageJSONChat(chatId, currentMessage);
            setChatLog([
                ...chatLog,
                {sender: 'user', message: currentMessage},
                {sender: 'bot', message: response.answer}])
            setCurrentMessage("")
        } catch (error) {
            console.error("Error sending message: ", error)
        }

        revalidatePath(`/chatbot/json/${chatId}/chat`)
        redirect(`/chatbot/json/${chatId}/history`);
    }

    function renderChatHistory(chatLog: ChatMessage[]) {
        if (chatLog) {
            return chatLog.map((entry: ChatMessage, index: number) => (
                <div key={index}>
                    <p>{entry.sender}: {entry.message}</p>
                </div>
            ));
        } else {
            return ;
        }
    }

    return (
        <div>
            {renderChatHistory(chatLog)}
            <ChatInput
                value={currentMessage}
                onChange={(e) => setCurrentMessage(e.target.value)}
                onSubmit={handleSendMessage}
            />
        </div>
    );
}
