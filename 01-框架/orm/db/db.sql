CREATE TABLE `user_info`
(
    `id`          int NOT NULL AUTO_INCREMENT,
    `name`        varchar(255) DEFAULT NULL,
    `age`         int          DEFAULT NULL,
    `birthday`    date         DEFAULT NULL,
    `address`     varchar(255) DEFAULT NULL,
    `create_time` datetime     DEFAULT CURRENT_TIMESTAMP,
    `update_time` datetime     DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;