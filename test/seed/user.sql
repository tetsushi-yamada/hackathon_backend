DROP TABLE IF EXISTS users;

CREATE TABLE IF NOT EXISTS users (
    user_id char(36) not null,
    user_name varchar(32) not null,
    email varchar(255) not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp on update current_timestamp,
    PRIMARY KEY (user_id)
    );

INSERT INTO users (user_id, user_name, email) VALUES ('1', 'JohnDoe', 'john@example.com');
