import axios from 'axios';

const apiClient = axios.create({
    baseURL: 'http://localhost:8080',
    headers: {
        'Content-Type': "application/json",
    },
});

export const sendMessageStreamChat = async (chatId: string, message: string) => {
    try {
        const response = await apiClient.post('/chat/stream', { chat_id: chatId, message: message });
        return response.data;
    } catch (error) {
        console.error('API request Error:', error);
        throw error;
    }
};

export const sendMessageJSONChat = async (chatId: string, message: string) => {
    try {
        const response = await apiClient.post('/chat/json', { chat_id: chatId, message: message });
        return response.data;
    } catch (error) {
        console.error('API request Error:', error);
        throw error;
    }
};

export const getStreamChatHistory = async (chatId: string) => {
    try {
        const response = await apiClient.get(`/chat/stream/history?id=${chatId}`);
        return response.data;
    } catch (error) {
        console.error('API request Error:', error);
        throw error;
    }
};

export const getJSONChatHistory = async (chatId: string) => {
    try {
        const response = await apiClient.get(`/chat/json/history?id=${chatId}`);
        return response.data;
    } catch (error) {
        console.error('API request Error:', error);
        throw error;
    }
};
