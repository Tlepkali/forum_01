create table if not exists comment_vote (
    author_id INTEGER NOT NULL,
    comment_id INTEGER,
    status INTEGER NOT NULL,
    PRIMARY KEY (author_id, comment_id),
    FOREIGN KEY (author_id) REFERENCES user(id) ON DELETE CASCADE,
    FOREIGN KEY (comment_id) REFERENCES comment(id) ON DELETE CASCADE
);