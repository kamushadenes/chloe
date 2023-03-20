CREATE TABLE messages
(
    id               TEXT PRIMARY KEY,
    created_at       DATETIME DEFAULT CURRENT_TIMESTAMP,
    user_id          TEXT NOT NULL,
    role             TEXT NOT NULL,
    content          TEXT NULL,
    summary          TEXT NULL,
    chain_of_thought TEXT NULL
);