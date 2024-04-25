DROP TABLE IF EXISTS follows;

CREATE TABLE IF NOT EXISTS follows (
    user_id char(36) not null,
    follow_id char(36) not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp on update current_timestamp,
    PRIMARY KEY (user_id)
    );

INSERT INTO follows (user_id,follow_id) VALUES ('2','100');