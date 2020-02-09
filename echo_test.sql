SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `username` varchar(20) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(50) NOT NULL DEFAULT '' COMMENT '密码',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态，1正常，2禁用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='后台用户表';

INSERT INTO `admin` (`id`, `username`, `password`, `status`) VALUES
(1,	'admin',	'e10adc3949ba59abbe56e057f20f883e',	1);

DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '分类名称',
  `weight` int(11) NOT NULL DEFAULT '0' COMMENT '权重（升序）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='资讯分类表';

INSERT INTO `category` (`id`, `name`, `weight`) VALUES
(1,	'最新资讯',	2),
(3,	'官方动态',	3);

DROP TABLE IF EXISTS `news`;
CREATE TABLE `news` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `category_id` int(11) NOT NULL DEFAULT '0' COMMENT '分类id',
  `title` varchar(250) NOT NULL DEFAULT '' COMMENT '标题',
  `description` varchar(500) NOT NULL DEFAULT '' COMMENT '简介',
  `pic` varchar(200) NOT NULL DEFAULT '' COMMENT '图片',
  `content` text NOT NULL COMMENT '内容',
  `weight` int(11) NOT NULL DEFAULT '0' COMMENT '权重（升序）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='资讯表';

INSERT INTO `news` (`id`, `category_id`, `title`, `description`, `pic`, `content`, `weight`) VALUES
(2,	3,	'需要需要需要323',	'需要需要需要323需要需要需要323需要需要需要323需要需要需要323',	'http://localhost:8888/static/upload/2019/December/13/4.jpg',	'<p>需要需要需要323需要需要需要323<img src=\"http://localhost:8888/static/upload/2019/December/13/5.jpg\" style=\"width: 310px;\"><br></p>',	3);

-- 2019-12-13 03:04:24
