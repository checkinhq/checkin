CREATE TABLE users (
    `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `uid` CHAR(27) COLLATE utf8_unicode_ci NOT NULL,

    `email` VARCHAR(255) COLLATE utf8_unicode_ci NOT NULL,
    `password` BINARY(60) NOT NULL,
    `first_name` VARCHAR(255) COLLATE utf8_unicode_ci NOT NULL,
    `last_name` VARCHAR(255) COLLATE utf8_unicode_ci NOT NULL,

    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,

    PRIMARY KEY (`id`),
    UNIQUE KEY (`uid`),
    UNIQUE KEY (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
