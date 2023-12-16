CREATE TABLE IF NOT EXISTS demoTable (
    id INT AUTO_INCREMENT PRIMARY KEY,
    answer TEXT NOT NULL,
    feedback TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS chatHistory (
    id INT,
    user_id INT,
    user_prompt TEXT NOT NULL,
    bot_response TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)
