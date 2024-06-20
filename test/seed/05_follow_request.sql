DROP TABLE IF EXISTS follow_requests;

CREATE TABLE IF NOT EXISTS follow_requests (
    user_id char(36) not null,
    follow_id char(36) not null,
    status varchar(32) not null default 'pending',
    created_at timestamp not null default current_timestamp,
    PRIMARY KEY (user_id, follow_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (follow_id) REFERENCES users(user_id) ON DELETE CASCADE
    );