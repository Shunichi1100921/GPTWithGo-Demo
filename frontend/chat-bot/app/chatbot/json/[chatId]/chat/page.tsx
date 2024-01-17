import React from "react";
import ChatInterface from "@/app/ui/chat-interface";
import {JSONChatHistory} from "@/app/ui/chat-history";

export default function Page({ params }: { params: { chatId: string }}) {
    return (
        <div className="flex flex-col h-screen">
            <div className="p-3 bg-gray-800 text-white">
                <h1>JSON Chat for Chat ID: {params.chatId}</h1>
                <JSONChatHistory chatId={params.chatId}/>
                <ChatInterface chatId={params.chatId}/>
            </div>
        </div>
    );
}