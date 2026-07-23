# 第十一关学习记录：Go 连接 MySQL

- 日期：2026-07-23
- 状态：已通关

## 我观察到了什么

- MySQL 中原有一条 ID 为 1 的 Todo，`GET /todos` 返回了相同的 ID、标题和完成状态，证明列表接口已经从数据库读取数据。
- `POST /todos` 返回 `201 Created`，MySQL 为新 Todo 分配了 ID 4；随后直接查询 `todos` 表也能看到同一条数据。
- 自增 ID 从 1 跳到 4 是正常现象：被删除或执行失败的历史写入可能留下空号，`AUTO_INCREMENT` 不保证编号连续。
- MySQL 客户端提示命令行密码可能不安全；当前凭据仅用于本地学习环境，生产环境不能把密码直接写在命令或源码中。

## 我的理解

- `database/sql` 是 Go 标准库提供的数据库操作接口，具体的 MySQL 协议由 `go-sql-driver/mysql` 驱动实现。
- `sql.Open` 创建 `*sql.DB` 连接池句柄；`Ping` 才会实际验证 Go 到 MySQL 的连接。
- `QueryContext` 使用 SQL 执行查询，并把 HTTP 请求的取消或超时信号传入数据库操作。
- `rows.Next` 移动到当前结果行，`Scan` 再把这一行的列写入传入的字段地址。
- `ExecContext` 执行不返回结果行的写操作；`LastInsertId` 取得本次 `INSERT` 生成的自增 ID。

## 延伸了解

- 真实项目可能按 controller、service、repository 或 DAO 分层，让 HTTP、业务规则和数据访问各自承担不同职责。
- Go 也有 ORM，例如 GORM、Ent 和 Bun；本课程当前使用 `database/sql` 与原生 SQL，先建立查询、参数和持久化的基础心智。
- 分层、ORM、连接池调优与生产环境配置目前只作概念了解，不在本关展开。

## 当前掌握证据

- 能通过 `GET /todos` 读取 MySQL 数据。
- 能通过 `POST /todos` 创建数据，并在 MySQL 客户端中交叉验证。
- 能区分 `database/sql` 与 MySQL 驱动，并解释 `Open`、`Ping`、`QueryContext`、`Scan`、`ExecContext` 和 `LastInsertId` 的作用。
