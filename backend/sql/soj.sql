/*
 Navicat Premium Data Transfer

 Source Server         : wsl-ubuntu22.04
 Source Server Type    : MySQL
 Source Server Version : 80032 (8.0.32)
 Source Host           : localhost:3306
 Source Schema         : soj

 Target Server Type    : MySQL
 Target Server Version : 80032 (8.0.32)
 File Encoding         : 65001

 Date: 11/03/2023 14:43:11
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for problem
-- ----------------------------
DROP TABLE IF EXISTS `problem`;
CREATE TABLE `problem`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `create_time` bigint UNSIGNED NOT NULL,
  `update_time` bigint NOT NULL,
  `delete_time` bigint UNSIGNED NOT NULL DEFAULT 0,
  `problem_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '问题的自定义ID 例如（HOJ-1000）',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '题目标题',
  `author` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '未知' COMMENT '作者',
  `type` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '0为ACM,1为OI',
  `time_limit` int NULL DEFAULT 1000 COMMENT '单位ms',
  `memory_limit` int NULL DEFAULT 65535 COMMENT '单位kb',
  `stack_limit` int NULL DEFAULT 128 COMMENT '单位mb',
  `description` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '描述',
  `input` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '输入描述',
  `output` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '输出描述',
  `examples` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '题面样例',
  `source` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '题目来源',
  `difficulty` tinyint NULL DEFAULT 0 COMMENT '题目难度,0简单，1中等，2困难',
  `hint` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '备注,提醒',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `problem_id`(`problem_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `create_time` bigint UNSIGNED NOT NULL COMMENT '创建时间',
  `update_time` bigint UNSIGNED NOT NULL COMMENT '更新时间',
  `delete_time` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
  `username` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `password` varchar(97) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户密码',
  `avatar` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'https://cdn.acwing.com/media/user/profile/photo/71847_lg_463b89cdb9.jpeg' COMMENT '头像',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uindex_username`(`username` ASC) USING BTREE,
  INDEX `id`(`id` ASC, `username` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户信息表' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
