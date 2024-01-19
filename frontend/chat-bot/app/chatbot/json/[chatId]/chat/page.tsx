import React from "react";
import JSONChatInterface from "@/app/ui/json-chat-interface";
import { JSONChatHistory } from "@/app/ui/chat-history";

export default function Page({ params }: { params: { chatId: string }}) {
    return (
        <div className="flex flex-col h-screen">
            <div className="p-3 bg-gray-800 text-white">
                <h1>JSON Chat for Chat ID: {params.chatId}</h1>
                <JSONChatHistory chatId={params.chatId}/>
                <JSONChatInterface chatId={params.chatId}/>
            </div>
        </div>
    );
}