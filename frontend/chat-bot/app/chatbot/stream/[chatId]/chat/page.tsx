import React from "react";
import ChatInterface from "@/app/ui/json-chat-interface";
import { StreamChatHistory } from "@/app/ui/chat-history";
import StreamChatInterface from "@/app/ui/stream-chat-interface";
import {sendMessageStreamChat} from "@/app/lib/chatService";

export default function Page({ params }: { params: { chatId: string }}) {
    return (
        <div className="flex flex-col h-screen">
            <div className="p-3 bg-gray-800 text-white">
                <h1>Stream Chat for Chat ID: {params.chatId}</h1>
                <StreamChatHistory chatId={params.chatId} />
                <StreamChatInterface chatId={params.chatId} />
            </div>
        </div>
    )
}