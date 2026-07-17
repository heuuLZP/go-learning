# Go Learning

一个面向 Go 纯新手的个人学习仓库。目标不是先学完整语法，而是通过 10–20 分钟的微关卡，逐步打通 HTTP、REST、SQL、MySQL 和 Go 服务之间的真实后端链路。

## 现在开始

第一关：启动一个 Go HTTP 服务，访问 `GET /hello`，再亲自修改响应文字。

```bash
go run .
```

然后访问 <http://localhost:8080/hello>。

课程页面：[`lessons/0001-first-http-response.html`](lessons/0001-first-http-response.html)

当前任务与通关条件：[`current.md`](current.md)

## 学习方式

- 平时只打开 [`current.md`](current.md)，一次只关注当前关。
- 每关都要完成“运行 → 观察 → 亲自改一处 → 用一句话解释”。
- AI 负责样板代码，学习者负责理解、验证和小改动。
- 不要求每日打卡；中断后从当前关继续，不补欠下的课程。
- 按每周三次设计；首月完成 8、10 或 12 关都算有效通关。

## 仓库结构

- `current.md`：唯一的日常学习入口
- `lessons/`：每次只学习一个目标的关卡页面
- `reference/`：学过内容的精简速查资料
- `learning-records/`：学习过程中的阶段记录
- `docs/`：学习目标、课程设计、资料和教师笔记
- `main.go`：当前练习代码

更多说明：[`docs/learning-guide.md`](docs/learning-guide.md)。教师用课程设计见 [`docs/course-design.md`](docs/course-design.md)。

## 当前范围

第一个月完成内存 REST CRUD，并逐步接入 Docker MySQL。使用 Go 标准库和原生 SQL，不引入 Gin、ORM、Prisma、Redis 或前端页面；事务、索引、鉴权等能力在真实场景出现后再解锁。
