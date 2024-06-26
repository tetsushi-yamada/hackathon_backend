DROP TABLE IF EXISTS tweets;

CREATE TABLE IF NOT EXISTS tweets (
    tweet_id char(36) not null,
    user_id char(36) not null,
    tweet_text varchar(255) not null,
    parent_id char(36) default null,
    retweet_id char(36) default null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp on update current_timestamp,
    PRIMARY KEY (tweet_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
    );

INSERT INTO tweets (tweet_id, user_id, tweet_text) VALUES ('1', '1', 'Hello, World!');
INSERT INTO tweets (tweet_id, user_id, tweet_text) VALUES ('2', '1', 'No way!');
INSERT INTO tweets (tweet_id, user_id, tweet_text) VALUES ('3', '2', 'I am Te!');
INSERT INTO tweets (tweet_id, user_id, tweet_text) VALUES ('4', '10000', 'Will be updated!');
INSERT INTO tweets (tweet_id, user_id, tweet_text) VALUES ('5', '100000', 'Will be searched!');
INSERT INTO tweets (tweet_id, user_id, tweet_text, parent_id) VALUES ('6','11','REPLY','5');
INSERT INTO tweets (tweet_id, user_id, tweet_text) VALUES ('7','3','TEST TWEET');
INSERT INTO tweets (tweet_id, user_id, tweet_text) VALUES ('8','1000','WILL BE RETWEETED!');
INSERT INTO tweets (tweet_id, user_id, tweet_text, retweet_id) VALUES ('9','100','RETWEET','8');
