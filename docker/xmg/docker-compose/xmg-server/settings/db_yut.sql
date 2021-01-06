drop table if exists `t_user`;
drop table if exists `t_group`;

CREATE TABLE  `t_group` (
    `id` int(10) NOT NULL AUTO_INCREMENT,
    `desc` varchar(40) NOT NULL DEFAULT '' COMMENT '描述',
    `menus` varchar(255) NOT NULL DEFAULT '' COMMENT '菜单列表',
    `auths` varchar(1024) NOT NULL DEFAULT '' COMMENT '权限url列表',
    PRIMARY KEY (`id`),
    UNIQUE KEY (`id`,`desc`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `t_user` (
  `username` varchar(40) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(33) NOT NULL DEFAULT '' COMMENT '密码',
  `email` varchar(25) NOT NULL DEFAULT '' COMMENT '邮箱',
  `group_id` int(10) NOT NULL DEFAULT 1 COMMENT '所属组ID',
  `nickname` varchar(25) NOT NULL DEFAULT '' COMMENT '昵称',
  `userdesc` varchar(100) NOT NULL DEFAULT '' COMMENT '简介',
  `wx_rcode_img` varchar(100) NOT NULL DEFAULT '' COMMENT '微信二维码',
  `zf_rcode_img` varchar(100) NOT NULL DEFAULT '' COMMENT '支付二维码',
  `sex` tinyint(1) NOT NULL DEFAULT 0 COMMENT '性别',
  PRIMARY KEY (`username`),
  FOREIGN KEY(`group_id`) REFERENCES `t_group`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

insert into t_group(`id`,`desc`,`menus`,`auths`) VALUES(1,"管理员",
"100,101,102,103,200,201,202,300,301,302,400,401,402,403,404,500,501,502,600,601",
"1001,1002,1003,1004,2001,2002,2003,2004,2005,2006,2007,2008,2009,3001,3002,3003,3004,4001,4002,4003,4004,5001,5002,5003,5004,6001,6002,6003,6004,6005,6006");


insert into t_user(`username`,`password`,`email`, `group_id`) VALUES("admin", "e10adc3949ba59abbe56e057f20f883e", "2538698032@qq.com", 1);

-- 粉丝
drop table if exists `t_friends`;
CREATE TABLE `t_user_friends` (
    `username` varchar(40) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '用户名',
    `friend_name` varchar(40) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '用户名',
    `create_tm` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建时间戳',
    PRIMARY KEY (`username`,`friend_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

drop table if exists `t_user_res`;
CREATE TABLE `t_user_res` (
    `id` int(10) NOT NULL AUTO_INCREMENT,
    `username` varchar(40) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '作者',
    `res_type` tinyint(2) NOT NULL DEFAULT 0 COMMENT 'img: 1 music:2 video:3',
    `res_desc` varchar(500) NOT NULL DEFAULT '' COMMENT '描述',
    `res_name` varchar(256) NOT NULL DEFAULT '' COMMENT '资源名称',
    `create_tm` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建时间戳',
    `update_tm` timestamp NOT NULL default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP  COMMENT '更新时间',
     PRIMARY KEY (`id`),
    KEY `idx_username_type_createtm` (`username`,`res_type`,`create_tm`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

drop table if exists `t_blog_category`;
CREATE TABLE `t_blog_category` (
    `category_id` smallint(3) NOT NULL AUTO_INCREMENT,
    `username` varchar(40) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '作者',
    `desc` varchar(256) NOT NULL DEFAULT '' COMMENT '描述',
    `update_tm` timestamp NOT NULL default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP  COMMENT '更新时间',
     PRIMARY KEY (`category_id`),
     KEY `idx_username`(`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

drop table if exists `t_blog`;
CREATE TABLE `t_blog` (
    `id` int(10) NOT NULL AUTO_INCREMENT,
    `status` tinyint(2) NOT NULL DEFAULT 0 COMMENT '状态 1： 已发布  0：未发布',
    `type` tinyint(2) NOT NULL DEFAULT 0 COMMENT '类型 1： 富文本 2：markdown',
    `username` varchar(40) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '作者',
    `name` varchar(256) NOT NULL DEFAULT '' COMMENT '博客名称',
    `url` varchar(256) NOT NULL DEFAULT '' COMMENT '博客地址',
    `category_id` smallint(3) NOT NULL DEFAULT 0 COMMENT '分类id',
    `content` text NOT NULL DEFAULT '' COMMENT '内容',
    `create_tm` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建时间戳',
    `publish_tm` bigint(20) NOT NULL DEFAULT 0 COMMENT '发布时间戳',
    `update_tm` timestamp NOT NULL default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP  COMMENT '更新时间',
     PRIMARY KEY (`id`),
     KEY `idx_category`(`category_id`),
     KEY `idx_state`(`status`),
     KEY `idx_username_createtm` (`username`,`create_tm`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;