delimiter $$

CREATE DATABASE `geekxbweb` /*!40100 DEFAULT CHARACTER SET utf8 */$$

delimiter $$

CREATE TABLE `DB_Article` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `arttItle` text NOT NULL,
  `artcontent` text NOT NULL,
  `postname` varchar(45) NOT NULL DEFAULT '' COMMENT '发布人名称',
  `posttime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `readnum` int(11) NOT NULL COMMENT '查看次数',
  `classid` int(11) NOT NULL DEFAULT '0' COMMENT '所属分类id',
  `sortid` int(11) NOT NULL DEFAULT '0',
  `ischeck` int(11) NOT NULL DEFAULT '0' COMMENT '是否审核通过 0 默认 未审核   1已审核',
  `ishot` int(11) NOT NULL DEFAULT '0' COMMENT '是否热门 0 不是   1热门',
  `istop` int(11) NOT NULL DEFAULT '0' COMMENT '是否置顶 0默认  1置顶',
  PRIMARY KEY (`id`),
  KEY `id_idx` (`classid`),
  CONSTRAINT `FK_classid` FOREIGN KEY (`classid`) REFERENCES `DB_ArticleClass` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8$$

delimiter $$

CREATE TABLE `DB_ArticleClass` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `classname` varchar(80) NOT NULL,
  `sortid` int(11) NOT NULL DEFAULT '0',
  `parentid` int(11) NOT NULL DEFAULT '0',
  `forumid` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8$$

delimiter $$

CREATE TABLE `DB_Authority` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Authname` varchar(30) NOT NULL DEFAULT '',
  `Authaction` varchar(100) NOT NULL DEFAULT '',
  `Authgroup` int(11) NOT NULL DEFAULT '1',
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8$$

delimiter $$

CREATE TABLE `DB_User` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL,
  `userpwd` varchar(256) NOT NULL,
  `nickname` varchar(50) NOT NULL DEFAULT '',
  `createtime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `regtime` datetime NOT NULL,
  `usergroupid` int(11) DEFAULT NULL,
  `email` varchar(45) NOT NULL DEFAULT '',
  `headpic` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_idx` (`usergroupid`),
  CONSTRAINT `FK_usergroupid` FOREIGN KEY (`usergroupid`) REFERENCES `DB_UserGroup` (`Id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8$$

delimiter $$

CREATE TABLE `DB_UserGroup` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Ugname` varchar(60) NOT NULL COMMENT '用户组名',
  `Ugmanageauthority` text NOT NULL COMMENT '权限',
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8$$

