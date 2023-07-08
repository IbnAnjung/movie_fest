create database if not exists movie_fest;

use movie_fest;


-- movie_fest.users definition

CREATE TABLE `users` (
  `id` binary(16) NOT NULL DEFAULT (uuid_to_bin(uuid())),
  `username` varchar(50) NOT NULL,
  `password` varchar(255) NOT NULL,
  `role` enum('ADMIN','USER') NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- movie_fest.movie_genres definition

CREATE TABLE `movie_genres` (
  `id` binary(16) NOT NULL DEFAULT (uuid_to_bin(uuid())),
  `name` varchar(50) DEFAULT NULL,
  `views_counter` bigint DEFAULT '0',
  `votes_counter` bigint DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- movie_fest.movies definition

CREATE TABLE `movies` (
  `id` binary(16) NOT NULL DEFAULT (uuid_to_bin(uuid())),
  `filename` varchar(255) NOT NULL,
  `views_counter` bigint DEFAULT '0',
  `votes_counter` bigint DEFAULT '0',
  `uploaded_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `uploaded_by` binary(16) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_movies_uploaded_by_users_id` (`uploaded_by`),
  CONSTRAINT `movies_ibfk_1` FOREIGN KEY (`uploaded_by`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- movie_fest.movie_has_genres definition

CREATE TABLE `movie_has_genres` (
  `movie_genre_id` binary(16) NOT NULL,
  `movie_id` binary(16) NOT NULL,
  UNIQUE KEY `unique_movie_genre` (`movie_genre_id`,`movie_id`),
  KEY `fk_movie_has_genres_movie_id_movies_id` (`movie_id`),
  CONSTRAINT `movie_has_genres_ibfk_1` FOREIGN KEY (`movie_genre_id`) REFERENCES `movie_genres` (`id`),
  CONSTRAINT `movie_has_genres_ibfk_2` FOREIGN KEY (`movie_id`) REFERENCES `movies` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;



-- movie_fest.movie_descriptions definition

CREATE TABLE `movie_descriptions` (
  `id` binary(16) NOT NULL DEFAULT (uuid_to_bin(uuid())),
  `movie_id` binary(16) DEFAULT NULL,
  `duration` int NOT NULL,
  `artists` varchar(255) DEFAULT NULL,
  `description` text,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_by` binary(16) NOT NULL,
  `updated_by` binary(16) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_movie_descriptions_movie_id_movies_id` (`movie_id`),
  KEY `fk_movie_descriptions_created_by_users_id` (`created_by`),
  KEY `fk_movie_descriptions_updated_by_users_id` (`updated_by`),
  CONSTRAINT `movie_descriptions_ibfk_1` FOREIGN KEY (`movie_id`) REFERENCES `movies` (`id`),
  CONSTRAINT `movie_descriptions_ibfk_2` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`),
  CONSTRAINT `movie_descriptions_ibfk_3` FOREIGN KEY (`updated_by`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- movie_fest.movie_votes definition

CREATE TABLE `movie_votes` (
  `movie_id` binary(16) NOT NULL,
  `user_id` binary(16) NOT NULL,
  UNIQUE KEY `unique_vote` (`movie_id`,`user_id`),
  KEY `fk_movie_votes_user_id_users_id` (`user_id`),
  CONSTRAINT `movie_votes_ibfk_1` FOREIGN KEY (`movie_id`) REFERENCES `movies` (`id`),
  CONSTRAINT `movie_votes_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


