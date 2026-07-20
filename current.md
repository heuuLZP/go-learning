# 当前关：用 POST 创建 Todo

预计用时：15–20 分钟。今天只做这一关。

第三关已经通关。第四关把数据方向反过来：上一关把 Go 数据写成响应 JSON，这一关从请求 JSON 中创建一个新的 Todo。

## 通关条件

- 运行 `go run .`
- 用 `POST /todos` 发送 JSON，看到 `201 Created` 和新 Todo
- 亲自修改请求中的 `title`，再创建一个 Todo
- 用 `GET /todos` 确认新 Todo 已追加到内存列表
- 能说出：`r.Body`、`Decode(&todo)`、`append` 和 `201 Created` 分别起什么作用

## 开始

打开 [第四关课程](lessons/0004-create-todo.html)。

想先看整体方向，可以打开 [Go 后端闯关地图](roadmap.html)。

如果今天只有 2 分钟：只发送一次课程中的 `POST /todos`，找到 `201 Created` 和响应 JSON，也算保温成功；下一次仍从本关继续。

## 向老师通关

完成后回复两样东西：

1. 修改标题后的 `POST` 响应，以及随后 `GET /todos` 的 JSON；
2. 你怎么理解 `r.Body`、`Decode(&todo)`、`append` 和 `201 Created`。

老师确认后才解锁下一关。
