-- 为 bx_users 表添加 Google 登录相关字段
ALTER TABLE `bx_users` 
ADD COLUMN `google_id` VARCHAR(255) NULL COMMENT 'Google用户ID' AFTER `coins`,
ADD COLUMN `google_email` VARCHAR(255) NULL COMMENT 'Google邮箱' AFTER `google_id`,
ADD COLUMN `login_provider` VARCHAR(50) NOT NULL DEFAULT 'local' COMMENT '登录提供商 local/google' AFTER `google_email`;

-- 为 google_id 添加索引以提高查询性能
ALTER TABLE `bx_users` ADD INDEX `idx_bx_users_google_id` (`google_id`);

-- 为 login_provider 添加索引
ALTER TABLE `bx_users` ADD INDEX `idx_bx_users_login_provider` (`login_provider`);