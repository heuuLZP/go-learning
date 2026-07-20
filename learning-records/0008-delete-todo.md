# 第七关学习记录：删除 Todo Boss

- 日期：2026-07-20
- 状态：已通关

## 我观察到了什么

- `POST /todos` 创建了 ID 为 3 的练习 Todo，并返回 `201 Created`。
- `DELETE /todos/3` 返回 `204 No Content`，响应没有正文。
- 随后 `GET /todos/3` 返回 `404 Not Found` 和 `Todo 不存在`，证明删除已经生效。

## 我的理解

- `POST` 创建资源，`GET` 读取资源，`PUT` 更新资源，`DELETE` 删除资源。
- 当前更新接口使用路径 ID 决定目标，只允许正文改变 `title` 和 `done`，不会编辑 ID。
- 删除时先根据路径 ID 找到 slice 中的索引，再把目标左侧和右侧两段拼接起来。

## 术语校准

- slice 中的位置叫 index（索引），不叫 key。
- `204 No Content` 表示操作成功且响应没有正文。
- `todos[:index]` 不包含目标位置，`todos[index+1:]` 从目标的下一个元素开始，两段连接后目标被移除。

## 当前掌握证据

- 能完成并解释内存中的 Todo 创建、查询、更新和删除。
- 能根据操作选择 POST、GET、PUT、DELETE，并理解主要成功状态码。
- 能用删除后的 `404` 再次验证内存状态确实发生变化。

## 下一次复现

第八关将启动 Docker MySQL，让数据第一次离开 Go 进程；这一关只确认容器、MySQL 客户端和目标数据库都可用。
