# 第二关学习记录：看清一次 HTTP 响应

- 日期：2026-07-20
- 状态：已通关

## 我观察到了什么

- `GET /ping` 返回状态码 `200 OK`。
- `POST /ping` 返回状态码 `405 Method Not Allowed`，因为 `/ping` 只注册了 `GET` 方法。
- `fmt.Fprint(w, "pong")` 会让客户端在 HTTP 响应正文中看到 `pong`。
- 临时换成 `fmt.Print("pong")` 后，`pong` 出现在运行 Go 服务的终端，而不是客户端响应中。

## 我的理解

- `w` 与 HTTP 响应输出有关。
- `fmt.Fprint` 会把内容写到第一个参数指定的地方。
- `fmt.Print` 会直接写到 Go 服务进程的标准输出。

## 术语校准

- `w` 的类型是 `http.ResponseWriter`。更准确地说，它是 handler 用来写 HTTP 响应的接口，不是请求本身，也不是响应正文这个值。
- handler 可以通过 `w` 设置响应头、写入状态码和写入正文，客户端会收到这些内容。
- `fmt.Fprint` 的第一个参数是一个写入目标；把 `w` 传给它，就是把文字写入 HTTP 响应。

## 当前掌握证据

- 能区分 `200 OK` 与 `405 Method Not Allowed`。
- 能解释请求方法也是路由匹配条件的一部分。
- 能解释 `Fprint` 与 `Print` 的输出位置差异。
- 对 `ResponseWriter` 已从“见过”推进到“能解释”。

## 下一次复现

第三关不再用 `fmt.Fprint` 写固定文字，而是用 `json.NewEncoder(w)` 写 JSON。届时继续确认：虽然内容格式变了，`w` 仍然代表 HTTP 响应的写入目标。
