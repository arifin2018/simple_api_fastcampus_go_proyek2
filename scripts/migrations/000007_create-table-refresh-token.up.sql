CREATE TABLE IF NOT EXISTS refresh_token (
    id int NOT NULL AUTO_INCREMENT,
    user_id int NOT NULL,
    refresh_token text NOT NULL,
    expired_at timestamp NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    created_by LONGTEXT NOT NULL,
    updated_by LONGTEXT NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT fk_user_id_refresh_token FOREIGN KEY (user_id) 
        REFERENCES users(id) 
        ON DELETE CASCADE 
        ON UPDATE CASCADE
);