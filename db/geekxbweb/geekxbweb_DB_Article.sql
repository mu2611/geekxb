CREATE DATABASE  IF NOT EXISTS `geekxbweb` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `geekxbweb`;
-- MySQL dump 10.13  Distrib 5.6.13, for osx10.6 (i386)
--
-- Host: localhost    Database: geekxbweb
-- ------------------------------------------------------
-- Server version	5.6.12

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `DB_Article`
--

DROP TABLE IF EXISTS `DB_Article`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DB_Article`
--

LOCK TABLES `DB_Article` WRITE;
/*!40000 ALTER TABLE `DB_Article` DISABLE KEYS */;
/*!40000 ALTER TABLE `DB_Article` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2013-12-12 15:33:53
