# 第三关学习记录：让 Todo 变成 JSON

- 日期：2026-07-20
- 状态：已通关

## 我观察到了什么

- `GET /todos` 返回 JSON 数组。
- 在 `todos` 中增加第二项后，响应正文也出现了两个 JSON 对象。
- 第二项是 `{"id":2,"title":"增加第二项","done":false}`。

## 我的理解

- Go 中的 `Todo` 是一个 `struct` 类型。
- `todos` 是一个 `[]Todo`，其中每个元素都是 `Todo` 类型。
- `encoding/json` 把 `todos` 编码成 JSON。
- `json.NewEncoder(w).Encode(todos)` 把编码结果写入 HTTP 响应，客户端因此能看到 JSON 正文。

## 术语校准

- `[]Todo` 更准确地叫 slice（切片），不是数组；这一关只把它理解成 Todo 列表即可。
- `json.NewEncoder(w)` 创建的是一个以 `w` 为写入目标的编码器；真正执行“编码并写入”的是后面的 `Encode(todos)`。
- `Todo` 字段后的 JSON 标签决定输出字段名，例如 `Title` 对应 `title`。

## 当前掌握证据

- 能亲自增加一个 `Todo`，并在响应中找到对应 JSON 对象。
- 能解释 `Todo`、`[]Todo`、JSON 数组和 HTTP 响应之间的关系。
- 能继续追踪 `w`：无论写固定文字还是 JSON，它都指向 HTTP 响应。

## 下一次复现

第四关将数据流反过来：客户端发送 JSON，`json.NewDecoder(r.Body)` 从 HTTP 请求正文读取它，再创建一个新的 Todo。
