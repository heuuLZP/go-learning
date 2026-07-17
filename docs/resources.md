# Go 后端入门资源

## Knowledge

- [Tutorial: Getting started — Go 官方](https://go.dev/doc/tutorial/getting-started)
  用于第一次运行 Go 程序和理解 module；只按关卡引用，不要求通读。
- [Tutorial: Create a Go module — Go 官方](https://go.dev/doc/tutorial/create-module)
  用于函数、包和错误的最小示例；需要拆包时再读。
- [net/http package — Go 官方](https://pkg.go.dev/net/http)
  HTTP 服务行为的权威参考；前几关只查看 `HandleFunc`、`ListenAndServe` 和 `ResponseWriter`。
- [encoding/json package — Go 官方](https://pkg.go.dev/encoding/json)
  JSON 编解码的权威参考；进入 Todo JSON 关卡后使用。
- [HTTP request methods — MDN](https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Methods)
  HTTP 方法的可靠索引；进入 REST CRUD 时只查 `GET`、`POST`、`PUT` 和 `DELETE`。
- [HTTP response status codes — MDN](https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Status)
  状态码语义索引；用于区分 `200`、`201`、`204`、`400`、`404` 和 `500`。
- [Accessing relational databases — Go 官方](https://go.dev/doc/database/)
  Go 数据库入门主资料；进入 MySQL 接入关后学习连接、查询和参数传递。
- [database/sql package — Go 官方](https://pkg.go.dev/database/sql)
  Go 标准数据库接口的权威参考；只在课程遇到具体 API 时查阅。
- [MySQL Tutorial — MySQL 官方](https://dev.mysql.com/doc/refman/8.4/en/tutorial.html)
  建表与 SQL CRUD 的主资料；第 8–10 关按需引用，不要求通读。
- [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)
  Go 连接 MySQL 的驱动文档；第 11 关用于 DSN、驱动注册和连接配置。
- [A Tour of Go — Go 官方](https://go.dev/tour/)
  按需查语法和类型，不作为必须从头刷完的课程。
- [jiangtao/go-together](https://github.com/jiangtao/go-together)
  课程结构参考。适合有 Node.js/SQL 后端经验的学习者；本工作区只复用其优秀教学机制，不照搬难度和顺序。

## Later Knowledge

- [MySQL InnoDB transaction model — MySQL 官方](https://dev.mysql.com/doc/refman/8.4/en/innodb-transaction-model.html)
  多表原子操作出现后用于事务与隔离概念；首月不展开。
- [How MySQL uses indexes — MySQL 官方](https://dev.mysql.com/doc/refman/8.4/en/mysql-indexes.html)
  出现稳定查询条件后配合 `EXPLAIN` 学习索引；首月不做性能调优。
- [Authentication Cheat Sheet — OWASP](https://cheatsheetseries.owasp.org/cheatsheets/Authentication_Cheat_Sheet.html)
  引入用户后学习认证安全边界；不用于首月 Todo API。
- [Redis data types — Redis 官方](https://redis.io/docs/latest/develop/data-types/)
  出现会话 TTL、缓存或限流需求后选择数据结构；不把 Redis 当作默认依赖。

## Wisdom (Communities)

- 暂不要求加入社区。纯新手阶段优先保持短反馈回路；出现真实、反复的问题后再选择社区求证。
