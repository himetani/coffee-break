DROP TABLE IF EXISTS `coffeetime`.`reservations`;
CREATE TABLE `coffeetime`.`reservations` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `date` varchar(64) NOT NULL,
  `name` varchar(64) NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`date`, `name`),
  UNIQUE KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
