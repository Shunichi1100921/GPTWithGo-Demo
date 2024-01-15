// 'use client'
// import React, { useEffect, useState } from "react";
import { fetchStreamChatHistory } from "@/app/lib/chatService";

interface ChatEntry {
    user_prompt: string;
    bot_response: string;
}

export default async function Page({ params }: { params: { chatId: string }} ) {
    const chatHistory: ChatEntry[] = await fetchStreamChatHistory(params.chatId)

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
            <h1>Stream Chat History for Chat ID: {params.chatId}</h1>
            <div>
                {renderChatHistory()}
            </div>
        </div>
    );
};