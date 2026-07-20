# 第五关学习记录：按 ID 查询 Todo

- 日期：2026-07-20
- 状态：已通关

## 我观察到了什么

- `GET /todos/1` 返回 `200 OK` 和 ID 为 1 的单个 Todo。
- `GET /todos/10` 返回 `404 Not Found` 和正文 `Todo 不存在`。
- 查询单个 Todo 时响应是 JSON 对象，不是 JSON 数组。

## 我的理解

- `r.PathValue("id")` 从匹配 `/todos/{id}` 的 URL 路径中取出 `id`。
- `strconv.Atoi` 把路径中的 `string` 转换成 `int`，以便和 `Todo.ID` 比较。
- `range` 依次遍历 `todos` 中的元素。
- `404 Not Found` 是 HTTP 响应状态，表示没有找到对应资源。

## 术语校准

- `PathValue` 读取的是路径变量，不是 query 参数。
- `/todos/1` 中的 `1` 是路径变量；`/todos?id=1` 中的 `id=1` 才是 query 参数。
- `404` 表示 ID 格式有效但资源不存在；无法转换成数字的 ID 返回 `400 Bad Request`。

## 当前掌握证据

- 能改变 URL 中的 ID，查询不同的 Todo。
- 能用实际响应区分存在资源的 `200` 与不存在资源的 `404`。
- 能解释路径文字如何转换成数字，再与内存数据比较。

## 下一次复现

第六关将组合已经学过的路径 ID 和请求 JSON，用 `PUT /todos/{id}` 替换一个已有 Todo。
