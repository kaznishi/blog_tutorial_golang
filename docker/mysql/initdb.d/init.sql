use defaultdb;

CREATE TABLE `articles` (
  `id` int(11) NOT NULL,
  `title` varchar(255) NOT NULL,
  `content` text NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

ALTER TABLE `articles`
  ADD PRIMARY KEY (`id`);

ALTER TABLE `articles`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

CREATE TABLE `users` (
  `id` int(10) NOT NULL,
  `name` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `salt` varchar(255) NOT NULL,
  `is_active` tinyint(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

ALTER TABLE `users`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT;

INSERT INTO `articles` (`id`, `title`, `content`, `created_at`, `updated_at`) VALUES
(1, 'dummy title', 'dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. dummy content. ', '2018-10-23 00:00:00', '2018-10-23 00:00:00');
