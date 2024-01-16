import React from "react";
import {fetchJSONChatHistory, sendMessageJSONChat} from "@/app/lib/chatService";
import JSONChatHistory from "@/app/ui/chat-history";
import ChatInput from "@/app/ui/chat-input";
// import SendMessageButton from "@/app/ui/send-message-button";

interface ChatHistory {
    user_prompt: string;
    bot_response: string;
}

export default function Page({ params }: { params: { chatId: string }}) {
    const handleSendMesage = () => {
        console.log("Clicked!!!!!");
        console.log("Sending message: ");
        try {
            const response = sendMessageJSONChat(params.chatId, 'ありがとうございます。');
        } catch (error) {
            console.error('Error sending message: ', error)
        }
    };

    return (
        <div>
            <h1>JSON Chat for Chat ID: {params.chatId}</h1>
            <JSONChatHistory chatId={params.chatId}/>
            <ChatInput chatId={params.chatId}/>
            {/*<SendMessageButton onClick={handleSendMessage}/>*/}
            {/*<p>current message: { currentMessage }</p>*/}
        </div>
    );
}