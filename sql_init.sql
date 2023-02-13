/*
 Navicat Premium Data Transfer

 Source Server         : dss_test
 Source Server Type    : MySQL
 Source Server Version : 80024
 Source Host           : 192.168.31.100:3306
 Source Schema         : dss_test

 Target Server Type    : MySQL
 Target Server Version : 80024
 File Encoding         : 65001

 Date: 13/11/2022 21:00:03
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for stock_company
-- ----------------------------
DROP TABLE IF EXISTS `stock_company`;
CREATE TABLE `stock_company`  (
  `ts_code` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'TS代码',
  `exchange` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '交易所代码',
  `chairman` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '法人代表',
  `manager` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '总经理',
  `secretary` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '董秘',
  `reg_capital` float(30, 2) NULL DEFAULT NULL COMMENT '注册资本',
  `setup_date` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '注册日期',
  `province` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '所在省份',
  `city` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '所在城市',
  `introduction` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '公司介绍',
  `website` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '公司主页',
  `email` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '电子邮件',
  `office` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '办公室',
  `employees` int(0) NULL DEFAULT NULL COMMENT '员工人数',
  `main_business` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '主要业务及产品',
  `business_scope` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '经营范围',
  PRIMARY KEY (`ts_code`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '上市公司基本信息' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for stock_info
-- ----------------------------
DROP TABLE IF EXISTS `stock_info`;
CREATE TABLE `stock_info`  (
  `ts_code` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'TS代码',
  `symbol` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '股票代码',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '股票名称',
  `area` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '地域',
  `industry` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '所属行业',
  `fullname` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '股票全称',
  `enname` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '英文全称',
  `cnspell` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '拼音缩写',
  `market` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '市场类型（主板/创业板/科创板/CDR）',
  `exchange` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '交易所代码',
  `curr_type` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '交易货币',
  `list_status` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '上市状态 L上市 D退市 P暂停上市',
  `list_date` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '上市日期',
  `delist_date` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '退市日期',
  `is_hs` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '是否沪深港通标的，N否 H沪股通 S深股通',
                               PRIMARY KEY (`ts_code`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '基础数据' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for task_info
-- ----------------------------
DROP TABLE IF EXISTS `task_info`;
CREATE TABLE `task_info`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `task_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '任务名称',
  `date` varchar(8) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '日期',
  `spend_time` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci  NOT NULL COMMENT '花费时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '定时任务执行记录' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for ths_gn
-- ----------------------------
DROP TABLE IF EXISTS `ths_gn`;
CREATE TABLE `ths_gn`  (
  `code` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '同花顺概念代码',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '同花顺概念名称',
  PRIMARY KEY (`code`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '同花顺概念' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for ths_gn_rel_symbol
-- ----------------------------
DROP TABLE IF EXISTS `ths_gn_rel_symbol`;
CREATE TABLE `ths_gn_rel_symbol`  (
  `gn_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '同花顺概念名称',
  `symbol` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '股票代码',
  PRIMARY KEY (`gn_name`, `symbol`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '同花顺概念与股票代码关联关系' ROW_FORMAT = Dynamic;


DROP TABLE IF EXISTS `long_hu`;
CREATE TABLE `long_hu`  (
  `id` varchar(36) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'id',
  `type` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '类型',
  `symbol` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '股票代码',
  `trade_date` varchar(8) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '交易日期',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '股票名称',
  `close` float(30, 2) NULL DEFAULT NULL COMMENT '收盘价',
  `pct_chg` float(30, 2) NULL DEFAULT NULL COMMENT '涨跌幅',
  `amount` float(30, 2) NULL DEFAULT NULL COMMENT '成交额',
  `net_worth` float(30, 2) NULL DEFAULT NULL COMMENT '净买入额',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `symbol_index`(`symbol`) USING BTREE,
  INDEX `trade_date_index`(`trade_date`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for long_hu_detail
-- ----------------------------
DROP TABLE IF EXISTS `long_hu_detail`;
CREATE TABLE `long_hu_detail`  (
  `long_hu_id` varchar(36) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '龙虎榜id',
  `dept` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '营业部',
  `label` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '标签',
  `buy` float(11, 2) NULL DEFAULT NULL COMMENT '买入额',
  `sell` float(11, 2) NULL DEFAULT NULL COMMENT '卖出额',
  `net_worth` float(11, 2) NULL DEFAULT NULL COMMENT '净买入额',
  INDEX `long_hu_id_index`(`long_hu_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '龙虎榜详情' ROW_FORMAT = Dynamic;

-- -- ----------------------------
-- -- Table structure for ths_hy
-- -- ----------------------------
-- DROP TABLE IF EXISTS `ths_hy`;
-- CREATE TABLE `ths_hy`  (
--   `code` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '同花顺行业代码',
--   `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '同花顺行业名称',
--   PRIMARY KEY (`code`) USING BTREE
-- ) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '同花顺行业' ROW_FORMAT = Dynamic;

-- -- ----------------------------
-- -- Table structure for ths_hy_rel_symbol
-- -- ----------------------------
-- DROP TABLE IF EXISTS `ths_hy_rel_symbol`;
-- CREATE TABLE `ths_hy_rel_symbol`  (
--   `hy_code` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '同花顺行业代码',
--   `symbol` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '股票代码',
--   PRIMARY KEY (`hy_code`, `symbol`) USING BTREE
-- ) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '同花顺行业与股票代码关联关系' ROW_FORMAT = Dynamic;

-- -- ----------------------------
-- -- Table structure for ths_hy_quote
-- -- ----------------------------
-- DROP TABLE IF EXISTS `ths_hy_quote`;
-- CREATE TABLE `ths_hy_quote`  (
--   `code` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '行业代码',
--   `trade_date` varchar(8) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '交易日期',
--   `open` float(30, 4) NULL DEFAULT NULL COMMENT '开盘价',
--   `close` float(30, 4) NULL DEFAULT NULL COMMENT '收盘价',
--   `low` float(30, 4) NULL DEFAULT NULL COMMENT '最低价',
--   `high` float(30, 4) NULL DEFAULT NULL COMMENT '最高价',
--   `pre_close` float(30, 4) NULL DEFAULT NULL COMMENT '昨收价',
--   `change` float(30, 4) NULL DEFAULT NULL COMMENT '资金流入(亿)',
--   `pct_chg` float(30, 4) NULL DEFAULT NULL COMMENT '涨跌幅',
--   `vol` float(30, 4) NULL DEFAULT NULL COMMENT '成交量(万手)',
--   `amount` float(30, 4) NULL DEFAULT NULL COMMENT '成交额(亿)',
--   `rank` int(0) NULL DEFAULT NULL COMMENT '涨幅排名',
--   `rise_count` int(0) NULL DEFAULT NULL COMMENT '上涨家数',
--   `fall_count` int(0) NULL DEFAULT NULL COMMENT '下跌家数',
--   PRIMARY KEY (`code`, `trade_date`) USING BTREE
-- ) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '同花顺行业行情' ROW_FORMAT = Dynamic;

-- -- ----------------------------
-- -- Table structure for ths_gn_quote
-- -- ----------------------------
-- DROP TABLE IF EXISTS `ths_gn_quote`;
-- CREATE TABLE `ths_gn_quote`  (
--   `code` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '概念代码',
--   `trade_date` varchar(8) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '交易日期',
--   `open` float(30, 4) NULL DEFAULT NULL COMMENT '开盘价',
--   `close` float(30, 4) NULL DEFAULT NULL COMMENT '收盘价',
--   `low` float(30, 4) NULL DEFAULT NULL COMMENT '最低价',
--   `high` float(30, 4) NULL DEFAULT NULL COMMENT '最高价',
--   `pre_close` float(30, 4) NULL DEFAULT NULL COMMENT '昨收价',
--   `change` float(30, 4) NULL DEFAULT NULL COMMENT '资金流入(亿)',
--   `pct_chg` float(30, 4) NULL DEFAULT NULL COMMENT '涨跌幅',
--   `vol` float(30, 4) NULL DEFAULT NULL COMMENT '成交量(万手)',
--   `amount` float(30, 4) NULL DEFAULT NULL COMMENT '成交额(亿)',
--   `rank` int(0) NULL DEFAULT NULL COMMENT '涨幅排名',
--   `rise_count` int(0) NULL DEFAULT NULL COMMENT '上涨家数',
--   `fall_count` int(0) NULL DEFAULT NULL COMMENT '下跌家数',
--   PRIMARY KEY (`code`, `trade_date`) USING BTREE
-- ) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '同花顺概念行情' ROW_FORMAT = Dynamic;

-- SET FOREIGN_KEY_CHECKS = 1;
