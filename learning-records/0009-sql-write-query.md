# 第九关学习记录：写入与查询

- 日期：2026-07-22
- 状态：已通关

## 我观察到了什么

- `DESCRIBE todos` 显示了 `id`、`title`、`done` 三个字段及其类型、主键、默认值和空值限制。
- `INSERT INTO todos (title)` 写入一条 Todo，`SELECT` 查询到 ID 为 1、`done` 为 0 的数据。
- MySQL 容器重启后再次执行 `SELECT`，之前写入的数据仍然存在。

## 我的理解

- 表结构有点像 TypeScript 的类型或 interface：它描述一行数据有哪些字段，并约束字段类型、主键、默认值和是否允许空值。
- `INSERT` 负责向表中写入数据，`SELECT` 负责查询符合条件的数据。
- Docker volume 保存 MySQL 数据，使容器重启后仍能读取之前的数据。

## 术语校准

- `IF NOT EXISTS` 表示表已经存在时跳过创建操作，并且不报错。
- MySQL 中的 `BOOLEAN` 是 `TINYINT(1)` 的别名，通常用 `0` 表示 false、`1` 表示 true；客户端可以选择不同的展示方式。
- `AUTO_INCREMENT` 让数据库自动生成 `id`，`DEFAULT FALSE` 让未指定的 `done` 使用默认值。
- SQL 字符串引号中的前导空格也属于数据，会被原样保存。

## 当前掌握证据

- 能根据 `DESCRIBE` 结果解释表的字段和主要约束。
- 能执行 `INSERT` 和 `SELECT`，并解释它们分别对应写入与查询。
- 能用容器重启后的查询结果验证 volume 中的数据仍然存在。

## 下一次复现

第十关会先练习同时写入 `title` 和 `done`，再用带 `WHERE` 的 `UPDATE` 与 `DELETE` 完成一轮 SQL CRUD。语句会先拆成动作、表、内容和条件，降低长 SQL 的记忆负担。
