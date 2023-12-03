CREATE TABLE IF NOT EXISTS `user_token`
(
    `id`            int auto_increment PRIMARY KEY,
    `token`         VARCHAR(500) NOT NULL,
    `refresh_token` VARCHAR(500) NOT NULL,
    `user_id`       int UNIQUE   NOT NULL,
    `created_at`    DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS `user`
(
    `id`           int auto_increment PRIMARY KEY,
    `first_name`   VARCHAR(50)         NOT NULL,
    `last_name`    VARCHAR(255)        NOT NULL,
    `phone`        VARCHAR(255)        NOT NULL,
    `login`        VARCHAR(255) UNIQUE NOT NULL,
    `password`     VARCHAR(255)        NOT NULL,
    `is_confirmed` bool                         DEFAULT 0 NOT NULL,
    `is_deleted`   bool                         DEFAULT 0 NOT NULL,
    `created_at`   DATETIME            NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`   DATETIME            NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS `secret_code`
(
    `id`         int auto_increment PRIMARY KEY,
    `code`       char(4)  NOT NULL,
    `user_id`    int      NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS `photo`
(
    `id`         int auto_increment PRIMARY KEY,
    `name`       varchar(255) NOT NULL,
    `image`      BLOB         NOT NULL,
    `created_at` DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);