CREATE TABLE `checkins` (
    `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` INT(10) UNSIGNED NOT NULL,
    `date` DATE NOT NULL,

    `previous` TEXT COLLATE utf8_unicode_ci NOT NULL,
    `goals_reached` TINYINT(3) UNSIGNED NOT NULL,
    `next` TEXT COLLATE utf8_unicode_ci NOT NULL,
    `blockers` TEXT COLLATE utf8_unicode_ci NOT NULL,

    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,

    PRIMARY KEY (`id`),
    KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
