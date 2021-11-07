/*
Navicat MySQL Data Transfer

Source Server         : tale
Source Server Version : 80021
Source Host           : localhost:3306
Source Database       : app_upgrade

Target Server Type    : MYSQL
Target Server Version : 80021
File Encoding         : 65001

Date: 2021-11-07 21:10:48
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for rules
-- ----------------------------
DROP TABLE IF EXISTS `rules`;
CREATE TABLE `rules` (
  `rid` int NOT NULL COMMENT '规则唯一标识',
  `state` tinyint(1) DEFAULT NULL COMMENT '规则是否可用',
  `aid` int DEFAULT NULL COMMENT 'app唯一标识',
  `platform` varchar(25) DEFAULT NULL,
  `channel` varchar(25) DEFAULT NULL,
  `device_id_list` text,
  `min_update_version_code` varchar(25) DEFAULT NULL,
  `max_update_version_code` varchar(25) DEFAULT NULL,
  `min_os_api` int DEFAULT NULL,
  `max_os_api` int DEFAULT NULL,
  `cpu_arch` varchar(25) DEFAULT NULL,
  `download_url` varchar(255) DEFAULT NULL,
  `update_version_code` varchar(255) DEFAULT NULL,
  `md5` varchar(255) DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `update_tips` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`rid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of rules
-- ----------------------------
INSERT INTO `rules` VALUES ('1', '0', '0', '', '', 'hello;uu;kk', '', '', '0', '0', '', '', '', '', '', '');
