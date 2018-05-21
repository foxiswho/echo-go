/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 50721
 Source Host           : localhost:3306
 Source Schema         : shop_go

 Target Server Type    : MySQL
 Target Server Version : 50721
 File Encoding         : 65001

 Date: 20/05/2018 21:50:32
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin
-- ----------------------------
DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `username` char(30) DEFAULT NULL COMMENT '用户名',
  `password` char(32) DEFAULT NULL COMMENT '密码',
  `mail` varchar(80) DEFAULT NULL COMMENT '邮箱',
  `salt` varchar(10) DEFAULT NULL COMMENT '干扰码',
  `gmt_create` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `gmt_modified` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `ip` char(15) DEFAULT NULL COMMENT '添加IP',
  `job_no` varchar(15) DEFAULT NULL COMMENT '工号',
  `nick_name` varchar(50) DEFAULT NULL COMMENT '昵称',
  `true_name` varchar(50) DEFAULT NULL COMMENT '真实姓名',
  `qq` varchar(50) DEFAULT NULL COMMENT 'qq',
  `phone` varchar(50) DEFAULT NULL COMMENT '电话',
  `mobile` varchar(20) DEFAULT NULL COMMENT '手机',
  `name` varchar(255) DEFAULT NULL COMMENT '显示名称',
  `is_del` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '删除0否1是',
  `department_id` int(11) NOT NULL DEFAULT '0' COMMENT '部门id',
  `team_id` int(11) NOT NULL COMMENT '团队ID',
  `master_id` int(11) NOT NULL COMMENT '师傅id',
  `leader_id` int(11) NOT NULL COMMENT '领导id',
  `post_id` int(11) NOT NULL COMMENT '职务id',
  `role_id` int(11) NOT NULL COMMENT '角色id(主)',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `username` (`username`),
  KEY `is_del` (`is_del`),
  KEY `role_id` (`role_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='管理员表';

-- ----------------------------
-- Records of admin
-- ----------------------------
BEGIN;
INSERT INTO `admin` VALUES (1, 'admin', 'admin', NULL, NULL, '2018-05-18 10:04:57', '2018-05-18 10:04:57', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 0, 0, 0, 0, 0, 1);
COMMIT;

-- ----------------------------
-- Table structure for admin_menu
-- ----------------------------
DROP TABLE IF EXISTS `admin_menu`;
CREATE TABLE `admin_menu` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` char(100) DEFAULT NULL COMMENT '名称',
  `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '上级菜单',
  `s` char(60) DEFAULT NULL COMMENT '模块/控制器/动作',
  `data` char(100) DEFAULT NULL COMMENT '其他参数',
  `sort` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `type` char(32) NOT NULL DEFAULT 'url' COMMENT '类别url菜单function独立功能user用户独有',
  `level` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '级别',
  `level1_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '1级栏目ID',
  `md5` char(32) DEFAULT NULL COMMENT 's的md5值',
  `is_show` tinyint(1) NOT NULL DEFAULT '1' COMMENT '显示隐藏;1显示;0隐藏',
  `is_unique` tinyint(1) NOT NULL DEFAULT '0' COMMENT '用户独有此功能1是0否',
  PRIMARY KEY (`id`),
  KEY `sort` (`sort`),
  KEY `parent_id` (`parent_id`),
  KEY `s` (`s`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='菜单';

-- ----------------------------
-- Table structure for admin_role_access
-- ----------------------------
DROP TABLE IF EXISTS `admin_role_access`;
CREATE TABLE `admin_role_access` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `aid` int(11) NOT NULL DEFAULT '0' COMMENT '管理员ID',
  `role_id` int(11) NOT NULL DEFAULT '0' COMMENT '角色ID',
  `is_default` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否默认',
  PRIMARY KEY (`id`),
  UNIQUE KEY `aid_role_id` (`aid`,`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='管理员与角色的关系、一个管理员可以有多个角色';

-- ----------------------------
-- Table structure for admin_role_priv
-- ----------------------------
DROP TABLE IF EXISTS `admin_role_priv`;
CREATE TABLE `admin_role_priv` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `role_id` smallint(3) unsigned NOT NULL DEFAULT '0' COMMENT '角色ID',
  `s` char(100) DEFAULT NULL COMMENT '模块/控制器/动作',
  `data` char(50) DEFAULT NULL COMMENT '其他参数',
  `aid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '管理员ID',
  `menu_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '菜单ID',
  `type` char(32) NOT NULL DEFAULT 'url' COMMENT '类别url菜单function独立功能user用户独有',
  PRIMARY KEY (`id`),
  KEY `role_id` (`role_id`),
  KEY `role_id_2` (`role_id`,`s`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='角色权限表';

-- ----------------------------
-- Table structure for admin_status
-- ----------------------------
DROP TABLE IF EXISTS `admin_status`;
CREATE TABLE `admin_status` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `login_time` timestamp NULL DEFAULT NULL COMMENT '登录时间',
  `login_ip` char(20) DEFAULT NULL COMMENT 'IP',
  `login` int(11) NOT NULL DEFAULT '0' COMMENT '登录次数',
  `aid_add` int(11) NOT NULL DEFAULT '0' COMMENT '添加人',
  `aid_update` int(11) NOT NULL DEFAULT '0' COMMENT '更新人',
  `gmt_modified` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='状态';

-- ----------------------------
-- Table structure for app
-- ----------------------------
DROP TABLE IF EXISTS `app`;
CREATE TABLE `app` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type_id` int(11) NOT NULL DEFAULT '0' COMMENT 'app_id,来源type表',
  `name` varchar(100) DEFAULT NULL COMMENT '名称',
  `mark` char(32) DEFAULT NULL COMMENT '标志',
  `setting` varchar(5000) DEFAULT NULL COMMENT '扩展参数',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `is_del` int(11) NOT NULL DEFAULT '0' COMMENT '是否删除0否1是',
  `gmt_create` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  `gmt_modified` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `type_id` (`type_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='应用';

-- ----------------------------
-- Table structure for area
-- ----------------------------
DROP TABLE IF EXISTS `area`;
CREATE TABLE `area` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` char(50) DEFAULT '' COMMENT '名称',
  `name_en` varchar(100) DEFAULT '' COMMENT '英文名称',
  `parent_id` int(11) DEFAULT '0' COMMENT '上级栏目ID',
  `type` tinyint(4) DEFAULT '0' COMMENT '类别;0默认;',
  `name_traditional` varchar(50) DEFAULT '' COMMENT '繁体名称',
  `sort` int(11) DEFAULT '0' COMMENT '排序',
  PRIMARY KEY (`id`),
  KEY `parent_id` (`parent_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='地区表';

-- ----------------------------
-- Table structure for area_ext
-- ----------------------------
DROP TABLE IF EXISTS `area_ext`;
CREATE TABLE `area_ext` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `area_id` int(11) DEFAULT '0' COMMENT 'ID',
  `name` char(50) DEFAULT '' COMMENT '名称',
  `name_en` varchar(100) DEFAULT '' COMMENT '英文名称',
  `parent_id` int(11) DEFAULT '0' COMMENT '上级栏目ID',
  `type` tinyint(4) DEFAULT '0' COMMENT '类别;0默认;1又名;2;3属于;11已合并到;12已更名为',
  `name_traditional` varchar(50) DEFAULT '' COMMENT '繁体名称',
  `sort` int(11) DEFAULT '0' COMMENT '排序',
  `type_name` varchar(50) DEFAULT '' COMMENT '类别名称',
  `other_name` varchar(50) DEFAULT '' COMMENT '根据类别名称填写',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `id` (`area_id`,`parent_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='地区扩展表';

-- ----------------------------
-- Table structure for attachment
-- ----------------------------
DROP TABLE IF EXISTS `attachment`;
CREATE TABLE `attachment` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '附件ID',
  `module` char(32) DEFAULT NULL COMMENT '模块',
  `mark` char(60) DEFAULT NULL COMMENT '标记标志',
  `type_id` int(5) unsigned NOT NULL DEFAULT '0' COMMENT '类别ID',
  `name` char(50) DEFAULT NULL COMMENT '保存的文件名称',
  `name_original` varchar(255) DEFAULT NULL COMMENT '原文件名',
  `path` char(200) DEFAULT NULL COMMENT '文件路径',
  `size` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '文件大小',
  `ext` char(10) DEFAULT NULL COMMENT '文件后缀',
  `is_image` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否图片1是0否',
  `is_thumb` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否缩略图1是0否',
  `downloads` int(8) unsigned NOT NULL DEFAULT '0' COMMENT '下载次数',
  `gmt_create` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间上传时间',
  `ip` char(15) DEFAULT NULL COMMENT '上传IP',
  `status` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '状态99正常;',
  `md5` char(32) DEFAULT NULL COMMENT 'md5',
  `sha1` char(40) DEFAULT NULL COMMENT 'sha1',
  `from_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '所属ID',
  `aid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '后台管理员ID',
  `uid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '前台用户ID',
  `is_show` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '是否显示1是0否',
  `http` varchar(100) DEFAULT NULL COMMENT '图片http地址',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `md5` (`md5`),
  KEY `module` (`module`),
  KEY `mark` (`mark`),
  KEY `id` (`from_id`),
  KEY `status` (`status`),
  KEY `aid` (`aid`),
  KEY `uid` (`uid`),
  KEY `is_show` (`is_show`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='附件表';

-- ----------------------------
-- Records of attachment
-- ----------------------------
BEGIN;
INSERT INTO `attachment` VALUES (1, '', '', 1, '302b4f78f934cf4a06c5d8a9257ed97c.jpg', '9358d109b3de9c82bb32fd2d6081800a19d84338.jpg', '/uploads/image/2018-05/', 313768, 'jpg', 0, 0, 0, '2018-05-16 22:15:11', '', 0, '8444e2b2a9f99cae7a855a5f887dd3e2', '', 0, 1, 0, 0, '');
COMMIT;

-- ----------------------------
-- Table structure for cart
-- ----------------------------
DROP TABLE IF EXISTS `cart`;
CREATE TABLE `cart` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `goods_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '商品ID',
  `product_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '商品信息id',
  `num` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '数量',
  `price` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '单价',
  `amount` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '合计总价',
  `warehouse_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '仓库ID',
  `sid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '供货商ID',
  `type_id` int(11) NOT NULL DEFAULT '1' COMMENT '类别:普通',
  PRIMARY KEY (`id`),
  KEY `uid` (`uid`),
  KEY `type_id` (`type_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='购物车';

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule` (
  `p_type` varchar(100) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL,
  KEY `IDX_casbin_rule_p_type` (`p_type`),
  KEY `IDX_casbin_rule_v0` (`v0`),
  KEY `IDX_casbin_rule_v1` (`v1`),
  KEY `IDX_casbin_rule_v2` (`v2`),
  KEY `IDX_casbin_rule_v3` (`v3`),
  KEY `IDX_casbin_rule_v4` (`v4`),
  KEY `IDX_casbin_rule_v5` (`v5`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
BEGIN;
INSERT INTO `casbin_rule` VALUES ('p', '1', '/admin/rbac/index', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '/admin/rdbc/*', 'GET', '', '', '');
COMMIT;

-- ----------------------------
-- Table structure for connect
-- ----------------------------
DROP TABLE IF EXISTS `connect`;
CREATE TABLE `connect` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `type_id` int(11) NOT NULL DEFAULT '0' COMMENT '类别id',
  `uid` int(11) NOT NULL DEFAULT '0' COMMENT '用户id',
  `open_id` char(80) DEFAULT NULL COMMENT '对应唯一开放id',
  `token` varchar(80) DEFAULT NULL COMMENT '开放密钥',
  `type` int(11) NOT NULL DEFAULT '1' COMMENT '登录类型1腾讯QQ2新浪微博',
  `type_login` int(11) NOT NULL DEFAULT '0' COMMENT '登录模块;302前台还是后台301',
  `gmt_create` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `extend` varchar(5000) DEFAULT '' COMMENT '扩展参数',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `openid` (`open_id`),
  KEY `uid` (`uid`) USING BTREE,
  KEY `type_id` (`type_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='快捷登陆/qq';

-- ----------------------------
-- Table structure for goods
-- ----------------------------
DROP TABLE IF EXISTS `goods`;
CREATE TABLE `goods` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `product_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '商品信息ID',
  `warehouse_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '仓库ID',
  `sid` int(11) NOT NULL DEFAULT '0' COMMENT '供应商ID',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '状态0未审核99已审核',
  `is_del` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除1是0否',
  `is_open` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否上架1是0否',
  `aid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '管理员（发布人）ID',
  `cat_id` int(10) unsigned NOT NULL DEFAULT '1' COMMENT '栏目id',
  `brand_id` int(10) NOT NULL DEFAULT '0' COMMENT '品牌',
  `title` varchar(100) DEFAULT NULL COMMENT '标题',
  `model` varchar(100) DEFAULT NULL COMMENT '规格',
  `number` char(100) DEFAULT NULL COMMENT '商品编号',
  `thumb` varchar(255) DEFAULT NULL COMMENT '缩略图',
  `original_img` varchar(255) DEFAULT NULL COMMENT '原始图',
  `sort` int(10) NOT NULL DEFAULT '0' COMMENT '排序',
  `price_base` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '底价',
  `price_plantform_cost` int(12) unsigned NOT NULL DEFAULT '0' COMMENT '平台成本',
  `attr_type_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '属性类别ID',
  `num_unit` int(11) unsigned NOT NULL DEFAULT '1' COMMENT '每个单位内多少个，每盒几罐',
  `type_stock` int(10) NOT NULL DEFAULT '0' COMMENT '是否仓库库存',
  `type_id` int(11) NOT NULL DEFAULT '10001' COMMENT '类别类目',
  `mark` char(32) NOT NULL DEFAULT '' COMMENT '标志:产品-仓库-供应商',
  `mark_id` int(11) NOT NULL DEFAULT '10001' COMMENT '标志类别',
  `is_combined` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否商品组合',
  `description` varchar(255) DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `aid` (`aid`) USING BTREE,
  KEY `cat_id` (`cat_id`) USING BTREE,
  KEY `brand_id` (`brand_id`) USING BTREE,
  KEY `is_open` (`is_open`),
  KEY `sort` (`sort`),
  KEY `status` (`status`),
  KEY `is_del` (`is_del`),
  KEY `sid` (`sid`),
  KEY `mark` (`mark`),
  KEY `mark_id` (`mark_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1006 DEFAULT CHARSET=utf8mb4 COMMENT='商品发布';

-- ----------------------------
-- Records of goods
-- ----------------------------
BEGIN;
INSERT INTO `goods` VALUES (1002, 1, 1, 1, 99, 0, 1, 0, 1, 1, '德国Aptamil爱他美婴幼儿配方奶粉1+段(适合1岁以上宝宝)600g', '1+段 600g', 'A000001', '', '', 0, 0, 0, 0, 1, 0, 202001, '𱄩-Ϫ', 201001, 0, '');
COMMIT;

-- ----------------------------
-- Table structure for goods_brand
-- ----------------------------
DROP TABLE IF EXISTS `goods_brand`;
CREATE TABLE `goods_brand` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL COMMENT '名称',
  `name_en` varchar(255) DEFAULT NULL COMMENT '品牌英文名称或是汉语拼音',
  `http` varchar(255) DEFAULT NULL COMMENT '品牌网站',
  `phone` varchar(255) DEFAULT NULL COMMENT '客服电话',
  `content` text COMMENT '品牌介绍',
  `letter` varchar(255) DEFAULT NULL COMMENT '品牌首字母',
  `sort` int(10) NOT NULL DEFAULT '0' COMMENT '品牌排序',
  `logo` varchar(255) DEFAULT NULL COMMENT '品牌logo',
  `parent_id` int(10) unsigned NOT NULL DEFAULT '0',
  `arr_parent_id` varchar(255) DEFAULT NULL COMMENT '所有父栏目ID',
  `is_child` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否有子栏目',
  `arr_child_id` text COMMENT '所有子栏目ID',
  `is_del` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除1是0否',
  PRIMARY KEY (`id`),
  KEY `is_del` (`is_del`),
  KEY `parent_id` (`parent_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='品牌';

-- ----------------------------
-- Table structure for goods_category
-- ----------------------------
DROP TABLE IF EXISTS `goods_category`;
CREATE TABLE `goods_category` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL COMMENT '名称',
  `description` text COMMENT '介绍',
  `parent_id` int(10) unsigned NOT NULL DEFAULT '0',
  `sort` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `arr_parent_id` varchar(255) DEFAULT NULL COMMENT '所有父栏目ID',
  `is_child` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否有子栏目',
  `arr_child_id` text COMMENT '所有子栏目ID',
  `is_del` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除1是0否',
  `gmt_create` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  `gmt_modified` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `parent_id` (`parent_id`) USING BTREE,
  KEY `sort` (`sort`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='栏目';

-- ----------------------------
-- Table structure for goods_combined
-- ----------------------------
DROP TABLE IF EXISTS `goods_combined`;
CREATE TABLE `goods_combined` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `goods_id` int(11) NOT NULL DEFAULT '0',
  `product_id` int(11) NOT NULL DEFAULT '0' COMMENT '产品ID',
  `aid` int(11) NOT NULL DEFAULT '0' COMMENT '添加人',
  `price_shop` bigint(20) NOT NULL DEFAULT '0' COMMENT '组合商品价格',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  `num_least` int(11) NOT NULL DEFAULT '1' COMMENT '最少购买数量',
  `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '顶级商品ID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='组合商品';

-- ----------------------------
-- Table structure for goods_content
-- ----------------------------
DROP TABLE IF EXISTS `goods_content`;
CREATE TABLE `goods_content` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `seo_title` varchar(50) DEFAULT NULL COMMENT 'seo标题',
  `seo_description` varchar(200) DEFAULT NULL COMMENT 'seo描述',
  `seo_keyword` varchar(50) DEFAULT NULL COMMENT 'seo关键字',
  `content` text COMMENT '内容',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注紧供自己查看',
  `title_other` varchar(5000) DEFAULT NULL COMMENT '其他名称',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1006 DEFAULT CHARSET=utf8mb4 COMMENT='商品内容';

-- ----------------------------
-- Records of goods_content
-- ----------------------------
BEGIN;
INSERT INTO `goods_content` VALUES (1002, '', '', '', '内容', '', '');
COMMIT;

-- ----------------------------
-- Table structure for goods_price
-- ----------------------------
DROP TABLE IF EXISTS `goods_price`;
CREATE TABLE `goods_price` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `price_market` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '市场价',
  `price_shop` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '商城价',
  `is_promote` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否促销1是0否',
  `promote_price` int(20) unsigned NOT NULL DEFAULT '0' COMMENT '促销价格',
  `promote_start_date` timestamp NULL DEFAULT NULL COMMENT '促销开始日期',
  `promote_end_date` timestamp NULL DEFAULT NULL COMMENT '促销结束日期',
  `is_free_shipping` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否包邮1是0否',
  `start_date` timestamp NULL DEFAULT NULL COMMENT '开始时间',
  `end_date` timestamp NULL DEFAULT NULL COMMENT '结束时间',
  `min_free_shipping` int(10) unsigned NOT NULL DEFAULT '1' COMMENT '最小包邮数量',
  `num_max` int(11) NOT NULL DEFAULT '9999' COMMENT '最大可一次购买数量',
  `num_least` int(11) NOT NULL DEFAULT '1' COMMENT '最少购买数量',
  `is_free_tax` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否包税使用包税价格',
  `is_group_price` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否使用用户组价格',
  `tax_price` bigint(20) NOT NULL DEFAULT '0' COMMENT '包税价格',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1006 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of goods_price
-- ----------------------------
BEGIN;
INSERT INTO `goods_price` VALUES (1002, 1400000, 1200000, 0, 0, NULL, NULL, 0, NULL, NULL, 0, 999, 1, 0, 0, 0);
COMMIT;

-- ----------------------------
-- Table structure for goods_statistics
-- ----------------------------
DROP TABLE IF EXISTS `goods_statistics`;
CREATE TABLE `goods_statistics` (
  `id` int(11) NOT NULL,
  `saless` bigint(20) NOT NULL DEFAULT '0' COMMENT '销量',
  `reading` bigint(20) NOT NULL DEFAULT '0' COMMENT '访问数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品销量';

-- ----------------------------
-- Table structure for log
-- ----------------------------
DROP TABLE IF EXISTS `log`;
CREATE TABLE `log` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `from_id` int(11) NOT NULL DEFAULT '0' COMMENT 'id',
  `aid` int(11) NOT NULL DEFAULT '0' COMMENT '管理员ID',
  `uid` int(11) NOT NULL DEFAULT '0' COMMENT '用户id',
  `gmt_create` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `mark` char(32) DEFAULT NULL COMMENT '标志自定义标志',
  `data` text COMMENT '其他内容',
  `no` char(50) DEFAULT NULL COMMENT '单号',
  `type_login` int(11) NOT NULL DEFAULT '0' COMMENT '登录方式;302前台还是后台301',
  `type_client` int(11) NOT NULL DEFAULT '0' COMMENT '登录客户端类别;321电脑;322安卓;323IOS;324手机网页;325其他',
  `ip` char(20) DEFAULT NULL COMMENT 'IP',
  `msg` varchar(255) DEFAULT NULL COMMENT '自定义说明',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `type_login` (`type_login`),
  KEY `type_client` (`type_client`),
  KEY `uid` (`uid`),
  KEY `aid` (`aid`),
  KEY `id` (`from_id`),
  KEY `no` (`no`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='日志表';

-- ----------------------------
-- Table structure for news
-- ----------------------------
DROP TABLE IF EXISTS `news`;
CREATE TABLE `news` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `aid` int(11) NOT NULL DEFAULT '0' COMMENT '管理员AID',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除1是0否',
  `is_open` tinyint(1) NOT NULL DEFAULT '1' COMMENT '启用1是0否',
  `status` int(11) NOT NULL DEFAULT '0' COMMENT '状态',
  `gmt_system` timestamp NULL DEFAULT NULL COMMENT '创建时间,系统时间不可修改',
  `gmt_modified` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `gmt_create` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间,可修改',
  `title` varchar(255) DEFAULT NULL COMMENT '标题',
  `author` varchar(255) DEFAULT NULL COMMENT '作者',
  `url` varchar(255) DEFAULT NULL COMMENT '网址',
  `url_source` varchar(255) DEFAULT NULL COMMENT '来源地址(转载)',
  `url_rewrite` char(150) DEFAULT NULL COMMENT '自定义伪静态Url',
  `description` varchar(255) DEFAULT NULL COMMENT '摘要',
  `content` text COMMENT '内容',
  `type` int(11) NOT NULL DEFAULT '0' COMMENT '类型0文章10001博客栏目',
  `module_id` int(10) NOT NULL DEFAULT '0' COMMENT '模块10019技术10018生活',
  `source_id` int(11) NOT NULL DEFAULT '0' COMMENT '来源:后台，接口，其他',
  `type_id` int(11) NOT NULL DEFAULT '0' COMMENT '类别ID，原创，转载，翻译',
  `cat_id` int(11) NOT NULL DEFAULT '0' COMMENT '分类ID，栏目',
  `tag` varchar(255) DEFAULT NULL COMMENT '标签',
  `thumb` varchar(255) DEFAULT NULL COMMENT '缩略图',
  `is_relevant` tinyint(1) NOT NULL DEFAULT '0' COMMENT '相关文章1是0否',
  `is_jump` tinyint(1) NOT NULL DEFAULT '0' COMMENT '跳转1是0否',
  `is_comment` tinyint(1) NOT NULL DEFAULT '1' COMMENT '允许评论1是0否',
  `is_read` int(11) NOT NULL DEFAULT '10014' COMMENT '是否阅读10014未看10015在看10016已看',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `is_del` (`is_del`,`is_open`,`status`,`type_id`,`cat_id`,`sort`),
  KEY `url_rewrite` (`url_rewrite`),
  KEY `type` (`type`),
  KEY `module_id` (`module_id`) USING BTREE,
  KEY `source_id` (`source_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='博客内容';

-- ----------------------------
-- Table structure for news_statistics
-- ----------------------------
DROP TABLE IF EXISTS `news_statistics`;
CREATE TABLE `news_statistics` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `news_id` int(11) NOT NULL DEFAULT '0' COMMENT '文章ID',
  `comment` int(11) NOT NULL DEFAULT '0' COMMENT '评论人数',
  `read` int(11) NOT NULL DEFAULT '0' COMMENT '阅读人数',
  `seo_title` varchar(255) DEFAULT NULL COMMENT 'SEO标题',
  `seo_description` varchar(255) DEFAULT NULL COMMENT 'SEO摘要',
  `seo_keyword` varchar(255) DEFAULT NULL COMMENT 'SEO关键词',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `blog_id` (`news_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='博客统计';

-- ----------------------------
-- Table structure for news_sync_mapping
-- ----------------------------
DROP TABLE IF EXISTS `news_sync_mapping`;
CREATE TABLE `news_sync_mapping` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `news_id` int(11) NOT NULL DEFAULT '0' COMMENT '本站blog的id',
  `type_id` int(11) NOT NULL DEFAULT '0' COMMENT '类别id',
  `to_id` varchar(64) DEFAULT NULL COMMENT 'csdn的id',
  `gmt_modified` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后一次更新时间',
  `gmt_create` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间',
  `mark` char(32) DEFAULT NULL COMMENT '标志',
  `is_sync` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否同步过',
  `extend` varchar(5000) DEFAULT NULL COMMENT '扩展参数',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='本站blog_id 与其他同步站点的id关系';

-- ----------------------------
-- Table structure for news_sync_queue
-- ----------------------------
DROP TABLE IF EXISTS `news_sync_queue`;
CREATE TABLE `news_sync_queue` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `news_id` int(11) NOT NULL DEFAULT '0' COMMENT '本站博客id',
  `type_id` int(11) NOT NULL DEFAULT '0' COMMENT '类型',
  `status` tinyint(3) NOT NULL DEFAULT '0' COMMENT '状态：0:待运行 10:失败 99:成功',
  `gmt_modified` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后一次更新时间',
  `gmt_create` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间',
  `msg` varchar(255) DEFAULT NULL COMMENT '内容',
  `map_id` int(11) NOT NULL DEFAULT '0' COMMENT '同步ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='博客同步队列';

-- ----------------------------
-- Table structure for news_tag
-- ----------------------------
DROP TABLE IF EXISTS `news_tag`;
CREATE TABLE `news_tag` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` char(100) DEFAULT NULL COMMENT '名称',
  `gmt_create` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  `aid` int(11) NOT NULL DEFAULT '0' COMMENT '管理员ID',
  `news_id` int(11) NOT NULL DEFAULT '0' COMMENT '文章ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='博客标签';

-- ----------------------------
-- Table structure for order
-- ----------------------------
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `order_no` char(32) DEFAULT NULL COMMENT '销售订单号',
  `order_sn` char(32) DEFAULT NULL COMMENT '单号淘宝苏宁等等',
  `uid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `order_status` int(11) NOT NULL DEFAULT '0' COMMENT '订单状态(DEFAULT用户未点发货,WAIT_CHECK等待审核,NO_CHECK审核不通过,WAIT_SEND等等发货,SEND卖家已发货,RECEIPT已收货,DROP交易关闭,SUCCESS订单交易成功,CANCEL交易取消,WAIT_CUSTOMS_CHECK等待海关审核 REFUND退款 DELETE删除 DRAFT 草稿)',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '状态0未审核99已审核',
  `is_del` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除1是0否',
  `type_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '类别,1普通订单;',
  `type_id_admin` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '类别,1普通订单;后台设置',
  `type_id_sub` int(11) NOT NULL DEFAULT '0' COMMENT '其他类别',
  `vat_fee` bigint(20) NOT NULL DEFAULT '0' COMMENT '增值税费',
  `sales_fee` bigint(20) NOT NULL DEFAULT '0' COMMENT '消费税',
  `amount_freight` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '物流费用',
  `amount_discount` bigint(20) NOT NULL DEFAULT '0' COMMENT '折扣金额',
  `amount_goods` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '商品总金额',
  `amount_other` bigint(20) NOT NULL COMMENT '其他价格费用',
  `amount_tax` bigint(20) NOT NULL DEFAULT '0' COMMENT '税费',
  `amount_order` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '订单总额',
  `amount_payment` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '支付总额,已付款金额(实际付款金额)',
  `total` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '总数量',
  `total_no_receipt` int(10) NOT NULL DEFAULT '0' COMMENT '未收货数量',
  `sid` int(11) NOT NULL DEFAULT '0' COMMENT '供应商ID',
  `warehouse_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '仓库ID',
  `store_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '店铺ID',
  `express_no` char(50) NOT NULL DEFAULT '' COMMENT '物流单号,运送单号',
  `express_id` int(10) NOT NULL DEFAULT '0' COMMENT '快递公司id',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注用户自己看',
  `remark_admin` varchar(255) DEFAULT NULL COMMENT '备注客服自己看',
  `gmt_create` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '下单时间',
  `gmt_modified` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
  `time_success` timestamp NULL DEFAULT NULL COMMENT '订单完成时间(整个订单完成，交易完成)',
  `time_check_admin` timestamp NULL DEFAULT NULL COMMENT '客服审核时间',
  `time_check` timestamp NULL DEFAULT NULL COMMENT '审核时间，海关审核时间',
  `time_receipt` timestamp NULL DEFAULT NULL COMMENT '收货时间',
  `declare` int(11) NOT NULL DEFAULT '0' COMMENT 'DEFAULT未申报 NOT_ALLOW禁止申报 PORT_ACCEPT申报中 SUCCESS申报成功 FAIL申报失败 WARING申报异常;总订单时SUCCESS表示本订单已全部添加完成',
  `declare_msg` varchar(200) DEFAULT NULL COMMENT '申报信息',
  `declare_time` timestamp NULL DEFAULT NULL COMMENT '申报时间',
  `is_send_time` timestamp NULL DEFAULT NULL COMMENT '发货动作时间',
  `is_send` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否发货1是0否',
  `is_refund` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否退款',
  `is_return` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '退货1是0否',
  `is_exchange` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '换货1是0否',
  `order_id_from` int(10) NOT NULL DEFAULT '0' COMMENT '來自哪个ID，修改价格前ID',
  `order_id_from_api` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '接口订单id',
  `order_id_master` int(10) NOT NULL DEFAULT '0' COMMENT '总订单号ID',
  `order_no_master` char(32) DEFAULT NULL COMMENT '总订单号',
  `sid_from` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '货源商家id',
  `pay_time` timestamp NULL DEFAULT NULL COMMENT '客户购买时间',
  `pay_id` int(11) NOT NULL DEFAULT '0' COMMENT '支付ID',
  `pay_no` char(50) DEFAULT NULL COMMENT '支付单号',
  `is_paid` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否已支付',
  `is_paid_system` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否已支付(系统自动)',
  `time_paid_system` datetime DEFAULT NULL COMMENT '系统支付时间',
  `exchange_rate` bigint(20) NOT NULL DEFAULT '0' COMMENT '汇率',
  `currency_mark` char(3) DEFAULT NULL COMMENT '币制',
  `get_id` int(11) NOT NULL DEFAULT '0' COMMENT '优惠券',
  `use_wallet` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '使用钱包',
  `use_credit` bigint(20) NOT NULL DEFAULT '0' COMMENT '使用积分',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `uid` (`uid`),
  KEY `is_send` (`is_send`),
  KEY `order_status` (`order_status`),
  KEY `status` (`status`),
  KEY `is_del` (`is_del`),
  KEY `type` (`type_id`) USING BTREE,
  KEY `is_refund` (`is_refund`) USING BTREE,
  KEY `add_time` (`gmt_create`),
  KEY `sales_no` (`order_no`),
  KEY `type_admin` (`type_id_admin`),
  KEY `sid` (`sid`),
  KEY `order_no_parent` (`order_no_master`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单';

-- ----------------------------
-- Table structure for order_consignee
-- ----------------------------
DROP TABLE IF EXISTS `order_consignee`;
CREATE TABLE `order_consignee` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '订单id',
  `consignee_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '收货人ID',
  `consignee` varchar(50) DEFAULT NULL COMMENT '收货人',
  `mobile` char(11) DEFAULT NULL COMMENT '手机号',
  `country` int(11) unsigned NOT NULL DEFAULT '1' COMMENT '国家',
  `province` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '省',
  `city` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '市',
  `district` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '区',
  `address` varchar(255) DEFAULT NULL COMMENT '地址',
  `address_en` varchar(255) DEFAULT NULL COMMENT '地址(英文)',
  `id_card` char(19) DEFAULT NULL COMMENT '身份证号',
  `id_card_front` varchar(255) DEFAULT NULL COMMENT '身份证正面',
  `id_card_back` varchar(255) DEFAULT NULL COMMENT '身份证反面',
  `province_name` char(30) DEFAULT NULL COMMENT '省',
  `city_name` char(50) DEFAULT NULL COMMENT '市',
  `district_name` char(50) DEFAULT NULL COMMENT '区',
  `address_all` varchar(255) DEFAULT NULL COMMENT '地址',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单收货人';

-- ----------------------------
-- Table structure for order_ext
-- ----------------------------
DROP TABLE IF EXISTS `order_ext`;
CREATE TABLE `order_ext` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `cost_amount_payment` bigint(20) NOT NULL DEFAULT '0' COMMENT '支付金额成本',
  `cost_amount_goods` bigint(20) NOT NULL DEFAULT '0' COMMENT '商品金额成本',
  `declare_tax` int(11) NOT NULL DEFAULT '0' COMMENT '申报税',
  `declare_vat_fee` int(11) NOT NULL DEFAULT '0' COMMENT '申报增值税',
  `declare_sales_fee` int(11) NOT NULL DEFAULT '0' COMMENT '申报消费税',
  `declare_amount_freight` int(11) NOT NULL DEFAULT '0' COMMENT '申报运费',
  `declare_amount_payment` int(11) NOT NULL DEFAULT '0' COMMENT '申报运费',
  `is_wms_sales` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否向WMS通信生成销售单   0未通信 1已通信',
  `is_wms_send_out` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否向WMS发送并生成出货单  0未完成 1已完成',
  `order_amount_declare` bigint(20) NOT NULL DEFAULT '0' COMMENT '订单总金额',
  `payment_amount_declare` bigint(20) NOT NULL DEFAULT '0' COMMENT '支付总金额',
  `goods_amount_declare` bigint(20) NOT NULL DEFAULT '0' COMMENT '商品小计',
  `billing_country` int(11) unsigned NOT NULL DEFAULT '1' COMMENT '账单国家',
  `billing_province` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '账单省',
  `billing_city` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '账单市',
  `billing_district` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '账单区',
  `billing_address` varchar(255) NOT NULL DEFAULT '' COMMENT '账单地址',
  `billing_mobile` char(11) NOT NULL DEFAULT '' COMMENT '手机号',
  `billing_consignee` varchar(255) NOT NULL DEFAULT '' COMMENT '账单收货人',
  `billing_mail` varchar(255) NOT NULL DEFAULT '' COMMENT '账单邮箱',
  `billing_address_en` varchar(255) DEFAULT NULL COMMENT '账单地址(英文)',
  `billing_zip_code` varchar(10) DEFAULT NULL COMMENT '账单邮编',
  `billing_tax_no` varchar(255) DEFAULT NULL COMMENT '税号',
  `packing_id` int(10) DEFAULT '0' COMMENT '包装ID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='扩展订单信息';

-- ----------------------------
-- Table structure for order_goods
-- ----------------------------
DROP TABLE IF EXISTS `order_goods`;
CREATE TABLE `order_goods` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `order_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '订单ID',
  `goods_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '商品ID',
  `product_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '商品信息id',
  `title` varchar(200) DEFAULT NULL COMMENT '商品名称',
  `num` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '数量',
  `number` char(100) DEFAULT NULL COMMENT '商品编号',
  `price` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '单价',
  `num_unit` int(11) NOT NULL DEFAULT '1' COMMENT '每个单位内多少个，每盒几罐',
  `num_total` int(11) NOT NULL DEFAULT '0' COMMENT '总数量 = 罐数x页面数量',
  `amount` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '合计总价',
  `freight` int(11) NOT NULL DEFAULT '0' COMMENT '运费',
  `warehouse_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '仓库ID',
  `sid` int(10) unsigned NOT NULL DEFAULT '1' COMMENT '商家ID',
  `sales_fee` bigint(20) NOT NULL DEFAULT '0' COMMENT '消费税费',
  `vat_fee` bigint(20) NOT NULL DEFAULT '0' COMMENT '增值税费',
  `price_tax` bigint(20) NOT NULL DEFAULT '0' COMMENT '总税费',
  `remark` text COMMENT '备注',
  `price_shop` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '商城价',
  `cost_price` bigint(20) NOT NULL DEFAULT '0' COMMENT '成本单价',
  `cost_amount` bigint(20) NOT NULL DEFAULT '0' COMMENT '成本金额',
  `mark_id` int(11) NOT NULL DEFAULT '0' COMMENT '商品标志ID',
  PRIMARY KEY (`id`),
  KEY `sales_id` (`order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='订单商品';

-- ----------------------------
-- Table structure for order_goods_structure
-- ----------------------------
DROP TABLE IF EXISTS `order_goods_structure`;
CREATE TABLE `order_goods_structure` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `order_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '订单ID',
  `goods_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '商品ID',
  `product_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '商品信息id',
  `title` varchar(200) DEFAULT NULL COMMENT '商品名称',
  `num` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '数量',
  `number` char(100) DEFAULT NULL COMMENT '商品编号',
  `price` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '单价',
  `num_unit` int(11) NOT NULL DEFAULT '1' COMMENT '每个单位内多少个，每盒几罐',
  `num_total` int(11) NOT NULL DEFAULT '0' COMMENT '总数量 = 罐数x页面数量',
  `amount` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '合计总价',
  `freight` bigint(20) NOT NULL DEFAULT '0' COMMENT '运费',
  `warehouse_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '仓库ID',
  `sid` int(10) unsigned NOT NULL DEFAULT '1' COMMENT '商家ID',
  `sales_fee` bigint(11) NOT NULL DEFAULT '0' COMMENT '消费税费',
  `vat_fee` bigint(10) NOT NULL DEFAULT '0' COMMENT '增值税费',
  `price_tax` bigint(10) NOT NULL DEFAULT '0' COMMENT '总税费',
  `remark` text COMMENT '备注',
  `price_shop` bigint(12) unsigned DEFAULT '0' COMMENT '商城价',
  `cost_price` bigint(11) NOT NULL DEFAULT '0' COMMENT '成本单价',
  `cost_amount` bigint(20) NOT NULL DEFAULT '0' COMMENT '成本金额',
  `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '所属组合商品ID',
  PRIMARY KEY (`id`),
  KEY `sales_id` (`order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='订单商品';

-- ----------------------------
-- Table structure for session
-- ----------------------------
DROP TABLE IF EXISTS `session`;
CREATE TABLE `session` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `uid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '用户UID',
  `ip` char(15) DEFAULT NULL COMMENT 'IP',
  `error_count` tinyint(1) NOT NULL DEFAULT '0' COMMENT '密码输入错误次数',
  `app_id` int(11) NOT NULL DEFAULT '0' COMMENT '登录应用',
  `md5` char(32) DEFAULT NULL COMMENT 'md5',
  `type_login` int(11) NOT NULL DEFAULT '0' COMMENT '登录方式;302前台还是后台301',
  `type_client` int(11) NOT NULL DEFAULT '0' COMMENT '登录客户端类别;321电脑;322安卓;323IOS;324手机网页;325其他',
  `gmt_create` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登录时间',
  `gmt_modified` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `uid` (`uid`,`type_login`,`type_client`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='SESSION';

-- ----------------------------
-- Table structure for tag
-- ----------------------------
DROP TABLE IF EXISTS `tag`;
CREATE TABLE `tag` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` char(50) DEFAULT NULL COMMENT '名称',
  `gmt_create` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='标签';

-- ----------------------------
-- Table structure for template
-- ----------------------------
DROP TABLE IF EXISTS `template`;
CREATE TABLE `template` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '模板ID',
  `name` varchar(80) DEFAULT NULL COMMENT '模板名称(中文)',
  `mark` varchar(80) DEFAULT NULL COMMENT '模板名称标志(英文)（调用时使用）',
  `title` varchar(255) DEFAULT NULL COMMENT '邮件标题',
  `type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '模板类型1短信模板2邮箱模板',
  `use` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '用途',
  `content` text,
  `remark` varchar(1024) DEFAULT NULL COMMENT '备注',
  `gmt_create` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `gmt_modified` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `code_num` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '验证码位数',
  `aid` int(11) NOT NULL DEFAULT '0' COMMENT '添加人',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='会员表';

-- ----------------------------
-- Table structure for type
-- ----------------------------
DROP TABLE IF EXISTS `type`;
CREATE TABLE `type` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` char(100) DEFAULT NULL COMMENT '名称',
  `name_en` char(100) DEFAULT NULL COMMENT '名称',
  `code` char(32) DEFAULT NULL COMMENT '代码',
  `mark` char(32) DEFAULT NULL COMMENT '标志',
  `type_id` int(11) NOT NULL DEFAULT '0' COMMENT '所属类别ID',
  `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '上级ID、属于/上级ID',
  `value` int(10) NOT NULL DEFAULT '0' COMMENT '值',
  `content` varchar(255) DEFAULT NULL COMMENT '内容',
  `is_del` int(11) NOT NULL DEFAULT '0' COMMENT '是否删除0否1是',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `gmt_create` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  `aid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '添加人',
  `module` char(50) DEFAULT NULL COMMENT '模块',
  `setting` varchar(255) DEFAULT NULL COMMENT '扩展参数',
  `is_default` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否默认',
  `is_child` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否有子类1是0否',
  `is_system` tinyint(1) NOT NULL DEFAULT '0' COMMENT '系统参数禁止修改',
  `is_show` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否显示在配置页面上',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `is_del` (`is_del`),
  KEY `parent_id` (`parent_id`),
  KEY `sort` (`sort`),
  KEY `mark` (`mark`),
  KEY `type_id` (`type_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='类别';

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `mobile` char(11) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `username` char(30) DEFAULT NULL COMMENT '用户名',
  `mail` char(32) DEFAULT NULL COMMENT '邮箱',
  `password` char(32) DEFAULT NULL COMMENT '密码',
  `salt` char(6) DEFAULT NULL COMMENT '干扰码',
  `reg_ip` char(15) DEFAULT NULL COMMENT '注册IP',
  `reg_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '注册时间',
  `is_del` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '状态0正常1删除',
  `group_id` int(11) unsigned NOT NULL DEFAULT '410' COMMENT '用户组ID',
  `true_name` varchar(32) DEFAULT NULL COMMENT '真实姓名',
  `name` varchar(100) DEFAULT NULL COMMENT '店铺名称',
  `gmt_create` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  `gmt_modified` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `mobile` (`mobile`),
  KEY `is_del` (`is_del`),
  KEY `username` (`username`),
  KEY `email` (`mail`),
  KEY `group_id` (`group_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='用户表';

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` VALUES (1, '', 'admin', '', 'admin', '', '', '2018-05-15 21:20:46', 0, 0, '', '', NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for user_group
-- ----------------------------
DROP TABLE IF EXISTS `user_group`;
CREATE TABLE `user_group` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '会员用户组ID',
  `name` char(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '名称',
  `discount` int(11) NOT NULL DEFAULT '0' COMMENT '折扣率',
  `is_show_price` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '是否显示价格1是0否',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `sort` int(5) unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `is_del` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除1是0否',
  `mark` char(15) DEFAULT NULL COMMENT '标志',
  `qq` varchar(15) DEFAULT NULL COMMENT '客服',
  PRIMARY KEY (`id`),
  KEY `is_del` (`is_del`),
  KEY `sort` (`sort`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='会员用户组';

-- ----------------------------
-- Table structure for user_group_ext
-- ----------------------------
DROP TABLE IF EXISTS `user_group_ext`;
CREATE TABLE `user_group_ext` (
  `group_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  PRIMARY KEY (`group_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='会员用户组扩展';

-- ----------------------------
-- Table structure for user_profile
-- ----------------------------
DROP TABLE IF EXISTS `user_profile`;
CREATE TABLE `user_profile` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `sex` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '性别1男2女3中性0保密',
  `job` varchar(50) DEFAULT NULL COMMENT '担任职务',
  `qq` varchar(20) DEFAULT NULL,
  `phone` varchar(20) DEFAULT NULL COMMENT '电话',
  `county` int(11) unsigned NOT NULL DEFAULT '1' COMMENT '国家',
  `province` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '省',
  `city` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '市',
  `district` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '区',
  `address` varchar(255) DEFAULT NULL COMMENT '地址',
  `wechat` varchar(20) DEFAULT NULL COMMENT '微信',
  `remark_admin` text COMMENT '客服备注',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='用户其他介绍';

-- ----------------------------
-- Table structure for user_status
-- ----------------------------
DROP TABLE IF EXISTS `user_status`;
CREATE TABLE `user_status` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `reg_ip` char(15) DEFAULT NULL COMMENT '注册IP',
  `reg_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '注册时间',
  `reg_type` int(11) NOT NULL DEFAULT '0' COMMENT '注册方式',
  `reg_app_id` int(11) NOT NULL DEFAULT '1' COMMENT '注册来源',
  `last_login_ip` char(15) DEFAULT NULL COMMENT '最后登录IP',
  `last_login_time` timestamp NULL DEFAULT NULL COMMENT '最后登录时间',
  `last_login_app_id` int(11) NOT NULL DEFAULT '0' COMMENT '最后登录app_id',
  `login` smallint(5) NOT NULL DEFAULT '0' COMMENT '登录次数',
  `is_mobile` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '手机号是否已验证1已验证0未验证',
  `is_email` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '邮箱是否已验证1已验证0未验证',
  `aid_add` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '客服AID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='用户状态';

SET FOREIGN_KEY_CHECKS = 1;
