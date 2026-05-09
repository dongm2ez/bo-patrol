# 博物馆智能巡检系统

基于 Go + Vue3 + UniApp 开发的博物馆智能巡检管理系统，实现巡检任务管理、扫码打卡、异常上报等核心功能。

## 📋 目录

- [项目简介](#项目简介)
- [技术栈](#技术栈)
- [功能特性](#功能特性)
- [项目结构](#项目结构)
- [快速开始](#快速开始)
- [部署说明](#部署说明)
- [API文档](#api文档)
- [贡献指南](#贡献指南)
- [许可证](#许可证)

## 🎯 项目简介

博物馆智能巡检系统是一套完整的巡检管理解决方案，旨在帮助博物馆实现巡检工作的数字化、智能化管理，提升巡检工作效率和质量。

系统包含三个端：
- **后端服务**：提供 RESTful API，处理业务逻辑
- **Web管理端**：供管理人员进行系统配置和数据查看
- **移动端**：供巡检人员使用，执行巡检任务

## 🛠️ 技术栈

### 后端
- **语言**: Go 1.21+
- **Web框架**: Gin 1.9+
- **ORM**: GORM 1.25+
- **数据库**: PostgreSQL 15+
- **认证**: JWT
- **缓存**: Redis (可选)
- **文件存储**: MinIO (可选)

### Web管理端
- **框架**: Vue 3.3+
- **TypeScript**: 5.0+
- **UI组件**: Element Plus 2.4+
- **路由**: Vue Router 4
- **状态管理**: Pinia
- **构建工具**: Vite 5

### 移动端
- **框架**: UniApp (Vue3)
- **支持平台**: App、H5、微信小程序

## ✨ 功能特性

### 核心功能
- [x] 用户管理 (管理员/巡检员)
- [x] 巡检路线管理
- [x] 巡检点位管理
- [x] 巡检任务创建与分配
- [x] 扫码打卡巡检
- [x] 异常问题上报
- [x] 工单流转处理
- [x] 数据统计分析

### 待实现功能
- [ ] 实时定位跟踪
- [ ] NFC打卡支持
- [ ] 巡检数据导出
- [ ] 消息通知推送
- [ ] 报表自动生成

## 📁 项目结构

```
bo-patrol/
├── api/                    # API定义
├── cmd/                    # 应用入口
│   ├── server/            # 主服务
│   └── migrate/           # 数据库迁移
├── configs/               # 配置文件
├── internal/              # 内部代码
│   ├── domain/           # 领域模型
│   ├── repository/       # 数据访问层
│   ├── service/          # 业务逻辑层
│   ├── handler/          # HTTP处理器
│   ├── middleware/       # 中间件
│   └── pkg/             # 内部工具包
├── pkg/                  # 公共包
├── web/                  # Vue3管理端
│   ├── src/
│   │   ├── views/       # 页面
│   │   ├── layout/      # 布局
│   │   ├── router/      # 路由
│   │   ├── stores/      # 状态管理
│   │   └── utils/       # 工具函数
│   └── package.json
├── mobile/               # UniApp移动端
│   ├── pages/           # 页面
│   ├── utils/           # 工具函数
│   ├── pages.json       # 页面配置
│   └── manifest.json    # 应用配置
├── deployments/          # 部署配置
├── .env.example         # 环境变量示例
├── Makefile             # 构建脚本
├── go.mod               # Go依赖
├── product-spec.md      # 产品文档
└── tech-spec.md         # 技术文档
```

## 🚀 快速开始

### 环境要求
- Go 1.21+
- Node.js 18+
- PostgreSQL 15+
- Redis 7+ (可选)
- MinIO (可选)

### 1. 克隆项目
```bash
git clone <repository-url>
cd bo-patrol
```

### 2. 配置环境变量
```bash
cp .env.example .env
# 编辑 .env 文件，配置数据库连接等信息
```

### 3. 启动后端服务
```bash
# 安装依赖
go mod download

# 运行服务
make run
# 或
go run cmd/server/main.go
```

服务将在 `http://localhost:8080` 启动

### 4. 启动Web管理端
```bash
cd web
npm install
npm run dev
```

访问 `http://localhost:3000`

### 5. 启动移动端
```bash
cd mobile
npm install
# H5开发模式
npm run dev:h5
# 微信小程序开发模式
npm run dev:mp-weixin
```

或使用 HBuilderX 打开 mobile 目录运行。

## 📦 部署说明

### Docker部署
```bash
# 构建后端镜像
docker build -f deployments/Dockerfile.server -t bo-patrol-server .

# 构建前端镜像
docker build -f deployments/Dockerfile.web -t bo-patrol-web .

# 使用docker-compose启动
docker-compose -f deployments/docker-compose.yml up -d
```

### 生产环境配置
1. 配置 PostgreSQL 数据库
2. 配置 JWT 密钥 (生产环境必须修改)
3. 配置 Redis 缓存 (可选)
4. 配置 MinIO 文件存储 (可选)
5. 配置 Nginx 反向代理

## 📚 API文档

### 认证接口
- `POST /api/v1/auth/register` - 用户注册
- `POST /api/v1/auth/login` - 用户登录

### 用户管理
- `GET /api/v1/users` - 获取用户列表
- `GET /api/v1/users/:id` - 获取用户详情
- `PUT /api/v1/users/:id` - 更新用户
- `DELETE /api/v1/users/:id` - 删除用户

### 巡检路线
- `GET /api/v1/routes` - 获取路线列表
- `POST /api/v1/routes` - 创建路线
- `GET /api/v1/routes/:id` - 获取路线详情
- `PUT /api/v1/routes/:id` - 更新路线
- `DELETE /api/v1/routes/:id` - 删除路线
- `GET /api/v1/routes/:id/points` - 获取点位列表
- `POST /api/v1/routes/:id/points` - 创建点位

### 巡检任务
- `GET /api/v1/tasks` - 获取任务列表
- `POST /api/v1/tasks` - 创建任务
- `GET /api/v1/tasks/:id` - 获取任务详情
- `PUT /api/v1/tasks/:id/start` - 开始任务
- `PUT /api/v1/tasks/:id/complete` - 完成任务

### 巡检记录
- `GET /api/v1/records` - 获取记录列表
- `POST /api/v1/records` - 创建打卡记录
- `GET /api/v1/records/:id` - 获取记录详情

### 工单管理
- `GET /api/v1/tickets` - 获取工单列表
- `POST /api/v1/tickets` - 创建工单
- `GET /api/v1/tickets/:id` - 获取工单详情
- `PUT /api/v1/tickets/:id/assign` - 派工单
- `PUT /api/v1/tickets/:id/process` - 处理工单
- `PUT /api/v1/tickets/:id/complete` - 完成工单

### 数据统计
- `GET /api/v1/stats/overview` - 概览统计
- `GET /api/v1/stats/completion` - 完成率统计
- `GET /api/v1/stats/abnormal` - 异常统计

## 📱 移动端页面

1. **登录页** - 用户登录
2. **首页** - 统计概览、快捷操作
3. **任务页** - 任务列表、扫码打卡
4. **记录页** - 巡检记录查看
5. **工单上报页** - 异常问题上报

## 💻 Web管理端页面

1. **登录页** - 管理员登录
2. **数据概览** - 仪表盘统计
3. **用户管理** - 用户CRUD
4. **路线管理** - 路线与点位管理
5. **任务管理** - 任务创建与管理
6. **巡检记录** - 打卡记录查看
7. **工单管理** - 工单处理与跟踪

## 🔧 开发说明

### 后端开发规范
- 遵循清晰分层架构 (Handler -> Service -> Repository)
- 使用 GORM 进行数据库操作
- 统一使用响应封装返回
- JWT 进行身份认证

### 前端开发规范
- 使用 TypeScript 编写
- 遵循 Vue3 Composition API 规范
- 使用 Element Plus 组件库
- 使用 Pinia 进行状态管理

## 🤝 贡献指南

1. Fork 本仓库
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 打开 Pull Request

## 📄 许可证

本项目采用 **MIT 开源许可证 + 商业许可证** 双许可模式。

### 开源 MIT 许可证
- ✅ 个人学习、研究使用
- ✅ 非商业用途部署
- ✅ 可自由修改、分发代码
- ✅ 需保留原始版权声明
- 详见 [LICENSE](LICENSE) 文件

### 商业许可证
适用于企业/机构商业部署、SaaS服务等场景，可提供专业技术支持与定制化开发服务。

商业合作请联系：dongm2ez@163.com（注明"博物馆巡检系统-商业合作"）

---

**说明**: 
- 开源版本完全免费，功能完整
- 选择双许可模式旨在平衡开源社区发展与可持续商业运营

## 📞 联系方式

如有问题或建议，欢迎提交 Issue 或联系项目维护者。

---

**注意**: 首次使用前请确保已正确配置数据库连接，并创建对应的数据库。
