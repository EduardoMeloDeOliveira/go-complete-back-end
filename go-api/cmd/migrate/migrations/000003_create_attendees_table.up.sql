CREATE TABLE IF NOT EXISTS attendees(
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    event_id INTEGER NOT NULL,
    CONSTRAINT fk_user
        FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    CONSTRAINT fk_event 
        FOREIGN KEY (event_id) REFERENCES events (id) ON DELETE CASCADE
)