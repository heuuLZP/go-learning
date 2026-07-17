# 当前关：看清一次 HTTP 响应

预计用时：15–20 分钟。今天只做这一关。

第一关已经通关。第二关沿着你对 `Fprint` 和参数 `w` 的疑问，观察状态码、响应正文和输出位置。

## 通关条件

- 运行 `go run .`
- 用 `curl -i http://localhost:8080/ping` 看到 `200 OK` 和 `pong`
- 用 `POST` 请求 `/ping`，看到 `405 Method Not Allowed`
- 完成一次 `fmt.Fprint` 与 `fmt.Print` 的对比实验，并把代码恢复
- 能说出：`w` 在 handler 中起什么作用

## 开始

打开 [第二关课程](lessons/0002-http-response.html)。

想先看整体方向，可以打开 [Go 后端闯关地图](roadmap.html)。

如果今天只有 2 分钟：只运行一次 `curl -i http://localhost:8080/ping`，找到状态码和正文，也算保温成功；下一次仍从本关继续。

## 向老师通关

完成后回复两样东西：

1. `GET /ping` 和 `POST /ping` 分别看到了什么状态码；
2. 你怎么理解 `w`，以及 `fmt.Fprint(w, ...)` 和 `fmt.Print(...)` 的区别。

老师确认后才解锁下一关。
