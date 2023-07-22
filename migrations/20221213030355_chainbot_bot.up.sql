CREATE TABLE `cex_order` (
 `id` int unsigned NOT NULL AUTO_INCREMENT,
 `symbol` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
 `price` int NOT NULL,
 `quantity` int NOT NULL,
 `status` int NOT NULL,
 `created_at` datetime DEFAULT NULL,
 `updated_at` datetime DEFAULT NULL,
 `deleted_at` datetime DEFAULT NULL,
 PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;