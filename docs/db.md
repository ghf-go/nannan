# 数据库

```mysql
-- Create syntax for TABLE 'tb_comment_feed'
CREATE TABLE `tb_comment_feed` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `feed_title` varchar(100) NOT NULL DEFAULT '' COMMENT '标题',
  `feed_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '类型',
  `feed_desc` varchar(500) NOT NULL DEFAULT '' COMMENT '描述',
  `feed_imgs` varchar(1024) NOT NULL DEFAULT '' COMMENT '图片',
  `feed_replys` int(11) NOT NULL DEFAULT '0' COMMENT '回复次数',
  `feed_praise` int(11) NOT NULL DEFAULT '0' COMMENT '点赞次数',
  `x` float NOT NULL DEFAULT '0' COMMENT '坐标',
  `y` float NOT NULL DEFAULT '0' COMMENT '坐标',
  `city` varchar(50) NOT NULL DEFAULT '' COMMENT '城市',
  `ext` varchar(1024) NOT NULL DEFAULT '' COMMENT '扩展信息',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0待审核，200审核成功，100删除',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_city` (`city`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Feed列表';

-- Create syntax for TABLE 'tb_comment_feed_content'
CREATE TABLE `tb_comment_feed_content` (
  `id` bigint(20) NOT NULL COMMENT '主键ID',
  `content` text NOT NULL COMMENT '详细内容',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='feed详情';

-- Create syntax for TABLE 'tb_comment_praise'
CREATE TABLE `tb_comment_praise` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `target_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '目标类型，',
  `target_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '被关注的目标ID',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_type_id_user` (`target_type`,`target_id`,`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='点赞信息表';

-- Create syntax for TABLE 'tb_comment_reply'
CREATE TABLE `tb_comment_reply` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `target_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '目标类型，',
  `target_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '被关注的目标ID',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `parent_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '评论的父节点',
  `content` varchar(4096) NOT NULL DEFAULT '' COMMENT '评论内容',
  `reply_count` int(11) NOT NULL DEFAULT '0' COMMENT '被评论次数',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态，0待审核，100 审核完毕，50删除',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_type_id` (`target_type`,`target_id`,`parent_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户评论';

-- Create syntax for TABLE 'tb_comment_score'
CREATE TABLE `tb_comment_score` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `target_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '目标类型，',
  `target_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '被关注的目标ID',
  `score` tinyint(4) NOT NULL DEFAULT '0' COMMENT '得分',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_type_id_user` (`target_type`,`target_id`,`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户打分';

-- Create syntax for TABLE 'tb_common_upload'
CREATE TABLE `tb_common_upload` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `file_key` varchar(32) NOT NULL DEFAULT '' COMMENT '文件key',
  `file_name` varchar(100) NOT NULL DEFAULT '' COMMENT '文件名字',
  `file_size` int(11) NOT NULL DEFAULT '0' COMMENT '文件大小',
  `path` varchar(100) NOT NULL DEFAULT '' COMMENT '文件路径',
  `file_mine` varchar(30) NOT NULL DEFAULT '' COMMENT '文件类型',
  `audit_time` int(11) NOT NULL DEFAULT '0' COMMENT '审核时间',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态：0带审核，100正常，50删除',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_file_key` (`file_key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='上传文件';

-- Create syntax for TABLE 'tb_order'
CREATE TABLE `tb_order` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `type_type` int(11) NOT NULL DEFAULT '0' COMMENT '订单类型自定义使用',
  `order_title` varchar(100) NOT NULL DEFAULT '' COMMENT '订单名字',
  `amount` int(11) NOT NULL DEFAULT '0' COMMENT '订单价格，单位分',
  `platform` varchar(10) NOT NULL DEFAULT '' COMMENT '支付平台，微信，支付宝',
  `platform_order_id` varchar(70) NOT NULL DEFAULT '' COMMENT '第三方支付订单',
  `platform_type` varchar(10) NOT NULL DEFAULT '' COMMENT '支付方式，app，pc,h5,小程序等',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态，0待支付，100 支付成功，50订单超时，110退款申请，120退款成功',
  `apply_refund_msg` varchar(100) NOT NULL DEFAULT '' COMMENT '申请退款原因',
  `apply_refund_at` datetime DEFAULT NULL COMMENT '申请退款时间',
  `pay_success_at` datetime DEFAULT NULL COMMENT '支付成功时间',
  `refund_success_at` date DEFAULT NULL COMMENT '退款成功时间',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单系统';

-- Create syntax for TABLE 'tb_relation_blacklist'
CREATE TABLE `tb_relation_blacklist` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `target_user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '目标用ID',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_user_id_target_user_id` (`user_id`,`target_user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='用户黑名单';

-- Create syntax for TABLE 'tb_relation_follow'
CREATE TABLE `tb_relation_follow` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `target_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '关注类型，0关注人',
  `target_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '被关注的目标ID',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_type_id_user` (`target_type`,`target_id`,`user_id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='用户关注';

-- Create syntax for TABLE 'tb_relation_friends'
CREATE TABLE `tb_relation_friends` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `target_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '好友用户ID',
  `mark_name` varchar(20) NOT NULL DEFAULT '' COMMENT '好友备注名字',
  `apply_msg` varchar(30) NOT NULL DEFAULT '' COMMENT '申请添加好友消息',
  `audit_msg` varchar(30) NOT NULL DEFAULT '' COMMENT '验证好友消息',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态：0申请中，100好友，10好友申请被拒绝，50好友添加申请，60拒绝好友申请',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COMMENT='好友列表';

-- Create syntax for TABLE 'tb_system_conf'
CREATE TABLE `tb_system_conf` (
  `conf_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `group_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '分类id',
  `key` varchar(30) NOT NULL DEFAULT '' COMMENT '设置的KEY值',
  `desc` varchar(30) NOT NULL DEFAULT '' COMMENT '描述',
  `val_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '目标值类型',
  `val` varchar(100) NOT NULL DEFAULT '' COMMENT '值',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
  PRIMARY KEY (`conf_id`),
  UNIQUE KEY `uniq_group_id_key` (`group_id`,`key`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='系统kv配置';

-- Create syntax for TABLE 'tb_system_group'
CREATE TABLE `tb_system_group` (
  `group_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `group_name` varchar(100) NOT NULL DEFAULT '' COMMENT '分类名字',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
  PRIMARY KEY (`group_id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COMMENT='系统分类';

-- Create syntax for TABLE 'tb_system_tags'
CREATE TABLE `tb_system_tags` (
  `tag_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `group_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '分类id',
  `tag_name` varchar(100) NOT NULL DEFAULT '' COMMENT '标签名字',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
  PRIMARY KEY (`tag_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COMMENT='标签表';

-- Create syntax for TABLE 'tb_user_base_info'
CREATE TABLE `tb_user_base_info` (
  `user_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `nick_name` varchar(30) NOT NULL DEFAULT '' COMMENT '昵称',
  `avatar` varchar(120) NOT NULL DEFAULT '' COMMENT '头像',
  `gender` tinyint(4) NOT NULL DEFAULT '0' COMMENT '性别：1男，2女',
  `birthday` date DEFAULT NULL COMMENT '生日',
  `province` varchar(50) NOT NULL DEFAULT '' COMMENT '省份',
  `city` varchar(50) NOT NULL DEFAULT '' COMMENT '城市',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='账号基本信息';

-- Create syntax for TABLE 'tb_user_email'
CREATE TABLE `tb_user_email` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `email` varchar(120) NOT NULL DEFAULT '' COMMENT '邮箱',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_email` (`email`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='用户邮箱绑定';

-- Create syntax for TABLE 'tb_user_id'
CREATE TABLE `tb_user_id` (
  `user_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态：1正常，10 禁止登录',
  `role` tinyint(4) NOT NULL DEFAULT '1' COMMENT '账号类型：1一般用户，200，管理员，0超级管理员',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COMMENT='用户ID发号器';

-- Create syntax for TABLE 'tb_user_mobile'
CREATE TABLE `tb_user_mobile` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `mobile` varchar(15) NOT NULL DEFAULT '' COMMENT '手机号',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_mobile` (`mobile`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COMMENT='用户手机号绑定';

-- Create syntax for TABLE 'tb_user_pass'
CREATE TABLE `tb_user_pass` (
  `user_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `passwd` varchar(32) NOT NULL DEFAULT '' COMMENT '密码',
  `sign` varchar(10) NOT NULL DEFAULT '' COMMENT '密码加盐',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户ID发号器';

-- Create syntax for TABLE 'tb_user_profile'
CREATE TABLE `tb_user_profile` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `key` varchar(30) NOT NULL DEFAULT '' COMMENT '关键字',
  `val` varchar(100) NOT NULL DEFAULT '' COMMENT '值',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_user_id_key` (`user_id`,`key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户ID发号器';

-- Create syntax for TABLE 'tb_user_wx'
CREATE TABLE `tb_user_wx` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `wx_openid` varchar(100) NOT NULL DEFAULT '',
  `wx_unid` varchar(100) NOT NULL DEFAULT '',
  `wx_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '类型：1服务号，2小程序，3微信',
  `wx_nickname` varchar(50) NOT NULL DEFAULT '' COMMENT '微信昵称',
  `wx_avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '微信头像',
  `wx_mobile` varchar(15) NOT NULL DEFAULT '' COMMENT '微信绑定的手机号',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_wx_openid_wx_type` (`wx_type`,`wx_openid`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='微信登录信息';
```