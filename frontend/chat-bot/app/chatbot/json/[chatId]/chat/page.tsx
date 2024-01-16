import React from "react";
import { JSONChatHistory } from "@/app/ui/chat-history";
import ChatInput from "@/app/ui/chat-input";

export default function Page({ params }: { params: { chatId: string }}) {
    return (
        <div>
            <h1>JSON Chat for Chat ID: {params.chatId}</h1>
            <JSONChatHistory chatId={params.chatId}/>
            <ChatInput chatId={params.chatId}/>
        </div>
    );
}