'use client';
import React, { useState } from "react";
import {sendMessageJSONChat} from "@/app/lib/chatService";
import {redirect} from "next/navigation";
import {revalidatePath} from "next/cache";


export default function ChatInput({ chatId }: { chatId: string }) {
    const [message, setMessage] = useState('')

    const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setMessage(event.target.value);
    }

    const handleSendMessage = async () => {
        console.log("Sending message: ", message);
        try {
            const response = await sendMessageJSONChat(chatId, message);
        } catch (error) {
            console.error('Error sending message: ', error)
        }

        revalidatePath(`/chatbot/json/${chatId}/chat`)
        redirect(`/chatbot/json/${chatId}/history`)
    }

    return (
        <form>
            <input
                className="w-full"
                id="chat-input"
                type="text"
                value={message}
                onChange={handleInputChange}
                placeholder={"Message to ChatBot..."}
            />
            <button onClick={handleSendMessage} className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">↑</button>
        </form>
    )
}
