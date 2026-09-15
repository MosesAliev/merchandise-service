CREATE TABLE db.public.transfers (
    sent_from VARCHAR(200),
    money int,
    sent_to VARCHAR(200),
    FOREIGN KEY (sent_from) REFERENCES users (username),
    FOREIGN KEY (sent_to) REFERENCES users(username)
);