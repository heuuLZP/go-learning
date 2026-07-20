# 当前关：启动 MySQL

预计用时：15–20 分钟；首次下载镜像可能额外等待。今天只做这一关。

第七关已经通关，内存 REST CRUD 完成。第八关让数据第一次离开 Go 进程：启动 MySQL，进入客户端，并确认目标数据库存在。

## 通关条件

- 运行 `go run .`
- 用 `docker compose up -d` 启动 MySQL
- 用 `docker compose ps` 看到 MySQL 状态为 `healthy`
- 进入 MySQL 客户端，`SELECT DATABASE()` 返回 `go_learning`
- 能区分 Docker Compose、MySQL 容器和 `go_learning` 数据库

## 开始

打开 [第八关课程](lessons/0008-mysql-docker.html)。

想先看整体方向，可以打开 [Go 后端闯关地图](roadmap.html)。

如果今天只有 2 分钟：只运行 `docker compose up -d` 和 `docker compose ps`，看到容器开始运行，也算保温成功；下一次仍从本关继续。

## 向老师通关

完成后回复两样东西：

1. `docker compose ps`、`SELECT DATABASE()` 和自定义文字查询的结果；
2. 你怎么区分 Compose、MySQL 容器与 `go_learning` 数据库。

老师确认后才解锁下一关。
