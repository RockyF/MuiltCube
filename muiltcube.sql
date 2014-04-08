/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50527
Source Host           : localhost:3306
Source Database       : muiltcube

Target Server Type    : MYSQL
Target Server Version : 50527
File Encoding         : 65001

Date: 2014-04-09 01:16:41
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for password
-- ----------------------------
DROP TABLE IF EXISTS `password`;
CREATE TABLE `password` (
  `player_id` int(4) NOT NULL,
  `password` varchar(40) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of password
-- ----------------------------
INSERT INTO `password` VALUES ('1001', '123');

-- ----------------------------
-- Table structure for player_info
-- ----------------------------
DROP TABLE IF EXISTS `player_info`;
CREATE TABLE `player_info` (
  `id` int(4) NOT NULL,
  `last_scene` int(4) DEFAULT NULL,
  `last_x` int(4) DEFAULT NULL,
  `last_y` int(4) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of player_info
-- ----------------------------
INSERT INTO `player_info` VALUES ('1001', '101', '1', '2');
INSERT INTO `player_info` VALUES ('1002', '102', '2', '3');

-- ----------------------------
-- Table structure for skin
-- ----------------------------
DROP TABLE IF EXISTS `skin`;
CREATE TABLE `skin` (
  `player_id` int(4) NOT NULL,
  `color_id` int(11) DEFAULT NULL,
  `bubble_id` int(11) DEFAULT NULL,
  `face` int(11) DEFAULT NULL,
  `vertex` int(11) DEFAULT NULL,
  `shadow` int(11) DEFAULT NULL,
  `surround` int(11) DEFAULT NULL,
  PRIMARY KEY (`player_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of skin
-- ----------------------------
INSERT INTO `skin` VALUES ('1001', '1', '1', '0', '0', '0', '0');

-- ----------------------------
-- Table structure for skin_bubble
-- ----------------------------
DROP TABLE IF EXISTS `skin_bubble`;
CREATE TABLE `skin_bubble` (
  `id` int(1) NOT NULL,
  `vlaue` int(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of skin_bubble
-- ----------------------------
INSERT INTO `skin_bubble` VALUES ('1', '1');
INSERT INTO `skin_bubble` VALUES ('2', '2');
INSERT INTO `skin_bubble` VALUES ('3', '3');
INSERT INTO `skin_bubble` VALUES ('4', '4');
INSERT INTO `skin_bubble` VALUES ('5', '5');

-- ----------------------------
-- Table structure for skin_color
-- ----------------------------
DROP TABLE IF EXISTS `skin_color`;
CREATE TABLE `skin_color` (
  `id` int(11) NOT NULL,
  `value` varchar(8) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of skin_color
-- ----------------------------
INSERT INTO `skin_color` VALUES ('1', 'FFFFFF');
INSERT INTO `skin_color` VALUES ('2', 'FF0000');
INSERT INTO `skin_color` VALUES ('3', '00FF00');
INSERT INTO `skin_color` VALUES ('4', '0000FF');
