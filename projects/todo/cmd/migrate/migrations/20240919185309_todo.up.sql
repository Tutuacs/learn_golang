CREATE TABLE IF NOT EXISTS `todos` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `userId` INT UNSIGNED,
    `name` VARCHAR(255) NOT NULL,
    `done` BOOLEAN DEFAULT false,

    PRIMARY KEY (id),
    FOREIGN KEY (userId) REFERENCES users(id)
);