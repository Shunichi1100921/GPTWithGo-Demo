'use client';

import React, {useState} from "react";
// import { sendMessageStreamChat } from "@/app/lib/chatService";
import ChatInput from "@/app/ui/chat-input";
import {revalidatePath} from "next/cache";
import {redirect} from "next/navigation";
import axios from "axios";
import {fetchEventSource} from "@microsoft/fetch-event-source";

export interface ChatMessage {
    sender: 'user' | 'bot';
    message: string;
}

export default function StreamChatInterface({ chatId }: { chatId: string }) {

    const [chatLog, setChatLog] = useState<ChatMessage[]>([]);
    const [currentMessage, setCurrentMessage] = useState<string>("");
    const [currentAnswer, setCurrentAnswer] = useState<string>("");

    const handleSendMessage = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        if (currentMessage === "") {
            return;
        }

        setChatLog([...chatLog, {sender: 'user', message: currentMessage}])

        try {
            const response = await fetch('http://localhost:8080/chat/stream', {
                    headers: {'Content-Type': 'application/json', Accept: "text/event-stream"},
                    method: 'POST',
                    body: JSON.stringify({chat_id: +chatId, message: currentMessage}),
            })

            if (response.body) {
                console.log('response.body: ', response.body)
                const reader = response.body.getReader();
                const decoder = new TextDecoder();
                const read = async () => {
                    while (true) {
                        console.log('reading...')
                        const {done, value} = await reader.read();
                        if (done) {
                            console.log("Stream Complete.");
                            setCurrentAnswer((preAnswer) => {
                                setChatLog([
                                    {sender: 'user', message: currentMessage},
                                    {sender: 'bot', message: preAnswer}
                                ])
                                return "";
                            })
                            setCurrentMessage("")
                            break;
                        }
                        const decoded = decoder.decode(value, {stream: true});
                        const lines = decoded.split('\n');
                        lines.forEach((line) => {
                            if (line.startsWith('data:')) {
                            const message = line.replace('data:', '').trim();
                            setCurrentAnswer((preAnswer) => preAnswer + message);
                            }
                        });
                    }
                    reader.releaseLock();
                }
                read();
            }

            } catch (error) {
                console.error("Error sending message: ", error)
            }

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

    function renderBotAnswer(currentAnswer: string) {
        if (currentAnswer) {
            return (
                <div>
                    <h2>bot answer</h2>
                    <p>bot: {currentAnswer}</p>
                    <p>type: {typeof currentAnswer}</p>
                </div>
            )
        } else {
            return ;
        }
    }

    return (
        <div>
            {renderChatHistory(chatLog)}
            <h2>CurrentAnswer</h2>
            {renderBotAnswer(currentAnswer)}
            <ChatInput
                value={currentMessage}
                onChange={(e) => setCurrentMessage(e.target.value)}
                onSubmit={handleSendMessage}
            />
        </div>
    );
}
