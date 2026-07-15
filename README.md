# Go Learning

一个面向 Go 纯新手的个人学习仓库。目标不是先学完整语法，而是通过 10–20 分钟的微关卡，逐步获得阅读、运行、修改和排错 Go 后端代码的能力。

## 学习方式

- 平时只打开 [`CURRENT.md`](CURRENT.md)，一次只关注当前关。
- 每关都要完成“运行 → 观察 → 亲自改一处 → 用一句话解释”。
- AI 负责样板代码，学习者负责理解、验证和小改动。
- 不要求每日打卡；中断后从当前关继续，不补欠下的课程。

## 当前关卡

第一关：启动一个 Go HTTP 服务，访问 `GET /hello`，再亲自修改响应文字。

```bash
go run .
```

然后访问 <http://localhost:8080/hello>。

课程页面：[`lessons/0001-first-http-response.html`](lessons/0001-first-http-response.html)

## 仓库结构

- `CURRENT.md`：唯一的日常学习入口
- `lessons/`：每次只学习一个目标的关卡页面
- `reference/`：学过内容的精简速查资料
- `MISSION.md`：学习目标与边界
- `COURSE-DESIGN.md`：教师使用的首月关卡池，学习时无需提前阅读
- `main.go`：当前练习代码

## 当前范围

第一个月优先练习 HTTP、JSON、struct、slice、输入处理、状态码、基础测试和排错。数据库、框架、gRPC 与并发等内容延后到真正需要时再学习。
