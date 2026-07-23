# 首月目标完成：12 / 12

第十二关已经通关。Go Todo API 的五条路由都已接入 MySQL，重启 Go 服务后数据仍然存在，首月后端主线完成。

## 已获得的能力

- 能说明 HTTP 请求如何经过 Go handler、`database/sql` 和 MySQL，再形成 JSON 或状态码响应
- 能分别使用 REST API 和 SQL 完成 Todo CRUD
- 能区分单行查询、多行查询和不返回结果行的 SQL 操作
- 能把数据库的无结果、执行失败和受影响行数翻译成 `404`、`500` 或成功响应
- 能用 API 与 MySQL 查询交叉验证数据，并通过重启 Go 服务证明持久化

## 本次需要记牢的一条链路

```text
数据库操作
  -> 检查 error
  -> 读取 Scan / LastInsertId / RowsAffected
  -> handler 判断结果
  -> 写入 HTTP 状态、响应头和 JSON
```

数据库结果不会自动决定 HTTP 状态，handler 负责翻译。尤其是 MySQL 的 `UPDATE` 影响 0 行时，还要区分“ID 不存在”和“值本来就相同”。

## 完成记录

查看 [第十二关学习记录](learning-records/0012-persistence-boss.md) 和 [12 / 12 闯关地图](roadmap.html)。

想复习实现，可以重新打开 [第十二关课程](lessons/0012-persistence-boss.html)。

## 下一步

先停在这个可运行成果上，不自动加入事务、ORM、鉴权或项目分层。下一阶段应从真实后端需求中选择一个新目标，再决定需要解锁哪项能力。
