-- Add Google login related fields to the bx_users table
ALTER TABLE `bx_users`
ADD COLUMN `google_id` VARCHAR(255) NULL COMMENT 'Google User ID' AFTER `coins`,
ADD COLUMN `google_email` VARCHAR(255) NULL COMMENT 'Google Email' AFTER `google_id`,
ADD COLUMN `login_provider` VARCHAR(50) NOT NULL DEFAULT 'local' COMMENT 'Login Provider local/google' AFTER `google_email`;

-- Add index to google_id for improved query performance
ALTER TABLE `bx_users` ADD INDEX `idx_bx_users_google_id` (`google_id`);

-- Add index to login_provider
ALTER TABLE `bx_users` ADD INDEX `idx_bx_users_login_provider` (`login_provider`);
