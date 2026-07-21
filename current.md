# 当前关：写入与查询

预计用时：15–20 分钟；首次下载镜像可能额外等待。今天只做这一关。

第八关已经通关，MySQL 和目标数据库已经准备好。第九关创建第一张 `todos` 表，写入一条 Todo，再查询并验证数据能够跨容器重启保存。

## 通关条件

- 用 `docker compose up -d` 启动 MySQL
- 用 `docker compose ps` 看到 MySQL 状态为 `healthy`
- 进入 MySQL 客户端，`SELECT DATABASE()` 返回 `go_learning`
- 创建 `todos` 表，并用 `DESCRIBE todos` 看到 `id`、`title`、`done`
- 插入一条自己命名的 Todo，再用 `SELECT` 查询出来
- 重启 MySQL 容器后，之前写入的 Todo 仍然存在

## 开始

打开 [第九关课程](lessons/0009-sql-write-query.html)。

想先看整体方向，可以打开 [Go 后端闯关地图](roadmap.html)。

如果今天只有 2 分钟：进入 MySQL 后执行 `SELECT DATABASE()` 和 `DESCRIBE todos`，能看到目标数据库和表结构，也算保温成功；下一次仍从本关继续。

## 向老师通关

完成后回复两样东西：

1. `DESCRIBE todos`、插入自定义 Todo 后的 `SELECT` 结果，以及重启容器后的查询结果；
2. 你怎么解释 `CREATE TABLE`、`INSERT`、`SELECT` 和 Docker volume 的关系。

老师确认后才解锁下一关。
