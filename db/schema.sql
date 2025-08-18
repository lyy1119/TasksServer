-- 创建数据库（如果不存在）
CREATE DATABASE IF NOT EXISTS todo
  DEFAULT CHARACTER SET utf8mb4
  COLLATE utf8mb4_0900_ai_ci;
USE todo;

-- 严格模式，防止 MySQL 自动截断或容错写脏数据
SET sql_mode = 'STRICT_ALL_TABLES';

-- ==========================
-- 用户表：保存账号和角色信息
-- ==========================
CREATE TABLE users (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,              -- 主键，自增
  email VARCHAR(255) NOT NULL UNIQUE,                -- 邮箱，唯一约束
  display_name VARCHAR(100) NOT NULL,                -- 用户显示名
  password_hash VARCHAR(255) NOT NULL,               -- 密码哈希（单用户时可放占位）
  role ENUM('user','admin') NOT NULL DEFAULT 'user', -- 角色：普通用户 / 管理员
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,                           -- 创建时间
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP -- 更新时间
) ENGINE=InnoDB;

-- ====================================
-- 任务表：保存所有用户的任务（核心表）
-- ====================================
CREATE TABLE tasks (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,   -- 主键，自增

  -- 多用户预留：单用户阶段固定为 1
  user_id BIGINT NOT NULL,

  title VARCHAR(200) NOT NULL,  -- 任务标题
  description TEXT,             -- 任务描述

  -- 四象限属性：重要性和紧急性
  is_urgent    TINYINT(1) NOT NULL DEFAULT 0, -- 是否紧急
  is_important TINYINT(1) NOT NULL DEFAULT 0, -- 是否重要

  -- 任务状态
  status ENUM('todo','done','archived') NOT NULL DEFAULT 'todo', -- 待办 / 已完成 / 已归档

  -- 时间字段（存储 UTC）
  due_at    DATETIME NULL,  -- 截止时间
  remind_at DATETIME NULL,  -- 提醒时间

  -- 循环规则（使用 iCal RRULE 格式，例如 FREQ=WEEKLY;BYDAY=MO,WE,FR）
  rrule           VARCHAR(300) NULL,
  next_occurrence DATETIME NULL,  -- 下一次循环发生时间（由后端计算保存，加速查询）

  -- 生成列：自动计算四象限编号
  -- Q1=1：重要且紧急；Q2=2：重要不紧急；Q3=3：不重要但紧急；Q4=4：不重要不紧急
  quadrant TINYINT AS (
    CASE
      WHEN is_important=1 AND is_urgent=1 THEN 1
      WHEN is_important=1 AND is_urgent=0 THEN 2
      WHEN is_important=0 AND is_urgent=1 THEN 3
      ELSE 4
    END
  ) STORED,

  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,                           -- 创建时间
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- 更新时间

  -- 外键：删除用户时，自动删除其任务
  CONSTRAINT fk_tasks_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,

  -- 常用查询索引
  INDEX idx_tasks_user_status_due (user_id, status, due_at),       -- 用户 + 状态 + 截止时间（列表/过期筛选）
  INDEX idx_tasks_user_quadrant (user_id, quadrant),               -- 用户 + 四象限（四象限视图）
  INDEX idx_tasks_user_nextocc (user_id, next_occurrence),         -- 用户 + 下一次循环时间
  INDEX idx_tasks_user_remind (user_id, remind_at),                -- 用户 + 提醒时间
  FULLTEXT INDEX ftx_tasks_title_desc (title, description)         -- 全文索引：标题/描述搜索
) ENGINE=InnoDB;

-- =====================
-- 插入一个初始用户（单用户用）
-- =====================
INSERT INTO users (email, display_name, password_hash, role)
VALUES ('single@local', 'Single User', '***placeholder***', 'user')
ON DUPLICATE KEY UPDATE email=email;
