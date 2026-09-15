CREATE TABLE db.public.purchases (
    username VARCHAR(200),
    merch VARCHAR(200),
    quantity INT, 
    FOREIGN KEY (username) REFERENCES users (username),
    FOREIGN KEY (merch) REFERENCES merch(name),
    PRIMARY KEY (username, merch)
);