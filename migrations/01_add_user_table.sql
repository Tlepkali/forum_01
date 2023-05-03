-- Add a new users table with id primary key
CREATE TABLE user (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL,
    hashed_pw TEXT NOT NULL,
    email TEXT NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL
);

-- Add a new index on username to enforce uniqueness
CREATE UNIQUE INDEX user_username_uindex ON user (username);

-- Add a new index on email to enforce uniqueness
CREATE UNIQUE INDEX user_email_uindex ON user (email);
