# 第四关学习记录：用 POST 创建 Todo

- 日期：2026-07-20
- 状态：已通关

## 我观察到了什么

- `POST /todos` 返回 `201 Created` 和新创建的 Todo JSON。
- 把请求字段 `title` 写成 `Title` 仍能解码，因为 `encoding/json` 匹配字段名时不区分大小写。
- 删除 `done` 但保持 JSON 合法时仍返回 `201`，`Done` 使用 Go `bool` 的零值 `false`。
- 删除字段后留下尾随逗号会返回 `400 Bad Request`，因为正文不再是合法 JSON。
- 请求正文为 `{}` 时仍返回 `201`，说明当前代码只检查 JSON 语法，没有检查必填字段。

## 我的理解

- `r.Body` 是 HTTP 请求正文，客户端发送的 JSON 位于其中。
- `Decode(&todo)` 从请求正文读取并解析 JSON，再把结果放进 `todo`。
- `append(todos, todo)` 把新 Todo 追加到 `todos` slice。
- `201 Created` 是 HTTP 响应状态码，表示请求成功创建了一个资源。

## 术语校准

- `201 Created` 位于 HTTP 响应状态行，不是响应头字段。
- 缺少 JSON 字段与 JSON 语法错误不同：前者仍可解码，后者会让 `Decode` 返回错误。
- `Decode` 不会自动执行“标题不能为空”这样的业务校验；严格字段校验将在后续真实需求中加入。

## 当前掌握证据

- 能修改请求 JSON 并观察新 Todo 被追加到内存列表。
- 能解释请求 JSON 从 `r.Body` 到 `todo`，再进入 `todos` 的完整方向。
- 能通过尾随逗号、缺少字段和空对象实验区分 JSON 解析与业务校验。

## 下一次复现

第五关从 URL 路径读取 Todo ID：存在时返回 `200 OK` 和单个 Todo，不存在时返回 `404 Not Found`。
