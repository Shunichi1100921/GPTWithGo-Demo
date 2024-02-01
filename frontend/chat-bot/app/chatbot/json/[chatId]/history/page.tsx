import { JSONChatHistory } from "@/app/ui/chat-history";


export default async function Page({ params }: { params: { chatId: string }} ) {
    return <JSONChatHistory chatId={ params.chatId }/>;
}