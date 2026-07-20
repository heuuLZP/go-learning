# 当前关：按 ID 查询 Todo

预计用时：15–20 分钟。今天只做这一关。

第四关已经通关。第五关继续使用 `GET`，但不再返回整个列表：从 URL 路径读取 ID，只返回匹配的 Todo，并区分 `200`、`400` 与 `404`。

## 通关条件

- 运行 `go run .`
- 用 `GET /todos/1` 看到 `200 OK` 和单个 Todo
- 亲自把路径 ID 改为另一个存在的 ID，看到不同 Todo
- 用 `GET /todos/999` 看到 `404 Not Found`
- 能说出：`PathValue`、`Atoi`、`range` 和 `404 Not Found` 分别起什么作用

## 开始

打开 [第五关课程](lessons/0005-get-todo-by-id.html)。

想先看整体方向，可以打开 [Go 后端闯关地图](roadmap.html)。

如果今天只有 2 分钟：只对比一次 `/todos/1` 和 `/todos/999` 的状态码，也算保温成功；下一次仍从本关继续。

## 向老师通关

完成后回复两样东西：

1. 查询一个存在 ID 和一个不存在 ID 的响应；
2. 你怎么理解 `PathValue`、`Atoi`、`range` 和 `404 Not Found`。

老师确认后才解锁下一关。
