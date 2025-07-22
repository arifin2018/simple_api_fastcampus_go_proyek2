CREATE TABLE IF NOT EXISTS user_activities (
    id int NOT NULL AUTO_INCREMENT,
    post_id int NOT NULL,
    user_id int NOT NULL,
    is_liked boolean NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    created_by VARCHAR(255) NOT NULL,
    updated_by VARCHAR(255) NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT fk_user_activity_post FOREIGN KEY (post_id) 
        REFERENCES posts(id) 
        ON DELETE CASCADE 
        ON UPDATE CASCADE,
    CONSTRAINT fk_user_activity_user FOREIGN KEY (user_id) 
        REFERENCES users(id) 
        ON DELETE CASCADE 
        ON UPDATE CASCADE
);