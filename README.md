# 🎬 go-jocy

> 基于 Gin 的高性能二次元视频聚合后端服务

> 前端开源地址: [https://github.com/nuanxinqing123/jocy-web-refactoring](https://github.com/nuanxinqing123/jocy-web-refactoring)

---

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.23%2B-blue" />
  <img src="https://img.shields.io/badge/Gin-1.10.0-brightgreen" />
  <img src="https://img.shields.io/badge/License-MIT-yellow" />
  <img src="https://img.shields.io/badge/Docker-Support-blue" />
</p>

## 项目简介

**go-jocy** 是一个基于 Go + Gin 框架开发的高性能二次元视频聚合后端服务，支持多源视频聚合、弹幕、评论、用户系统、收藏、历史记录等功能，适合二次元内容聚合平台、学习和二次开发。

---

## 主要特性

- 🚀 高性能：基于 Gin 框架，支持高并发
- 🧩 多源聚合：支持多站点视频内容聚合
- 🗂️ 丰富接口：涵盖用户、视频、弹幕、评论、收藏等
- 🔒 支持鉴权与限流中间件
- 🛠️ 灵活配置，支持热更新
- 🐳 一键 Docker 部署
- 📦 完善的目录结构，易于维护和扩展

---

## 目录结构

```bash
├── build.sh / build.bat         # 构建脚本
├── config/                      # 配置文件及加载逻辑
│   ├── autoload/                # 配置结构体定义
│   └── config.go                # 配置入口
├── const/                       # 常量定义
├── docker-compose.yml           # Docker Compose 部署
├── docker-entrypoint.sh         # Docker 容器启动脚本
├── Dockerfile                   # Docker 构建文件
├── initialize/                  # 初始化相关（路由、日志、配置）
├── internal/
│   ├── controllers/             # 业务控制器（API实现）
│   ├── middleware/              # Gin中间件
│   ├── model/                   # 数据结构定义
│   └── router/                  # 路由注册
├── logs/                        # 日志目录
├── main.go                      # 程序入口
├── test/                        # 测试脚本
├── utils/                       # 工具函数
└── README.md                    # 项目说明
```

---

## 主要依赖

| 依赖包                                              | 说明       |
|--------------------------------------------------|----------|
| [Gin](https://github.com/gin-gonic/gin)          | Web 框架   |
| [Viper](https://github.com/spf13/viper)          | 配置管理     |
| [Zap](https://github.com/uber-go/zap)            | 日志系统     |
| [Resty](https://github.com/go-resty/resty)       | HTTP 客户端 |
| [Gopher-Lua](https://github.com/yuin/gopher-lua) | Lua 支持   |
| 见 `go.mod` 获取完整依赖列表                              |

---

## 快速开始

### 本地运行

1. **克隆项目**
   ```bash
   git clone https://github.com/你的仓库/go-jocy.git
   cd go-jocy
   ```
2. **配置环境**
   - 复制或编辑 `config/config.yaml`，参考 `config/autoload/app.go` 字段说明
3. **安装依赖**
   ```bash
   go mod tidy
   ```
4. **运行项目**
   ```bash
   go run main.go
   ```

### Docker 一键部署

1. **构建镜像并运行**
   ```bash
   docker-compose up -d
   ```
2. **首次启动会自动生成配置文件**，可在 `config/` 目录下自定义

---

## 配置说明

`config/autoload/app.go` 字段如下：

| 字段           | 类型       | 说明                |
|--------------|----------|-------------------|
| mode         | string   | 运行模式（dev/release） |
| address      | string   | 监听地址              |
| port         | int      | 监听端口              |
| baseURL      | []string | 视频源基础URL          |
| appVersion   | string   | 应用版本号             |
| play_aes_key | string   | 播放加密Key(备用)       |
| play_aes_iv  | string   | 播放加密IV(备用)        |

配置文件路径: config/config.yaml

```yaml
app:
  # 运行模式【debug、release】
  mode:        "release"
  # 运行地址【不懂请写: 0.0.0.0】
  address:     "127.0.0.1"
  # 运行端口
  port:        8090
  # BaseURL
  baseURL:
    - "http://wudi.yaxuan.top"
  # App版本
  appVersion:  "1.5.7.9"
  # 播放地址密钥
  play_aes_key: "wcyjmnnnawozmydn"
  play_aes_iv:  "wcivwyjmlnzbhlmq"
```

---

## API 概览

> 详细接口请参考 `internal/controllers/app.go` 和 `internal/router/app.go`

- **用户相关**
  - `GET /users/avatar` 随机头像
  - `POST /users/captcha` 获取验证码
  - `POST /users/smscode` 发送验证码
  - `POST /users/register` 用户注册
  - `POST /users/login` 用户登录
  - `POST /users/logout` 用户登出
  - `GET /users/info` 用户信息
- **视频相关**
  - `GET /video/list` 视频列表
  - `GET /video/detail` 视频详情
  - `GET /video/play` 视频播放
  - `GET /video/play/params` 播放参数
  - `GET /video/search` 视频搜索
  - `GET /video/key` 预搜索
  - `POST /play/resources` 获取播放资源
- **弹幕/评论/消息**
  - `GET /danmu` 弹幕
  - `GET /vod_comment/getlist` 评论列表
  - `GET /vod_comment/getsublist` 子评论
  - `GET /vod_comment/gethitstop` 热评
  - `GET /messagebox` 消息盒子
  - `GET /messagebox/:type` 指定类型消息
- **收藏/历史**
  - `GET /collect` 我的收藏
  - `POST /collect` 添加收藏
  - `DELETE /collect` 删除收藏
  - `GET /history` 观看历史
  - `POST /history` 上传历史

---

## 贡献指南

欢迎 Issue、PR 和 Star！

1. Fork 本仓库
2. 新建分支 (`git checkout -b feature/xxx`)
3. 提交更改 (`git commit -am 'feat: xxx'`)
4. 推送分支 (`git push origin feature/xxx`)
5. 新建 Pull Request
