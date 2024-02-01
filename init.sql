/*
 Navicat Premium Data Transfer

 Source Server         : dss
 Source Server Type    : MySQL
 Source Server Version : 80200
 Source Host           : debian:3306
 Source Schema         : dss

 Target Server Type    : MySQL
 Target Server Version : 80200
 File Encoding         : 65001

 Date: 01/02/2024 18:01:07
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for bk
-- ----------------------------
DROP TABLE IF EXISTS `bk`;
CREATE TABLE `bk`  (
  `code` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '板块代码',
  `name` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '板块名称',
  `type` int(0) NULL DEFAULT NULL COMMENT '板块类型 1:地域;2:行业;3:概念',
  PRIMARY KEY (`code`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '同花顺概念' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for bk_quote
-- ----------------------------
DROP TABLE IF EXISTS `bk_quote`;
CREATE TABLE `bk_quote`  (
  `bk_code` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '板块代码',
  `trade_date` varchar(8) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '交易日期',
  `close` float(30, 4) NULL DEFAULT NULL COMMENT '收盘价',
  `change` float(30, 4) NULL DEFAULT NULL COMMENT '涨跌额',
  `pct_chg` float(30, 4) NULL DEFAULT NULL COMMENT '涨跌幅',
  `total` float(30, 4) NULL DEFAULT NULL COMMENT '总市值',
  `rate` float(30, 4) NULL DEFAULT NULL COMMENT '换手率',
  `rank` int(0) NULL DEFAULT NULL COMMENT '涨幅排名',
  `rise_count` int(0) NULL DEFAULT NULL COMMENT '上涨家数',
  `fall_count` int(0) NULL DEFAULT NULL COMMENT '下跌家数',
  `lead` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '领涨股',
  `lead_pct_chg` float(30, 4) NULL DEFAULT NULL COMMENT '领涨股涨跌幅',
  `open` float(30, 4) NULL DEFAULT NULL COMMENT '开盘价',
  `high` float(30, 4) NULL DEFAULT NULL COMMENT '最高价',
  `low` float(30, 4) NULL DEFAULT NULL COMMENT '最低价',
  `pre_close` float(30, 4) NULL DEFAULT NULL COMMENT '昨收价',
  `vol` float(30, 4) NULL DEFAULT NULL COMMENT '成交量(万手)',
  `amount` float(30, 4) NULL DEFAULT NULL COMMENT '成交额(亿)',
  PRIMARY KEY (`bk_code`, `trade_date`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '板块行情' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for bk_rel_symbol
-- ----------------------------
DROP TABLE IF EXISTS `bk_rel_symbol`;
CREATE TABLE `bk_rel_symbol`  (
  `bk_code` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '板块编码',
  `symbol` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '股票代码',
  PRIMARY KEY (`bk_code`, `symbol`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '同花顺概念与股票代码关联关系' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for long_hu
-- ----------------------------
DROP TABLE IF EXISTS `long_hu`;
CREATE TABLE `long_hu`  (
  `id` varchar(36) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT 'id',
  `type` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '类型',
  `symbol` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '股票代码',
  `trade_date` varchar(8) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '交易日期',
  `name` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '股票名称',
  `close` float(30, 2) NULL DEFAULT NULL COMMENT '收盘价',
  `pct_chg` float(30, 2) NULL DEFAULT NULL COMMENT '涨跌幅',
  `amount` float(30, 2) NULL DEFAULT NULL COMMENT '成交额',
  `net_worth` float(30, 2) NULL DEFAULT NULL COMMENT '净买入额',
  `volume` float(30, 2) NULL DEFAULT NULL COMMENT '成交量',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for long_hu_detail
-- ----------------------------
DROP TABLE IF EXISTS `long_hu_detail`;
CREATE TABLE `long_hu_detail`  (
  `long_hu_id` varchar(36) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '龙虎榜id',
  `dept` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '营业部',
  `buy` float(11, 2) NULL DEFAULT NULL COMMENT '买入额',
  `sell` float(11, 2) NULL DEFAULT NULL COMMENT '卖出额',
  `net_worth` float(11, 2) NULL DEFAULT NULL COMMENT '净买入额',
  `ratio` float(11, 2) NULL DEFAULT NULL COMMENT '占比'
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '龙虎榜详情' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for stock_info
-- ----------------------------
DROP TABLE IF EXISTS `stock_info`;
CREATE TABLE `stock_info`  (
  `ts_code` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT 'TS代码',
  `symbol` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '股票代码',
  `name` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '股票名称',
  `area` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '地域',
  `industry` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '所属行业',
  `fullname` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '股票全称',
  `enname` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '英文全称',
  `cnspell` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '拼音缩写',
  `market` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '市场类型（主板/创业板/科创板/CDR）',
  `exchange` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '交易所代码',
  `curr_type` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '交易货币',
  `list_status` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '上市状态 L上市 D退市 P暂停上市',
  `list_date` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '上市日期',
  `delist_date` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '退市日期',
  `is_hs` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '是否沪深港通标的，N否 H沪股通 S深股通',
  PRIMARY KEY (`ts_code`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '基础数据' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for stock_quote_2010
-- ----------------------------
DROP TABLE IF EXISTS `stock_quote_2010`;
CREATE TABLE `stock_quote_2010`  (
  `ts_code` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT 'TS代码',
  `trade_date` varchar(8) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '交易日期',
  `open` float(30, 2) NULL DEFAULT NULL COMMENT '开盘价',
  `high` float(30, 2) NULL DEFAULT NULL COMMENT '最高价',
  `low` float(30, 2) NULL DEFAULT NULL COMMENT '最低价',
  `close` float(30, 2) NULL DEFAULT NULL COMMENT '收盘价',
  `pre_close` float(30, 2) NULL DEFAULT NULL COMMENT '昨收价(前复权)',
  `change` float(30, 2) NULL DEFAULT NULL COMMENT '涨跌额',
  `pct_chg` float(30, 2) NULL DEFAULT NULL COMMENT '涨跌幅(未复权)',
  `vol` float(30, 2) NULL DEFAULT NULL COMMENT '成交量(万手)',
  `amount` float(30, 2) NULL DEFAULT NULL COMMENT '成交额(万元)',
  `limit_up` tinyint(1) NULL DEFAULT NULL COMMENT '涨停板',
  PRIMARY KEY (`ts_code`, `trade_date`) USING BTREE,
  INDEX `trade_date_index`(`trade_date`) USING BTREE,
  INDEX `ts_code_index`(`ts_code`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '股票行情' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for stock_quote_2011
-- ----------------------------
DROP TABLE IF EXISTS `stock_quote_2011`;
CREATE TABLE `stock_quote_2011`  (
  `ts_code` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT 'TS代码',
  `trade_date` varchar(8) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '交易日期',
  `open` float(30, 2) NULL DEFAULT NULL COMMENT '开盘价',
  `high` float(30, 2) NULL DEFAULT NULL COMMENT '最高价',
  `low` float(30, 2) NULL DEFAULT NULL COMMENT '最低价',
  `close` float(30, 2) NULL DEFAULT NULL COMMENT '收盘价',
  `pre_close` float(30, 2) NULL DEFAULT NULL COMMENT '昨收价(前复权)',
  `change` float(30, 2) NULL DEFAULT NULL COMMENT '涨跌额',
  `pct_chg` float(30, 2) NULL DEFAULT NULL COMMENT '涨跌幅(未复权)',
  `vol` float(30, 2) NULL DEFAULT NULL COMMENT '成交量(万手)',
  `amount` float(30, 2) NULL DEFAULT NULL COMMENT '成交额(万元)',
  `limit_up` tinyint(1) NULL DEFAULT NULL COMMENT '涨停板',
  PRIMARY KEY (`ts_code`, `trade_date`) USING BTREE,
  INDEX `trade_date_index`(`trade_date`) USING BTREE,
  INDEX `ts_code_index`(`ts_code`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '股票行情' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for stock_quote_2012
-- ----------------------------
DROP TABLE IF EXISTS `stock_quote_2012`;
CREATE TABLE `stock_quote_2012`  (
  `ts_code` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT 'TS代码',
  `trade_date` varchar(8) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '交易日期',
  `open` float(30, 2) NULL DEFAULT NULL COMMENT '开盘价',
  `high` float(30, 2) NULL DEFAULT NULL COMMENT '最高价',
  `low` float(30, 2) NULL DEFAULT NULL COMMENT '最低价',
  `close` float(30, 2) NULL DEFAULT NULL COMMENT '收盘价',
  `pre_close` float(30, 2) NULL DEFAULT NULL COMMENT '昨收价(前复权)',
  `change` float(30, 2) NULL DEFAULT NULL COMMENT '涨跌额',
  `pct_chg` float(30, 2) NULL DEFAULT NULL COMMENT '涨跌幅(未复权)',
  `vol` float(30, 2) NULL DEFAULT NULL COMMENT '成交量(万手)',
  `amount` float(30, 2) NULL DEFAULT NULL COMMENT '成交额(万元)',
  `limit_up` tinyint(1) NULL DEFAULT NULL COMMENT '涨停板',
  PRIMARY KEY (`ts_code`, `trade_date`) USING BTREE,
  INDEX `trade_date_index`(`trade_date`) USING BTREE,
  INDEX `ts_code_index`(`ts_code`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '股票行情' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for stock_quote_2013
-- ----------------------------
DROP TABLE IF EXISTS `stock_quote_2013`;
CREATE TABLE `stock_quote_2013`  (
  `ts_code` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT 'TS代码',
  `trade_date` varchar(8) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '交易日期',
  `open` float(30, 2) NULL DEFAULT NULL COMMENT '开盘价',
  `high` float(30, 2) NULL DEFAULT NULL COMMENT '最高价',
  `low` float(30, 2) NULL DEFAULT NULL COMMENT '最低价',
  `close` float(30, 2) NULL DEFAULT NULL COMMENT '收盘价',
  `pre_close` float(30, 2) NULL DEFAULT NULL COMMENT '昨收价(前复权)',
  `change` float(30, 2) NULL DEFAULT NULL COMMENT '涨跌额',
  `pct_chg` float(30, 2) NULL DEFAULT NULL COMMENT '涨跌幅(未复权)',
  `vol` float(30, 2) NULL DEFAULT NULL COMMENT '成交量(万手)',
  `amount` float(30, 2) NULL DEFAULT NULL COMMENT '成交额(万元)',
  `limit_up` tinyint(1) NULL DEFAULT NULL COMMENT '涨停板',
  PRIMARY KEY (`ts_code`, `trade_date`) USING BTREE,
  INDEX `trade_date_index`(`trade_date`) USING BTREE,
  INDEX `ts_code_index`(`ts_code`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '股票行情' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for stock_quote_2014
-- ----------------------------
DROP TABLE IF EXISTS `stock_quote_2014`;
CREATE TABLE `stock_quote_2014`  (
  `ts_code` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT 'TS代码',
  `trade_date` varchar(8) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '交易日期',
  `open` float(30, 2) NULL DEFAULT NULL COMMENT '开盘价',
  `high` float(30, 2) NULL DEFAULT NULL COMMENT '最高价',
  `low` float(30, 2) NULL DEFAULT NULL COMMENT '最低价',
  `close` float(30, 2) NULL DEFAULT NULL COMMENT '收盘价',
  `pre_close` float(30, 2) NULL DEFAULT NULL COMMENT '昨收价(前复权)',
  `change` float(30, 2) NULL DEFAULT NULL COMMENT '涨跌额',
  `pct_chg` float(30, 2) NULL DEFAULT NULL COMMENT '涨跌幅(未复权)',
  `vol` float(30, 2) NULL DEFAULT NULL COMMENT '成交量(万手)',
  `amount` float(30, 2) NULL DEFAULT NULL COMMENT '成交额(万元)',
  `limit_up` tinyint(1) NULL DEFAULT NULL COMMENT '涨停板',
  PRIMARY KEY (`ts_code`, `trade_date`) USING BTREE,
  INDEX `trade_date_index`(`trade_date`) USING BTREE,
  INDEX `ts_code_index`(`ts_code`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '股票行情' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for stock_quote_2015
-- ----------------------------
DROP TABLE IF EXISTS `stock_quote_2015`;
CREATE TABLE `stock_quote_2015`  (
  `ts_code` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT 'TS代码',
  `trade_date` varchar(8) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '交易日期',
  `open` float(30, 2) NULL DEFAULT NULL COMMENT '开盘价',
  `high` float(30, 2) NULL DEFAULT NULL COMMENT '最高价',
  `low` float(30, 2) NULL DEFAULT NULL COMMENT '最低价',
  `close` float(30, 2) NULL DEFAULT NULL COMMENT '收盘价',
  `pre_close` float(30, 2) NULL DEFAULT NULL COMMENT '昨收价(前复权)',
  `change` float(30, 2) NULL DEFAULT NULL COMMENT '涨跌额',
  `pct_chg` float(30, 2) NULL DEFAULT NULL COMMENT '涨跌幅(未复权)',
  `vol` float(30, 2) NULL DEFAULT NULL COMMENT '成交量(万手)',
  `amount` float(30, 2) NULL DEFAULT NULL COMMENT '成交额(万元)',
  `limit_up` tinyint(1) NULL DEFAULT NULL COMMENT '涨停板',
  PRIMARY KEY (`ts_code`, `trade_date`) USING BTREE,
  INDEX `trade_date_index`(`trade_date`) USING BTREE,
  INDEX `ts_code_index`(`ts_code`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '股票行情' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for stock_quote_2016
-- ----------------------------
DROP TABLE IF EXISTS `stock_quote_2016`;
CREATE TABLE `stock_quote_2016`  (
  `ts_code` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT 'TS代码',
  `trade_date` varchar(8) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '交易日期',
  `open` float(30, 2) NULL DEFAULT NULL COMMENT '开盘价',
  `high` float(30, 2) NULL DEFAULT NULL COMMENT '最高价',
  `low` float(30, 2) NULL DEFAULT NULL COMMENT '最低价',
  `close` float(30, 2) NULL DEFAULT NULL COMMENT '收盘价',
  `pre_close` float(30, 2) NULL DEFAULT NULL COMMENT '昨收价(前复权)',
  `change` float(30, 2) NULL DEFAULT NULL COMMENT '涨跌额',
  `pct_chg` float(30, 2) NULL DEFAULT NULL COMMENT '涨跌幅(未复权)',
  `vol` float(30, 2) NULL DEFAULT NULL COMMENT '成交量(万手)',
  `amount` float(30, 2) NULL DEFAULT NULL COMMENT '成交额(万元)',
  `limit_up` tinyint(1) NULL DEFAULT NULL COMMENT '涨停板',
  PRIMARY KEY (`ts_code`, `trade_date`) USING BTREE,
  INDEX `trade_date_index`(`trade_date`) USING BTREE,
  INDEX `ts_code_index`(`ts_code`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '股票行情' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for stock_quote_2017
-- ----------------------------
DROP TABLE IF EXISTS `stock_quote_2017`;
CREATE TABLE `stock_quote_2017`  (
  `ts_code` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT 'TS代码',
  `trade_date` varchar(8) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '交易日期',
  `open` float(30, 2) NULL DEFAULT NULL COMMENT '开盘价',
  `high` float(30, 2) NULL DEFAULT NULL COMMENT '最高价',
  `low` float(30, 2) NULL DEFAULT NULL COMMENT '最低价',
  `close` float(30, 2) NULL DEFAULT NULL COMMENT '收盘价',
  `pre_close` float(30, 2) NULL DEFAULT NULL COMMENT '昨收价(前复权)',
  `change` float(30, 2) NULL DEFAULT NULL COMMENT '涨跌额',
  `pct_chg` float(30, 2) NULL DEFAULT NULL COMMENT '涨跌幅(未复权)',
  `vol` float(30, 2) NULL DEFAULT NULL COMMENT '成交量(万手)',
  `amount` float(30, 2) NULL DEFAULT NULL COMMENT '成交额(万元)',
  `limit_up` tinyint(1) NULL DEFAULT NULL COMMENT '涨停板',
  PRIMARY KEY (`ts_code`, `trade_date`) USING BTREE,
  INDEX `trade_date_index`(`trade_date`) USING BTREE,
  INDEX `ts_code_index`(`ts_code`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '股票行情' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for stock_quote_2018
-- ----------------------------
DROP TABLE IF EXISTS `stock_quote_2018`;
CREATE TABLE `stock_quote_2018`  (
  `ts_code` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT 'TS代码',
  `trade_date` varchar(8) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '交易日期',
  `open` float(30, 2) NULL DEFAULT NULL COMMENT '开盘价',
  `high` float(30, 2) NULL DEFAULT NULL COMMENT '最高价',
  `low` float(30, 2) NULL DEFAULT NULL COMMENT '最低价',
  `close` float(30, 2) NULL DEFAULT NULL COMMENT '收盘价',
  `pre_close` float(30, 2) NULL DEFAULT NULL COMMENT '昨收价(前复权)',
  `change` float(30, 2) NULL DEFAULT NULL COMMENT '涨跌额',
  `pct_chg` float(30, 2) NULL DEFAULT NULL COMMENT '涨跌幅(未复权)',
  `vol` float(30, 2) NULL DEFAULT NULL COMMENT '成交量(万手)',
  `amount` float(30, 2) NULL DEFAULT NULL COMMENT '成交额(万元)',
  `limit_up` tinyint(1) NULL DEFAULT NULL COMMENT '涨停板',
  PRIMARY KEY (`ts_code`, `trade_date`) USING BTREE,
  INDEX `trade_date_index`(`trade_date`) USING BTREE,
  INDEX `ts_code_index`(`ts_code`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '股票行情' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for stock_quote_2019
-- ----------------------------
DROP TABLE IF EXISTS `stock_quote_2019`;
CREATE TABLE `stock_quote_2019`  (
  `ts_code` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT 'TS代码',
  `trade_date` varchar(8) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '交易日期',
  `open` float(30, 2) NULL DEFAULT NULL COMMENT '开盘价',
  `high` float(30, 2) NULL DEFAULT NULL COMMENT '最高价',
  `low` float(30, 2) NULL DEFAULT NULL COMMENT '最低价',
  `close` float(30, 2) NULL DEFAULT NULL COMMENT '收盘价',
  `pre_close` float(30, 2) NULL DEFAULT NULL COMMENT '昨收价(前复权)',
  `change` float(30, 2) NULL DEFAULT NULL COMMENT '涨跌额',
  `pct_chg` float(30, 2) NULL DEFAULT NULL COMMENT '涨跌幅(未复权)',
  `vol` float(30, 2) NULL DEFAULT NULL COMMENT '成交量(万手)',
  `amount` float(30, 2) NULL DEFAULT NULL COMMENT '成交额(万元)',
  `limit_up` tinyint(1) NULL DEFAULT NULL COMMENT '涨停板',
  PRIMARY KEY (`ts_code`, `trade_date`) USING BTREE,
  INDEX `trade_date_index`(`trade_date`) USING BTREE,
  INDEX `ts_code_index`(`ts_code`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '股票行情' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for stock_quote_2020
-- ----------------------------
DROP TABLE IF EXISTS `stock_quote_2020`;
CREATE TABLE `stock_quote_2020`  (
  `ts_code` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT 'TS代码',
  `trade_date` varchar(8) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '交易日期',
  `open` float(30, 2) NULL DEFAULT NULL COMMENT '开盘价',
  `high` float(30, 2) NULL DEFAULT NULL COMMENT '最高价',
  `low` float(30, 2) NULL DEFAULT NULL COMMENT '最低价',
  `close` float(30, 2) NULL DEFAULT NULL COMMENT '收盘价',
  `pre_close` float(30, 2) NULL DEFAULT NULL COMMENT '昨收价(前复权)',
  `change` float(30, 2) NULL DEFAULT NULL COMMENT '涨跌额',
  `pct_chg` float(30, 2) NULL DEFAULT NULL COMMENT '涨跌幅(未复权)',
  `vol` float(30, 2) NULL DEFAULT NULL COMMENT '成交量(万手)',
  `amount` float(30, 2) NULL DEFAULT NULL COMMENT '成交额(万元)',
  `limit_up` tinyint(1) NULL DEFAULT NULL COMMENT '涨停板',
  PRIMARY KEY (`ts_code`, `trade_date`) USING BTREE,
  INDEX `trade_date_index`(`trade_date`) USING BTREE,
  INDEX `ts_code_index`(`ts_code`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '股票行情' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for stock_quote_2021
-- ----------------------------
DROP TABLE IF EXISTS `stock_quote_2021`;
CREATE TABLE `stock_quote_2021`  (
  `ts_code` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT 'TS代码',
  `trade_date` varchar(8) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '交易日期',
  `open` float(30, 2) NULL DEFAULT NULL COMMENT '开盘价',
  `high` float(30, 2) NULL DEFAULT NULL COMMENT '最高价',
  `low` float(30, 2) NULL DEFAULT NULL COMMENT '最低价',
  `close` float(30, 2) NULL DEFAULT NULL COMMENT '收盘价',
  `pre_close` float(30, 2) NULL DEFAULT NULL COMMENT '昨收价(前复权)',
  `change` float(30, 2) NULL DEFAULT NULL COMMENT '涨跌额',
  `pct_chg` float(30, 2) NULL DEFAULT NULL COMMENT '涨跌幅(未复权)',
  `vol` float(30, 2) NULL DEFAULT NULL COMMENT '成交量(万手)',
  `amount` float(30, 2) NULL DEFAULT NULL COMMENT '成交额(万元)',
  `limit_up` tinyint(1) NULL DEFAULT NULL COMMENT '涨停板',
  PRIMARY KEY (`ts_code`, `trade_date`) USING BTREE,
  INDEX `trade_date_index`(`trade_date`) USING BTREE,
  INDEX `ts_code_index`(`ts_code`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '股票行情' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for stock_quote_2022
-- ----------------------------
DROP TABLE IF EXISTS `stock_quote_2022`;
CREATE TABLE `stock_quote_2022`  (
  `ts_code` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT 'TS代码',
  `trade_date` varchar(8) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '交易日期',
  `open` float(30, 2) NULL DEFAULT NULL COMMENT '开盘价',
  `high` float(30, 2) NULL DEFAULT NULL COMMENT '最高价',
  `low` float(30, 2) NULL DEFAULT NULL COMMENT '最低价',
  `close` float(30, 2) NULL DEFAULT NULL COMMENT '收盘价',
  `pre_close` float(30, 2) NULL DEFAULT NULL COMMENT '昨收价(前复权)',
  `change` float(30, 2) NULL DEFAULT NULL COMMENT '涨跌额',
  `pct_chg` float(30, 2) NULL DEFAULT NULL COMMENT '涨跌幅(未复权)',
  `vol` float(30, 2) NULL DEFAULT NULL COMMENT '成交量(万手)',
  `amount` float(30, 2) NULL DEFAULT NULL COMMENT '成交额(万元)',
  `limit_up` tinyint(1) NULL DEFAULT NULL COMMENT '涨停板',
  PRIMARY KEY (`ts_code`, `trade_date`) USING BTREE,
  INDEX `trade_date_index`(`trade_date`) USING BTREE,
  INDEX `ts_code_index`(`ts_code`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '股票行情' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for stock_quote_2023
-- ----------------------------
DROP TABLE IF EXISTS `stock_quote_2023`;
CREATE TABLE `stock_quote_2023`  (
  `ts_code` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT 'TS代码',
  `trade_date` varchar(8) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '交易日期',
  `open` float(30, 2) NULL DEFAULT NULL COMMENT '开盘价',
  `high` float(30, 2) NULL DEFAULT NULL COMMENT '最高价',
  `low` float(30, 2) NULL DEFAULT NULL COMMENT '最低价',
  `close` float(30, 2) NULL DEFAULT NULL COMMENT '收盘价',
  `pre_close` float(30, 2) NULL DEFAULT NULL COMMENT '昨收价(前复权)',
  `change` float(30, 2) NULL DEFAULT NULL COMMENT '涨跌额',
  `pct_chg` float(30, 2) NULL DEFAULT NULL COMMENT '涨跌幅(未复权)',
  `vol` float(30, 2) NULL DEFAULT NULL COMMENT '成交量(万手)',
  `amount` float(30, 2) NULL DEFAULT NULL COMMENT '成交额(万元)',
  `limit_up` tinyint(1) NULL DEFAULT NULL COMMENT '涨停板',
  PRIMARY KEY (`ts_code`, `trade_date`) USING BTREE,
  INDEX `trade_date_index`(`trade_date`) USING BTREE,
  INDEX `ts_code_index`(`ts_code`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '股票行情' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for stock_quote_2024
-- ----------------------------
DROP TABLE IF EXISTS `stock_quote_2024`;
CREATE TABLE `stock_quote_2024`  (
  `ts_code` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT 'TS代码',
  `trade_date` varchar(8) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '交易日期',
  `open` float(30, 2) NULL DEFAULT NULL COMMENT '开盘价',
  `high` float(30, 2) NULL DEFAULT NULL COMMENT '最高价',
  `low` float(30, 2) NULL DEFAULT NULL COMMENT '最低价',
  `close` float(30, 2) NULL DEFAULT NULL COMMENT '收盘价',
  `pre_close` float(30, 2) NULL DEFAULT NULL COMMENT '昨收价(前复权)',
  `change` float(30, 2) NULL DEFAULT NULL COMMENT '涨跌额',
  `pct_chg` float(30, 2) NULL DEFAULT NULL COMMENT '涨跌幅(未复权)',
  `vol` float(30, 2) NULL DEFAULT NULL COMMENT '成交量(万手)',
  `amount` float(30, 2) NULL DEFAULT NULL COMMENT '成交额(万元)',
  `limit_up` tinyint(1) NULL DEFAULT NULL COMMENT '涨停板',
  PRIMARY KEY (`ts_code`, `trade_date`) USING BTREE,
  INDEX `trade_date_index`(`trade_date`) USING BTREE,
  INDEX `ts_code_index`(`ts_code`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '股票行情' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for task_info
-- ----------------------------
DROP TABLE IF EXISTS `task_info`;
CREATE TABLE `task_info`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `task_name` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '任务名称',
  `date` varchar(8) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '日期',
  `spend_time` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '花费时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 88 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '定时任务执行记录' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
