DROP TABLE IF EXISTS users;

CREATE TABLE IF NOT EXISTS users (
    user_id char(36) not null,
    user_name varchar(32) not null,
    user_description char(36) default null,
    is_private boolean not null default false,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp on update current_timestamp,
    PRIMARY KEY (user_id)
    );

INSERT INTO users (user_id, user_name) VALUES ('1', 'JohnDoe');
INSERT INTO users (user_id, user_name) VALUES ('2', 'Te');
INSERT INTO users (user_id, user_name) VALUES ('3', 'TEST TWEET');
INSERT INTO users (user_id, user_name) VALUES ('100', 'FOLLOW ME');
INSERT INTO users (user_id, user_name) VALUES ('101', 'FOLLOW YOU');
INSERT INTO users (user_id, user_name) VALUES ('1000', 'GOOD YOU');
INSERT INTO users (user_id, user_name) VALUES ('10000', 'UPDATE');
INSERT INTO users (user_id, user_name) VALUES ('100000', 'SEARCH');
INSERT INTO users (user_id, user_name) VALUES ('11', 'REPLY');
INSERT INTO users (user_id, user_name) VALUES ('111', 'UPDATE USERNAME');
INSERT INTO users (user_id, user_name, user_description) VALUES ('1111', 'description','my name is description');