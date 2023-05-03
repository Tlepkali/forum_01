CREATE TABLE session (
    uuid TEXT PRIMARY KEY,
    user_id INTEGER NOT NULL,
    expire_at DATETIME NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE
);