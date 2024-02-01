import React, {ChangeEventHandler, FormEventHandler } from "react";


export default function ChatInput({ value, onChange, onSubmit }: {
    value: string,
    onChange: ChangeEventHandler<HTMLInputElement>,
    onSubmit: FormEventHandler,
}) {

    return (
        <form onSubmit={onSubmit}>
            <input
                className="w-full text-black"
                id="chat-input"
                type="text"
                value={value}
                onChange={onChange}
                placeholder={"Message to ChatBot..."}
            />
            <button
                type={"submit"}
                className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
            >
                ↑
            </button>
        </form>
    )
}
