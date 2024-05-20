CREATE TABLE IF NOT EXISTS tasclock.users (
    `id` INTEGER PRIMARY KEY AUTO_INCREMENT,
    `name` TEXT NOT NULL,
    `email` TEXT NOT NULL,
    `password` TEXT NOT NULL,
    `created_at` TEXT NOT NULL,
    `updated_at` TEXT NOT NULL
) DEFAULT CHARSET=utf8mb4 ENGINE=InnoDB;
CREATE TABLE IF NOT EXISTS tasclock.lists (
    `id` INTEGER PRIMARY KEY AUTO_INCREMENT,
    `user_id` INTEGER NOT NULL,
    `title` TEXT NOT NULL,
    `description` TEXT,
    `created_at` TEXT NOT NULL,
    `updated_at` TEXT NOT NULL,
    FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
) DEFAULT CHARSET=utf8mb4 ENGINE=InnoDB;
CREATE TABLE IF NOT EXISTS tasclock.tasks (
    `id` INTEGER PRIMARY KEY AUTO_INCREMENT,
    `user_id` INTEGER NOT NULL,
    `list_id` INTEGER NOT NULL,
    `title` TEXT NOT NULL,
    `description` TEXT,
    `completed` BOOLEAN NOT NULL DEFAULT FALSE,
    `created_at` TEXT NOT NULL,
    `updated_at` TEXT NOT NULL,
    FOREIGN KEY (`user_id`) REFERENCES `users`(`id`),
    FOREIGN KEY (`list_id`) REFERENCES `lists`(`id`)
) DEFAULT CHARSET=utf8mb4 ENGINE=InnoDB;