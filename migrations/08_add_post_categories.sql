CREATE TABLE post_categories (
  post_id INTEGER,
  category_name INTEGER,
  FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
  FOREIGN KEY (category_name) REFERENCES categories(name) ON DELETE CASCADE
);