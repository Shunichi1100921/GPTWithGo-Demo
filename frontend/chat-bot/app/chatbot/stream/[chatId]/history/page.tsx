import { StreamChatHistory } from "@/app/ui/chat-history";

export default async function Page({ params }: { params: { chatId: string }} ) {
    return <StreamChatHistory chatId={ params.chatId }/>;
}