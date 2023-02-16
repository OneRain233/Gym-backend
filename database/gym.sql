/*
 Navicat Premium Data Transfer

 Source Server         : 172.250.0.233
 Source Server Type    : MySQL
 Source Server Version : 80031 (8.0.31)
 Source Host           : 172.250.0.233:3306
 Source Schema         : gym

 Target Server Type    : MySQL
 Target Server Version : 80031 (8.0.31)
 File Encoding         : 65001

 Date: 16/02/2023 15:42:29
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for announcement
-- ----------------------------
DROP TABLE IF EXISTS `announcement`;
CREATE TABLE `announcement` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `title` varchar(255) DEFAULT NULL COMMENT 'Title',
  `content` varchar(255) DEFAULT NULL COMMENT 'Content',
  `update_time` datetime DEFAULT NULL COMMENT 'Update time',
  `user_id` int DEFAULT NULL COMMENT 'User ID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='Announcement';

-- ----------------------------
-- Table structure for bank
-- ----------------------------
DROP TABLE IF EXISTS `bank`;
CREATE TABLE `bank` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(255) DEFAULT NULL COMMENT 'Bank name',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=368 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='Bank';

-- ----------------------------
-- Table structure for comment_response
-- ----------------------------
DROP TABLE IF EXISTS `comment_response`;
CREATE TABLE `comment_response` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `parent_response_id` int DEFAULT NULL COMMENT 'PARENT_RESPONSE_ID',
  `title` varchar(255) DEFAULT NULL COMMENT 'title',
  `content` varchar(255) DEFAULT NULL COMMENT 'content',
  `user_id` int DEFAULT NULL COMMENT 'USER_ID',
  `creat_time` datetime DEFAULT NULL COMMENT 'creat_time',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='Comment_response;';

-- ----------------------------
-- Table structure for config
-- ----------------------------
DROP TABLE IF EXISTS `config`;
CREATE TABLE `config` (
  `id` int NOT NULL AUTO_INCREMENT,
  `key` varchar(255) DEFAULT NULL,
  `value` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for course
-- ----------------------------
DROP TABLE IF EXISTS `course`;
CREATE TABLE `course` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `status` int DEFAULT NULL COMMENT 'status',
  `start_time` datetime DEFAULT NULL COMMENT 'start_time',
  `end_time` datetime DEFAULT NULL COMMENT 'end_time',
  `title` varchar(255) DEFAULT NULL COMMENT 'title',
  `description` varchar(255) DEFAULT NULL COMMENT 'course discription',
  `teacher` varchar(255) DEFAULT NULL COMMENT 'teacher',
  `type` int DEFAULT NULL COMMENT 'type',
  `place` varchar(255) DEFAULT NULL COMMENT 'place',
  `max_member_num` int DEFAULT NULL COMMENT 'max member num',
  `current_num` int DEFAULT NULL COMMENT 'now member num',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='Course;';

-- ----------------------------
-- Table structure for credit
-- ----------------------------
DROP TABLE IF EXISTS `credit`;
CREATE TABLE `credit` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` int DEFAULT NULL COMMENT 'User ID',
  `credit` int DEFAULT NULL COMMENT 'Credit of User',
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='Credit ';

-- ----------------------------
-- Table structure for evaluation
-- ----------------------------
DROP TABLE IF EXISTS `evaluation`;
CREATE TABLE `evaluation` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` int DEFAULT NULL COMMENT 'User id',
  `facility_id` int DEFAULT NULL COMMENT 'Facility Id',
  `score` varchar(255) DEFAULT NULL COMMENT 'Score',
  `description` varchar(255) DEFAULT NULL COMMENT 'Description',
  `anonymous` varchar(255) DEFAULT NULL COMMENT 'IS Anonymous',
  `images` varchar(255) DEFAULT NULL COMMENT 'Image urls',
  `videos` varchar(255) DEFAULT NULL COMMENT 'Videos urls',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='Evaluation';

-- ----------------------------
-- Table structure for evaluation_response
-- ----------------------------
DROP TABLE IF EXISTS `evaluation_response`;
CREATE TABLE `evaluation_response` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `parent_response_id` int DEFAULT NULL COMMENT 'reponse_id',
  `user_id` int DEFAULT NULL COMMENT 'response_user_id',
  `evaluation_id` int DEFAULT NULL COMMENT 'evaluation_id',
  `content` varchar(255) DEFAULT NULL COMMENT 'response_content',
  `update_time` datetime DEFAULT NULL COMMENT 'response_time',
  `images` varchar(255) DEFAULT NULL COMMENT 'Image urls',
  `videos` varchar(255) DEFAULT NULL COMMENT 'Videos urls',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='Evaluation Response;';

-- ----------------------------
-- Table structure for facility
-- ----------------------------
DROP TABLE IF EXISTS `facility`;
CREATE TABLE `facility` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'Facility ID',
  `name` varchar(255) DEFAULT NULL COMMENT 'Facility Name',
  `description` varchar(900) DEFAULT NULL COMMENT 'Facility Description',
  `location` varchar(900) DEFAULT NULL COMMENT 'Facility Location',
  `images` varchar(255) DEFAULT NULL COMMENT 'Image urls',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='Facility';

-- ----------------------------
-- Table structure for facility_place
-- ----------------------------
DROP TABLE IF EXISTS `facility_place`;
CREATE TABLE `facility_place` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `facility_id` int DEFAULT NULL COMMENT 'Facility ID',
  `name` varchar(255) DEFAULT NULL COMMENT 'Place Name',
  `description` varchar(900) DEFAULT NULL COMMENT 'Place Description',
  `cost` decimal(24,6) DEFAULT NULL COMMENT 'Place order code',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='Facility Place';

-- ----------------------------
-- Table structure for facility_score
-- ----------------------------
DROP TABLE IF EXISTS `facility_score`;
CREATE TABLE `facility_score` (
  `facility_id` int NOT NULL COMMENT 'Facility ID',
  `score` decimal(24,6) DEFAULT NULL COMMENT 'Current Score',
  PRIMARY KEY (`facility_id`),
  UNIQUE KEY `facility_id` (`facility_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='Facility Score';

-- ----------------------------
-- Table structure for order
-- ----------------------------
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` int DEFAULT NULL COMMENT 'User id',
  `place_id` int DEFAULT NULL COMMENT 'Place ID',
  `time` datetime DEFAULT NULL COMMENT 'Order Time',
  `start_time` datetime DEFAULT NULL COMMENT 'Start time of the order',
  `end_time` datetime DEFAULT NULL COMMENT 'End time of the order',
  `amount` decimal(24,6) DEFAULT NULL COMMENT 'Amount of the order',
  `order_code` varchar(255) DEFAULT NULL COMMENT 'Order code',
  `status` int DEFAULT NULL COMMENT 'Status of the order',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='Order';

-- ----------------------------
-- Table structure for payment
-- ----------------------------
DROP TABLE IF EXISTS `payment`;
CREATE TABLE `payment` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `wallet_id` int NOT NULL COMMENT 'Wallet ID',
  `order_id` int NOT NULL COMMENT 'Order ID',
  `payment_code` varchar(255) NOT NULL COMMENT 'Payment code',
  `time` datetime NOT NULL COMMENT 'Payment Time',
  `amount` decimal(24,6) NOT NULL COMMENT 'Payment Amount',
  `payment_type` int NOT NULL COMMENT 'Payment type',
  `status` int NOT NULL COMMENT 'Payment status',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='Payment';

-- ----------------------------
-- Table structure for receipt
-- ----------------------------
DROP TABLE IF EXISTS `receipt`;
CREATE TABLE `receipt` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `order_id` int DEFAULT NULL COMMENT 'Order Id',
  `receipt_path` varchar(255) DEFAULT NULL COMMENT 'PDF Receipt path',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='Receipt';

-- ----------------------------
-- Table structure for subscription
-- ----------------------------
DROP TABLE IF EXISTS `subscription`;
CREATE TABLE `subscription` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` int NOT NULL,
  `start_time` datetime NOT NULL,
  `END_TIME` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='VIP subscription';

-- ----------------------------
-- Table structure for topic_approves
-- ----------------------------
DROP TABLE IF EXISTS `topic_approves`;
CREATE TABLE `topic_approves` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` int DEFAULT NULL COMMENT 'USER_ID',
  `topic_id` int DEFAULT NULL COMMENT 'TOPIC_ID',
  `type` int DEFAULT NULL COMMENT 'TYPE',
  `update_time` datetime DEFAULT NULL COMMENT 'create_time',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='Topicapproves;';

-- ----------------------------
-- Table structure for topic_comment
-- ----------------------------
DROP TABLE IF EXISTS `topic_comment`;
CREATE TABLE `topic_comment` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` int DEFAULT NULL COMMENT 'USER_ID',
  `topic_id` int DEFAULT NULL COMMENT 'TOPIC_ID',
  `parent_comment_id` int DEFAULT NULL COMMENT 'Parent comment ID',
  `title` varchar(255) DEFAULT NULL COMMENT 'Comment_title',
  `content` varchar(255) DEFAULT NULL COMMENT 'Comment_conrtent',
  `approve` varchar(255) DEFAULT NULL COMMENT 'comment_approve',
  `disapprove` varchar(255) DEFAULT NULL COMMENT 'comment_disapprove',
  `update_time` datetime DEFAULT NULL COMMENT 'creat_time',
  `images` varchar(255) DEFAULT NULL COMMENT 'Image urls',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='Topic_comment;';

-- ----------------------------
-- Table structure for topic_comment_approves
-- ----------------------------
DROP TABLE IF EXISTS `topic_comment_approves`;
CREATE TABLE `topic_comment_approves` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `type` int DEFAULT NULL COMMENT 'TYPE',
  `update_time` datetime DEFAULT NULL COMMENT 'Update time',
  `user_id` int DEFAULT NULL COMMENT 'USER_ID',
  `comment_id` int DEFAULT NULL COMMENT 'topic_comment_ID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='Topiccommentsapproves;';

-- ----------------------------
-- Table structure for topics
-- ----------------------------
DROP TABLE IF EXISTS `topics`;
CREATE TABLE `topics` (
  `id` int NOT NULL COMMENT 'ID',
  `user_id` int DEFAULT NULL COMMENT 'USER_ID',
  `content` varchar(255) DEFAULT NULL COMMENT 'content',
  `title` varchar(255) DEFAULT NULL COMMENT 'title',
  `update_time` datetime DEFAULT NULL COMMENT 'time',
  `approve_count` int DEFAULT NULL COMMENT 'Approve count',
  `disapprove_count` int DEFAULT NULL COMMENT 'Disapprove count',
  `images` varchar(255) DEFAULT NULL COMMENT 'Image urls',
  `videos` varchar(255) DEFAULT NULL COMMENT 'Video urls',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='Topics;';

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'USER ID',
  `username` varchar(90) NOT NULL COMMENT 'USERNAME',
  `password` varchar(255) DEFAULT NULL COMMENT 'PASSWORD',
  `avatar` varchar(255) DEFAULT NULL COMMENT 'AVATAR URL',
  `email` varchar(255) DEFAULT NULL COMMENT 'EMAIL ADDRESS',
  `phone` varchar(255) DEFAULT NULL COMMENT 'PHONE NUMBER',
  `gender` int DEFAULT NULL COMMENT 'GENDER',
  `update_time` datetime DEFAULT NULL COMMENT 'CREATE AT',
  `role` int NOT NULL DEFAULT '0' COMMENT 'ROLE',
  `login_time` datetime DEFAULT NULL COMMENT 'LAST LOGIN TIME',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='User';

-- ----------------------------
-- Table structure for wallet
-- ----------------------------
DROP TABLE IF EXISTS `wallet`;
CREATE TABLE `wallet` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` int DEFAULT NULL COMMENT 'User ID',
  `amount` decimal(24,6) DEFAULT NULL COMMENT 'Money amount',
  `status` int DEFAULT NULL COMMENT 'Wallet status',
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='Wallet';

-- ----------------------------
-- Table structure for wallet_card
-- ----------------------------
DROP TABLE IF EXISTS `wallet_card`;
CREATE TABLE `wallet_card` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `bank_id` int DEFAULT NULL COMMENT 'Bank ID of the card',
  `wallet_id` int DEFAULT NULL COMMENT 'Wallet ID of the card',
  `card_account` varchar(255) DEFAULT NULL COMMENT 'Card number',
  `phone` varchar(255) DEFAULT NULL COMMENT 'Phone of the card',
  `amount` decimal(24,6) DEFAULT NULL COMMENT 'Money amount',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='Wallet Card';

SET FOREIGN_KEY_CHECKS = 1;
