create database if not exists movie_fest;

use movie_fest;


-- movie_fest.users definition

CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL,
  `password` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


-- movie_fest.user_tokens definition

CREATE TABLE `user_tokens` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint unsigned NOT NULL,
  `token` text NOT NULL,
  `is_block` tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `user_tokens_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


-- movie_fest.movie_genres definition

CREATE TABLE `movie_genres` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  `views_counter` bigint DEFAULT '0',
  `votes_counter` bigint DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO movie_genres(name) VALUES ("horor"), ("comedy"),  ("drama"), ("action"), ("love");

-- movie_fest.movies definition

CREATE TABLE `movies` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `filename` varchar(255) NOT NULL,
  `title` varchar(100) NOT NULL DEFAULT '',
  `duration` int NOT NULL DEFAULT '0',
  `artists` varchar(255) NOT NULL DEFAULT '',
  `description` text,
  `views_counter` bigint DEFAULT '0',
  `votes_counter` bigint DEFAULT '0',
  `watch_duration_counter` int NOT NULL DEFAULT '0',
  `uploaded_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


-- movie_fest.movie_has_genres definition

CREATE TABLE `movie_has_genres` (
  `movie_genre_id` bigint unsigned NOT NULL,
  `movie_id` bigint unsigned NOT NULL,
  UNIQUE KEY `unique_movie_genre` (`movie_genre_id`,`movie_id`),
  KEY `fk_movie_has_genres_movie_id_movies_id` (`movie_id`),
  CONSTRAINT `movie_has_genres_ibfk_1` FOREIGN KEY (`movie_genre_id`) REFERENCES `movie_genres` (`id`),
  CONSTRAINT `movie_has_genres_ibfk_2` FOREIGN KEY (`movie_id`) REFERENCES `movies` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- movie_fest.movie_descriptions definition

CREATE TABLE `movie_descriptions` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `movie_id` bigint unsigned NOT NULL,
  `title` varchar(100) NOT NULL,
  `duration` int NOT NULL,
  `artists` varchar(255) DEFAULT NULL,
  `description` text,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_movie_descriptions_movie_id_movies_id` (`movie_id`),
  CONSTRAINT `movie_descriptions_ibfk_1` FOREIGN KEY (`movie_id`) REFERENCES `movies` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


-- movie_fest.user_votes definition

CREATE TABLE `user_votes` (
  `user_id` bigint unsigned NOT NULL,
  `movie_id` bigint unsigned NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE KEY `unique_votes` (`user_id`,`movie_id`),
  KEY `user_votes_movie_id_movies_id` (`movie_id`),
  CONSTRAINT `user_votes_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  CONSTRAINT `user_votes_ibfk_2` FOREIGN KEY (`movie_id`) REFERENCES `movies` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


-- movie_fest.user_watches definition

CREATE TABLE `user_watches` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint unsigned NOT NULL,
  `movie_id` bigint unsigned NOT NULL,
  `start_time` bigint NOT NULL DEFAULT '0',
  `end_time` bigint NOT NULL DEFAULT '0',
  `duration` bigint NOT NULL DEFAULT '0',
  `expired_at` timestamp NOT NULL,
  PRIMARY KEY (`id`),
  KEY `user_votes_user_id_users_id` (`user_id`),
  KEY `user_votes_movie_id_movies_id` (`movie_id`),
  CONSTRAINT `user_watches_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  CONSTRAINT `user_watches_ibfk_2` FOREIGN KEY (`movie_id`) REFERENCES `movies` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=45 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;