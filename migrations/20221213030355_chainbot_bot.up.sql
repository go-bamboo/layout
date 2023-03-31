/*
 Navicat Premium Data Transfer

 Source Server         : vultr的db
 Source Server Type    : SQLite
 Source Server Version : 3030001
 Source Schema         : main

 Target Server Type    : SQLite
 Target Server Version : 3030001
 File Encoding         : 65001

 Date: 13/12/2022 15:22:06
*/

-- ----------------------------
-- Table structure for layout_bot
-- ----------------------------
CREATE TABLE "layout_bot" (
    "id" bigint unsigned,
    "source" integer,
    "bot_id" varchar(200),
    "sub_id" varchar(200),
    "driver_type" varchar(200),
    "notify_template" text,
    "template_id" bigint,
    "created_at" datetime,
    "updated_at" datetime,
    "deleted_at" datetime,
    PRIMARY KEY ("id")
);
