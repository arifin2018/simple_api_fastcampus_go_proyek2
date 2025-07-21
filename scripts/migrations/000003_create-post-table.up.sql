CREATE TABLE IF NOT EXISTS posts (
    id int NOT NULL AUTO_INCREMENT,
    user_id INT NOT NULL UNIQUE,
    post_title varchar(255) NOT NULL,
    post_content LONGTEXT NOT NULL,
    post_hastags LONGTEXT NOT NULL,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp,
    created_by LONGTEXT not null,
    updated_by LONGTEXT not null,
    PRIMARY KEY (id)
);