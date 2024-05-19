CREATE DATABASE IF NOT EXISTS tasclock;
CREATE TABLE IF NOT EXISTS tasclock.users {
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `name` TEXT NOT NULL,
    `email` TEXT NOT NULL UNIQUE,
    `password` TEXT NOT NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL
};
CREATE TABLE IF NOT EXISTS tasclock.lists {
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `user_id` INTEGER NOT NULL,
    `title` TEXT NOT NULL,
    `description` TEXT,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL
    FOREIGN KEY (`user_id`) references `users`(`id`)
};
CREATE TABLE IF NOT EXISTS tasclock.tasks {
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `user_id` INTEGER NOT NULL,
    `list_id` INTEGER NOT NULL,
    `title` TEXT NOT NULL,
    `description` TEXT,
    `completed` BOOLEAN NOT NULL DEFAULT FALSE,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    FOREIGN KEY (`user_id`) references `users`(`id`),
    FOREIGN KEY (`list_id`) references `lists`(`id`)
};