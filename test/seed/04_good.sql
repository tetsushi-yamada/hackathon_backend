DROP TABLE IF EXISTS goods;

CREATE TABLE IF NOT EXISTS goods (
    tweet_id char(36) not null,
    user_id char(36) not null,
    created_at timestamp not null default current_timestamp,
    PRIMARY KEY (tweet_id, user_id),
    FOREIGN KEY (tweet_id) REFERENCES tweets(tweet_id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
    );

INSERT INTO goods (tweet_id, user_id) VALUES ('2','1000');
INSERT INTO goods (tweet_id, user_id) VALUES ('3','1000');