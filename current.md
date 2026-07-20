# 当前关：让 Todo 变成 JSON

预计用时：15–20 分钟。今天只做这一关。

第二关已经通关。第三关继续使用 `GET`、`200 OK` 和 `w`，只增加一层：用 Go 的 `struct` 表示 Todo，再把它写成 JSON 响应。

## 通关条件

- 运行 `go run .`
- 用 `curl -i http://localhost:8080/todos` 看到 `200 OK`、`Content-Type: application/json` 和一个 Todo
- 在 `todos` 中亲自增加第二个 Todo，重启后看到 JSON 数组中有两项
- 能说出：Go 的 `Todo` 如何对应 JSON，以及 `json.NewEncoder(w)` 中的 `w` 起什么作用

## 开始

打开 [第三关课程](lessons/0003-todo-json.html)。

想先看整体方向，可以打开 [Go 后端闯关地图](roadmap.html)。

如果今天只有 2 分钟：只运行一次 `curl -i http://localhost:8080/todos`，找到状态码、`Content-Type` 和 JSON 正文，也算保温成功；下一次仍从本关继续。

## 向老师通关

完成后回复两样东西：

1. 增加第二个 Todo 前后的 JSON 正文；
2. 你怎么理解 `Todo`、JSON 字段和 `json.NewEncoder(w)` 的关系。

老师确认后才解锁下一关。
