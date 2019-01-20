

CREATE TABLE `list_url` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `page_url` varchar(255) NOT NULL,
  `pid` int(11) DEFAULT NULL,
  `status` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `page_url` (`page_url`) USING HASH
) ENGINE=InnoDB AUTO_INCREMENT=6742 DEFAULT CHARSET=utf8;

CREATE TABLE `page_url` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `page_url` varchar(255) NOT NULL,
  `pid` int(11) DEFAULT NULL,
  `status` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `page_url` (`page_url`) USING HASH
) ENGINE=InnoDB AUTO_INCREMENT=134759 DEFAULT CHARSET=utf8;

CREATE TABLE `img_url` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `page_url` varchar(255) NOT NULL,
  `pid` int(11) DEFAULT NULL,
  `status` tinyint(1) DEFAULT '0',
  `size` int(11) DEFAULT NULL,
  `createtime` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=195764 DEFAULT CHARSET=utf8;