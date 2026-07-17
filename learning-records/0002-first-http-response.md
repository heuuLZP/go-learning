# 第一关学习记录：第一次 HTTP 响应

- 日期：2026-07-17
- 状态：已通关
- 最终响应：`Hello, Go First`

## 我完成了什么

- 启动了 Go HTTP 服务并访问 `GET /hello`。
- 找到并修改了 `helloHandler` 中的固定响应文字。
- 重启服务后确认浏览器显示 `Hello, Go First`。

## 我的理解

- `http.HandleFunc("GET /hello", helloHandler)` 定义了请求方法、路径和负责处理请求的函数。
- 请求匹配 `GET /hello` 后会执行 `helloHandler`。
- `helloHandler` 负责产生这次请求的具体响应，目前写回的是固定文字。

补充校准：`HandleFunc` 是 `net/http` 包提供的函数；`"GET /hello"` 同时包含请求方法和路径。

## 我的疑问

- `fmt.Fprint` 和普通的 `fmt.Print` 有什么区别？
- `helloHandler` 参数中的 `w` 具体起什么作用？
- 响应除了正文之外，还包含哪些信息？

## 当前整理

- `fmt.Print(...)` 默认把内容写到程序的标准输出，通常是在终端中看到。
- `fmt.Fprint(w, ...)` 把内容写到指定目标；handler 中的 `w` 代表 HTTP 响应，所以浏览器或客户端能够收到内容。
- `w` 的完整类型是 `http.ResponseWriter`，可以用来设置响应头、状态码和正文。

这些内容已经有初步解释，但还需要通过第二关的实际请求和对比实验确认理解。

## 下一步

第二关通过 `GET /ping` 观察 `200 OK`、响应正文和请求方法，并亲自比较 `fmt.Fprint` 与 `fmt.Print` 的输出位置。
