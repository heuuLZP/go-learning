# 第十二关学习记录：持久化 Boss

- 日期：2026-07-23
- 状态：已通关

## 我观察到了什么

- `POST /todos` 返回 `201 Created`，MySQL 为练习 Todo 分配了 ID 5。
- `PUT /todos/5` 返回 `200 OK`，标题更新为 `level12 after restart`，完成状态变为 `true`。
- 重启 Go 服务后，`GET /todos/5` 仍返回更新后的内容，证明数据没有保存在 Go 进程的内存中。
- 直接查询 MySQL 能看到相同的 ID、标题和 `done = 1`，证明 API 和数据库操作的是同一行数据。
- `DELETE /todos/5` 返回 `204 No Content`，随后查询相同 ID 返回 `404 Not Found`。

## 我的理解

- `QueryRowContext` 用于预期最多返回一行的查询；错误通常在调用 `Scan` 时返回，没有匹配行时得到 `sql.ErrNoRows`。
- `QueryContext` 用于多行查询，返回 `*sql.Rows`；需要依次检查查询错误、每行的 `Scan` 错误和遍历后的 `rows.Err()`。
- `ExecContext` 用于不返回结果行的 SQL，例如 `INSERT`、`UPDATE` 和 `DELETE`。它返回的 `sql.Result` 可以继续提供 `LastInsertId` 或 `RowsAffected`。
- 数据库结果不会自动变成 HTTP 状态。handler 根据 `error`、`sql.ErrNoRows` 和 `RowsAffected` 主动选择 `200`、`204`、`404` 或 `500`。

## 需要继续巩固

- `err` 是实现了 `error` 接口的错误值，不是一种与 `sql.ErrNoRows` 并列的具体错误类型；`sql.ErrNoRows` 是可以被识别的特定错误值。
- 对 `DELETE` 而言，`RowsAffected() == 0` 表示没有可删除的目标，可以返回 `404`。
- 对 MySQL 的 `UPDATE` 而言，影响 0 行既可能表示 ID 不存在，也可能表示新值与原值相同。因此还要按 ID 查询一次：查不到返回 `404`，查得到仍返回成功。

## 当前掌握证据

- 五条 Todo 路由都已通过 `database/sql` 读写 MySQL，内存 `todos` 和 `nextTodoID` 已移除。
- 能通过 API 完成创建、查询、更新和删除，并用 MySQL 查询交叉验证。
- 能通过重启 Go 服务证明数据持久化。
- 已达到首月 12/12 目标：HTTP 请求经过 Go handler 和 MySQL，最终形成正确的 JSON 或 HTTP 状态响应。
