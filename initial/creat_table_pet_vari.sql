/*
Navicat MySQL Data Transfer

Source Server         : hello
Source Server Version : 50720
Source Host           : localhost:3306
Source Database       : actestdb

Target Server Type    : MYSQL
Target Server Version : 50720
File Encoding         : 65001

Date: 2018-05-19 22:11:42
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for pet_vari
-- ----------------------------
DROP TABLE IF EXISTS `pet_vari`;
CREATE TABLE `pet_vari` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `pet_speci_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=37 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of pet_vari
-- ----------------------------
INSERT INTO `pet_vari` VALUES ('1', '暹罗猫', '1');
INSERT INTO `pet_vari` VALUES ('2', '布偶猫', '1');
INSERT INTO `pet_vari` VALUES ('3', '苏格兰折耳猫', '1');
INSERT INTO `pet_vari` VALUES ('4', '波斯猫', '1');
INSERT INTO `pet_vari` VALUES ('5', '英国短毛猫', '1');
INSERT INTO `pet_vari` VALUES ('6', '缅因猫', '1');
INSERT INTO `pet_vari` VALUES ('7', '美国短尾猫', '1');
INSERT INTO `pet_vari` VALUES ('8', '埃及猫', '1');
INSERT INTO `pet_vari` VALUES ('9', '橘猫', '1');
INSERT INTO `pet_vari` VALUES ('10', '哈士奇', '2');
INSERT INTO `pet_vari` VALUES ('11', '贵宾犬', '2');
INSERT INTO `pet_vari` VALUES ('12', '松狮', '2');
INSERT INTO `pet_vari` VALUES ('13', '博美', '2');
INSERT INTO `pet_vari` VALUES ('14', '萨摩耶', '2');
INSERT INTO `pet_vari` VALUES ('15', '柯基犬', '2');
INSERT INTO `pet_vari` VALUES ('16', '秋田犬', '2');
INSERT INTO `pet_vari` VALUES ('17', '泰迪犬', '2');
INSERT INTO `pet_vari` VALUES ('18', '金毛', '2');
INSERT INTO `pet_vari` VALUES ('19', '荷兰兔', '3');
INSERT INTO `pet_vari` VALUES ('20', '安哥拉兔', '3');
INSERT INTO `pet_vari` VALUES ('21', '海棠兔', '3');
INSERT INTO `pet_vari` VALUES ('22', '暹罗兔', '3');
INSERT INTO `pet_vari` VALUES ('23', '云雀', '4');
INSERT INTO `pet_vari` VALUES ('24', '金丝雀', '4');
INSERT INTO `pet_vari` VALUES ('25', '黄鹂', '4');
INSERT INTO `pet_vari` VALUES ('26', '鹦鹉', '4');
INSERT INTO `pet_vari` VALUES ('27', '仓鼠', '5');
INSERT INTO `pet_vari` VALUES ('28', '龙猫', '5');
INSERT INTO `pet_vari` VALUES ('29', '荷兰猪', '5');
INSERT INTO `pet_vari` VALUES ('30', '观赏鱼', '6');
INSERT INTO `pet_vari` VALUES ('31', '金鱼', '6');
INSERT INTO `pet_vari` VALUES ('32', '龙虾', '6');
INSERT INTO `pet_vari` VALUES ('33', '乌龟', '7');
INSERT INTO `pet_vari` VALUES ('34', '蜥蜴', '7');
INSERT INTO `pet_vari` VALUES ('35', '蛇', '7');
INSERT INTO `pet_vari` VALUES ('36', '蛙', '7');
