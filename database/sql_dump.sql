create database if not exists movie_fest;

use movie_fest;


-- movie_fest.users definition

CREATE TABLE `users` (
  `id` bigint unsigned not null AUTO_INCREMENT,
  `username` varchar(50) NOT NULL,
  `password` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- movie_fest.movie_genres definition

CREATE TABLE `movie_genres` (
  `id` bigint unsigned not null AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  `views_counter` bigint DEFAULT '0',
  `votes_counter` bigint DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- movie_fest.movies definition

CREATE TABLE `movies` (
  `id` bigint unsigned not null AUTO_INCREMENT,
  `filename` varchar(255) NOT NULL,
  `title` varchar(100) not null default "unkown",
  `duration` int NOT NULL default 0,
  `artists` varchar(255) NOT NULL DEFAULT "",
  `description` text NOT NULL DEFAULT "",
  `views_counter` bigint DEFAULT 0,
  `votes_counter` bigint DEFAULT 0,
  `uploaded_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- movie_fest.movie_has_genres definition

CREATE TABLE `movie_has_genres` (
  `movie_genre_id` bigint unsigned not null,
  `movie_id` bigint unsigned not null,
  UNIQUE KEY `unique_movie_genre` (`movie_genre_id`,`movie_id`),
  KEY `fk_movie_has_genres_movie_id_movies_id` (`movie_id`),
  CONSTRAINT `movie_has_genres_ibfk_1` FOREIGN KEY (`movie_genre_id`) REFERENCES `movie_genres` (`id`),
  CONSTRAINT `movie_has_genres_ibfk_2` FOREIGN KEY (`movie_id`) REFERENCES `movies` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;



-- movie_fest.movie_descriptions definition

CREATE TABLE `movie_descriptions` (
  `id` bigint unsigned not null AUTO_INCREMENT,
  `movie_id` bigint unsigned not null,
  `title` varchar(100) not null,
  `duration` int NOT NULL,
  `artists` varchar(255) DEFAULT NULL,
  `description` text,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_movie_descriptions_movie_id_movies_id` (`movie_id`),
  CONSTRAINT `movie_descriptions_ibfk_1` FOREIGN KEY (`movie_id`) REFERENCES `movies` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- movie_fest.movie_votes definition

CREATE TABLE `movie_votes` (
  `movie_id` bigint unsigned not null,
  `user_id` bigint unsigned not null,
  UNIQUE KEY `unique_vote` (`movie_id`,`user_id`),
  KEY `fk_movie_votes_user_id_users_id` (`user_id`),
  CONSTRAINT `movie_votes_ibfk_1` FOREIGN KEY (`movie_id`) REFERENCES `movies` (`id`),
  CONSTRAINT `movie_votes_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


