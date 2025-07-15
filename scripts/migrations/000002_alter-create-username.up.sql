ALTER TABLE Users
ADD username varchar(255);

ALTER TABLE Users
ADD CONSTRAINT UNIQUE unique_username (username)