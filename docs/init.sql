SET NAMES utf8;
CREATE DATABASE app_shorturl;
USE app_shorturl;

CREATE TABLE IF NOT EXISTS `shortens` (
  `id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `token` CHAR(8) NOT NULL COMMENT '短网址token',
  `hash` CHAR(32) NOT NULL COMMENT '网址 hash 值',
  `business` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '业务来源(sms,push,...)',
  `long_url` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '重定向的目标网址',
  `visits` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '请求次数',
  `visit_usrs` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '请求用户数',
  `created_at` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
  `last_visited` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '最后请求时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_token` (`token`),
  UNIQUE KEY `unique_hash` (`hash`)
) ENGINE=INNODB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `visit_logs` (
  `id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `token` CHAR(8) NOT NULL COMMENT '短网址id',
  `request_url` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '请求页面网址',
  `user_hash` CHAR(32) NOT NULL DEFAULT '' COMMENT '用户HASH值，用于计算UV',
  `user_agent` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '用户浏览器信息',
  `referer_url` VARCHAR(200) NOT NULL DEFAULT '' COMMENT 'HTTP REFERER信息',
  `client_address` VARCHAR(16) NOT NULL DEFAULT '' COMMENT '用户IP地址',
  `http_method` VARCHAR(20) NOT NULL DEFAULT 'GET' COMMENT 'HTTP请求方式',
  `visit_time` INT(11) NOT NULL DEFAULT '0' COMMENT '访问时间',
  PRIMARY KEY (`id`),
  KEY `index_token` (`token`)
) ENGINE=INNODB DEFAULT CHARSET=utf8;
