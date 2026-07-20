# 当前关：删除 Todo Boss

预计用时：约 20 分钟。今天只做这一关。

第六关已经通关。第七关完成内存 CRUD 的最后一个操作：删除 Todo，并用一次 Boss 复述串起 POST、GET、PUT、DELETE。

## 通关条件

- 运行 `go run .`
- 创建一条练习 Todo，再用 `DELETE /todos/{id}` 删除它
- 删除响应为 `204 No Content`，没有响应正文
- 随后查询相同 ID，看到 `404 Not Found`
- 能解释 slice 删除表达式，并复述四个 CRUD 方法各自操作什么

## 开始

打开 [第七关课程](lessons/0007-delete-todo.html)。

想先看整体方向，可以打开 [Go 后端闯关地图](roadmap.html)。

如果今天只有 2 分钟：只创建并删除一条 Todo，找到 `204 No Content`，也算保温成功；下一次仍从本关继续。

## 向老师通关

完成后回复两样东西：

1. 创建、删除以及随后查询得到 `404` 的关键响应；
2. 你怎么理解四个 CRUD 方法和 slice 删除表达式。

老师确认后才解锁下一关。
