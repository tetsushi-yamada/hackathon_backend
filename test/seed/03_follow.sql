DROP TABLE IF EXISTS follows;

CREATE TABLE IF NOT EXISTS follows (
    user_id char(36) not null,
    follow_id char(36) not null,
    created_at timestamp not null default current_timestamp,
    PRIMARY KEY (user_id, follow_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (follow_id) REFERENCES users(user_id) ON DELETE CASCADE
    );

INSERT INTO follows (user_id,follow_id) VALUES ('2','100');
INSERT INTO follows (user_id,follow_id) VALUES ('101','100');