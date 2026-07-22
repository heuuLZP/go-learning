# 当前关：完成 SQL CRUD

预计用时：20–25 分钟。今天只做这一关。

第九关已经通关，能够创建表、写入并查询数据，也验证了 Docker volume 的持久化。第十关同时写入两个字段，再用 `UPDATE` 和 `DELETE` 完成一轮 SQL CRUD。

## 通关条件

- 用 `docker compose up -d` 启动 MySQL
- 用一条 `INSERT` 同时写入自定义 `title` 和 `done`
- 用带 `WHERE` 的 `UPDATE` 更新目标 Todo，再查询验证
- 用带 `WHERE` 的 `DELETE` 删除目标 Todo，再查询得到 `Empty set`
- 能解释 `SET`、`WHERE`，以及更新或删除时漏掉 `WHERE` 的风险
- 遮住完整示例后，借助语句骨架独立完成一次 SQL CRUD

## 开始

打开 [第十关课程](lessons/0010-sql-update-delete.html)。

想先看整体方向，可以打开 [Go 后端闯关地图](roadmap.html)。

如果今天只有 2 分钟：用一条 `INSERT` 同时写入 `title` 和 `done`，再用 `SELECT` 查出来，也算保温成功；下一次仍从本关继续。

## 向老师通关

完成后回复两样东西：

1. 自定义 Todo 的 `INSERT`、更新后的查询结果、`DELETE` 和删除后的查询结果；
2. 你怎么解释 `SET`、`WHERE`，以及漏掉 `WHERE` 的风险。

老师确认后才解锁下一关。
