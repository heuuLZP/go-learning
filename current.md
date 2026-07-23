# 当前关：Go 连接 MySQL

预计用时：25–30 分钟。今天只做这一关。

第十关已经通关，能够独立完成一轮 SQL CRUD。第十一关把 Go 服务接到同一张 `todos` 表，先验证连接，再让列表查询和创建走数据库。

## 通关条件

- 用 `docker compose up -d` 启动 MySQL
- 安装 `go-sql-driver/mysql`，并用 `database/sql` 打开连接
- 用 `db.Ping()` 验证 Go 确实能到达 MySQL
- 把 `GET /todos` 改为使用 `QueryContext` 和 `Scan` 查询数据库
- 把 `POST /todos` 改为使用参数化 `INSERT`，返回数据库生成的 ID
- 能解释驱动、`sql.Open`、`Ping`、`QueryContext` 和 `Scan` 的分工

## 开始

打开 [第十一关课程](lessons/0011-go-mysql.html)。

想先看整体方向，可以打开 [Go 后端闯关地图](roadmap.html)。

如果今天只有 2 分钟：启动 MySQL，在 `main.go` 保存驱动导入后执行 `go mod tidy`，再读一遍 DSN 和 `db.Ping()`，也算保温成功；下一次仍从本关继续。

## 向老师通关

完成后回复两样东西：

1. Go 启动时的连接结果、`GET /todos`、`POST /todos` 和 MySQL 中对应的查询结果；
2. 你怎么解释驱动、`sql.Open`、`Ping`、`QueryContext` 和 `Scan`。

老师确认后才解锁下一关。
