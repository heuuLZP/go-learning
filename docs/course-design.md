# 第一个月教师用关卡设计

> 这是课程后台，不是每日待办。学习时只打开 [`current.md`](../current.md)。

## 首月成果

用 Go 完成一条最小真实后端链路：

```text
HTTP 请求 -> REST handler -> Todo -> SQL -> MySQL -> JSON 响应
```

课程按每周三次设计，共 12 个微关，预计主动学习 3–5 小时。Go、Docker 和数据库样板由 AI 主写；学习者负责运行、观察、亲自改一处和解释关键链路。

## 通关协议

每个微关只包含四步：

1. **回忆**：不看答案回忆上关的一点，最多 1 分钟。
2. **行动**：运行或修改一个很小的代码点，约 5–12 分钟。
3. **验证**：用命令、浏览器、`curl` 或测试立刻看到结果。
4. **复述**：用一句自己的话解释刚才发生了什么。

普通关 10–20 分钟，数据库 Boss 关最多 30 分钟。超过上限就拆成两关；中断后从当前关继续，不补课。

## 分级达成

- **8/12 基础通关**：能运行和修改内存版 REST CRUD。
- **10/12 良好通关**：能启动 MySQL，并亲自执行 SQL CRUD。
- **12/12 目标通关**：Go API 使用 MySQL，重启服务后数据仍存在。

任一档都算首月完成。达成度描述能力，不作为排名或追赶压力。

## 首月关卡池

### 阶段一：HTTP 与 JSON（3 关）

1. 启动 `GET /hello`，修改响应并解释 handler。
2. 新增 `GET /ping`，认识路由、方法和 `200 OK`。
3. 用 `struct` 和 JSON 返回一个及多个 Todo。

### 阶段二：REST 与内存 CRUD（4 关）

4. `POST /todos`：读取 JSON，创建 Todo，返回 `201 Created`。
5. `GET /todos/{id}`：读取路径 ID，区分 `200` 与 `404`。
6. `PUT /todos/{id}`：完整更新 `title` 和 `done`，非法输入返回 `400`。
7. `DELETE /todos/{id}`：删除后返回 `204 No Content`；Boss 复述 REST CRUD。

### 阶段三：SQL 与 MySQL（3 关）

8. 用 Docker Compose 启动 MySQL，进入客户端并看到目标数据库。
9. 建立 `todos` 表，亲自执行 `INSERT` 和 `SELECT`。
10. 亲自执行 `UPDATE` 和 `DELETE`；不看笔记完成一次 SQL CRUD。

### 阶段四：Go 接入 MySQL（2 关）

11. 用 `database/sql` 连接 MySQL，让创建和查询真正持久化。
12. 完成更新和删除；Boss 重启 Go 服务并证明数据仍然存在。

## 固定 API 契约

| 方法与路径 | 行为 | 成功状态 |
|---|---|---:|
| `GET /todos` | 查询列表 | `200` |
| `GET /todos/{id}` | 查询单项 | `200` |
| `POST /todos` | 创建 Todo | `201` |
| `PUT /todos/{id}` | 完整更新 Todo | `200` |
| `DELETE /todos/{id}` | 删除 Todo | `204` |

Todo 固定包含 `id`、`title`、`done`、`created_at`、`updated_at`。非法输入返回 `400`，不存在返回 `404`，未知内部错误返回不泄露细节的 `500`。

## 技术边界

- 首月使用 `net/http`、`database/sql`、`go-sql-driver/mysql` 和 Docker MySQL。
- 不引入 Gin、ORM、Prisma、Redis 或前端页面。
- 测试作为验证命令保留，系统学习测试顺延。
- MySQL 配置只在第 8 关解锁时生成，不提前打扰当前学习。

## 动态规则

- 连续卡住 5 分钟就缩小动作；同一关两次需要大量提示时插入同构练习。
- 一次通过仍只解锁下一关，可选挑战默认折叠。
- 每次只更新 `current.md`、当前 HTML 课程和路线图状态，不一次生成 12 堂课。
- 事务、索引、鉴权和 Redis 的解锁条件见 [`backend-capabilities.md`](backend-capabilities.md)。
