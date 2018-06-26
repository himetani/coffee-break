DROP TABLE IF EXISTS `coffeetime`.`reservation`;
CREATE TABLE `coffeetime`.`reservation` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `date` varchar(64),
  `name` varchar(64),
  `creation_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `modification_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`date`, `name`),
  UNIQUE KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
