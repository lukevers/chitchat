-- +migrate Up
CREATE TABLE messages
(
id INTEGER UNSIGNED NOT NULL AUTO_INCREMENT,
sender INTEGER UNSIGNED NOT NULL,
receiver INTEGER UNSIGNED NOT NULL,
message VARCHAR(255) NOT NULL,

PRIMARY KEY (id),
CONSTRAINT message_sender_foreign FOREIGN KEY (sender) REFERENCES users (id),
CONSTRAINT message_receiver_foreign FOREIGN KEY (receiver) REFERENCES users (id)
) DEFAULT CHARACTER SET utf8, DEFAULT COLLATE utf8_unicode_ci;

-- +migrate Down
DROP TABLE messages;
