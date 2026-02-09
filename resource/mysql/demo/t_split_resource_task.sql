-- AI 应用 DEMO - utf8mb4_general_ci 不区分大小写排序规则
create database IF NOT EXISTS mysql_demo CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

use mysql_demo;

CREATE TABLE IF NOT EXISTS `t_split_resource_task` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT 'Primary key, t_audio_resource::id',
    
    `series_resource_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '资源 id',
    `language_code` varchar(16) NOT NULL default '' COMMENT '拆剧 Language code',
    
    `statemachine` SMALLINT NOT NULL default 0 COMMENT '0 无意义, 标识阶段, 随着涉及变更',
    
    `status` int NOT NULL default 0 COMMENT '默认 0 无意义, 200 认为成功, 其他认为失败',
    
    `message` varchar(1024) NOT NULL default '' COMMENT 'message 算法提示信息',
    
    `update_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update time',
    `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create time',

    UNIQUE KEY `unique_series_resource` (`series_resource_id`)
) ENGINE = InnoDB DEFAULT CHARSET = UTF8MB4 COMMENT = '拆资源任务表';
