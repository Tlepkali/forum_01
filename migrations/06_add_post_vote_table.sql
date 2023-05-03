create table if not exists post_vote (
    author_id INTEGER NOT NULL,
    post_id INTEGER,
    status INTEGER NOT NULL,
    PRIMARY KEY (author_id, post_id),
    FOREIGN KEY (author_id) REFERENCES user(id) ON DELETE CASCADE,
    FOREIGN KEY (post_id) REFERENCES post(id) ON DELETE CASCADE
);