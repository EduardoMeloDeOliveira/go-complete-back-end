CREATE TABLE IF NOT EXISTS events (
    id SERIAL PRIMARY KEY,
    owner_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    date TIMESTAMP NOT NULL,
    location TEXT NOT NULL,
    CONSTRAINT fk_owner
      FOREIGN KEY (owner_id) REFERENCES users (id) ON DELETE CASCADE
);
