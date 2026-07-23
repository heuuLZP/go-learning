# 第十关学习记录：完成 SQL CRUD

- 日期：2026-07-23
- 状态：已通关

## 我观察到了什么

- `INSERT INTO todos (title, done) VALUES ('my insert', false)` 写入了一条 Todo，数据库分配了 ID 3。
- `UPDATE` 可以只修改指定字段；先更新标题，再把 `done` 更新为 `true`，查询结果中的 `done` 从 `0` 变成了 `1`。
- 删除或查询后需要根据实际 ID 验证结果；自增 ID 删除后不会复用，下一条新数据会继续使用更大的 ID。

## 我的理解

- `SET` 在 `UPDATE` 中指定要更新的字段和值。
- `WHERE` 是条件判断，用来精确控制要操作哪一条数据。
- `UPDATE` 或 `DELETE` 漏掉 `WHERE` 会分别更新或删除整张表的数据，因此执行前可以先用相同条件 `SELECT` 检查命中范围。
- `SELECT` 和 `DELETE` 使用 `FROM` 指定表；`UPDATE` 在语句开头直接指定表，所以语法中没有 `FROM`。

## 术语校准

- `AUTO_INCREMENT` 保证新记录获得新的编号，不保证编号连续；删除记录后留下空号是正常的。
- `BOOLEAN` 在 MySQL 中通常以 `0` 和 `1` 存储或显示。

## 环境提示

- Mac 终端通过 `docker compose exec mysql mysql -ugo_learning -pgo_learning go_learning` 进入 MySQL 时，中文输入可能出现乱码或无法录入。练习数据尽量使用英文、数字和下划线，避免因此影响 SQL 操作。

## 当前掌握证据

- 能独立完成一轮 `INSERT`、`SELECT`、`UPDATE` 和 `DELETE`。
- 能解释 `SET`、`WHERE`、自增 ID 空号，以及 `UPDATE` 不使用 `FROM` 的原因。
