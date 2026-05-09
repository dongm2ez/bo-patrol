<div align="center">
  <a href="https://github.com/dongm2ez/bo-patrol">
    <img alt="BoPatrol Logo" width="120" src="https://trae-api-cn.mchost.guru/api/v1/text_to_image?prompt=ancient chinese museum security shield logo with traditional patterns blue gold style&image_size=square">
  </a>
</div>

<h1 align="center">🛡️ 博巡 - 博物馆智能巡检系统</h1>

<p align="center">
  博物馆智能巡检，开源守护文化
</p>

<p align="center">
  <a href="https://github.com/dongm2ez/bo-patrol/stargazers">
    <img alt="GitHub stars" src="https://img.shields.io/github/stars/dongm2ez/bo-patrol?style=flat-square">
  </a>
  <a href="https://github.com/dongm2ez/bo-patrol/network/members">
    <img alt="GitHub forks" src="https://img.shields.io/github/forks/dongm2ez/bo-patrol?style=flat-square">
  </a>
  <a href="https://github.com/dongm2ez/bo-patrol/issues">
    <img alt="GitHub issues" src="https://img.shields.io/github/issues/dongm2ez/bo-patrol?style=flat-square">
  </a>
  <a href="https://github.com/dongm2ez/bo-patrol/blob/master/LICENSE">
    <img alt="GitHub license" src="https://img.shields.io/github/license/dongm2ez/bo-patrol?style=flat-square">
  </a>
</p>

---

## ✨ 项目简介

**博巡 (BoPatrol)** 是一套专为博物馆打造的开源智能巡检管理系统。采用 **Go + Vue3 + TypeScript** 技术栈，围绕博物馆特色业务场景构建，帮助博物馆实现巡检工作的数字化、智能化管理，守护文化遗产安全。

> 🏛️ **博物馆的需求，我们更懂**

### 🎯 核心理念

- **专业** - 深度贴合博物馆巡检业务场景
- **高效** - 扫码巡检，快速上报，智能派单
- **开源** - 代码完全开源，社区共同维护
- **易用** - 界面简洁，操作流畅，上手即会

## 🌟 功能特性

### 📊 核心功能模块

| 模块 | 功能说明 | 状态 |
|-----|---------|------|
| 🏢 **科室管理** | 灵活的组织架构，支持水科、强电科、电梯科等专业科室 | ✅ |
| 📦 **设备台账** | 全生命周期管理，设备归属、维保记录、部署位置 | ✅ |
| 🗺️ **空间地图** | 建筑-楼层-区域三级空间结构，可视化地图编辑 | ✅ |
| 🛤️ **路线管理** | 自定义巡检路线，博物馆展区专属配置 | ✅ |
| 🎯 **点位管理** | 巡检点位二维码，关联展品与安防设备 | ✅ |
| 📋 **任务管理** | 定时/临时/专项任务，智能分配到对应科室 | ✅ |
| 📱 **移动巡检** | UniApp 移动端，扫码打卡，GPS 定位（可选） | ✅ |
| 🚨 **异常上报** | 照片+文字描述，自动生成工单流转 | ✅ |
| 📈 **统计分析** | 巡检完成率、异常趋势、人员绩效统计 | ✅ |
| 🏛️ **展品管理** | 文物展品专属巡检，特殊保护需求 | ✅ |
| 🌡️ **环境监测** | 展厅温湿度、光照度等环境数据记录 | ✅ |

### 🎨 博物馆特色功能

- ✅ 展品完整性专项巡检
- ✅ 展厅温湿度环境监测
- ✅ 安防设备（摄像头、报警器）巡检
- ✅ 按科室专业分工的巡检内容定制
- ✅ 文物库房重点区域巡检
- ✅ 闭馆后夜间安全巡检

## 🛠️ 技术架构

### 后端技术栈

| 技术 | 版本 | 用途 |
|-----|-----|------|
| **Go** | 1.21+ | 核心编程语言 |
| **Gin** | 1.9+ | Web 框架 |
| **GORM** | 1.25+ | ORM 框架 |
| **PostgreSQL** | 15+ | 关系型数据库 |
| **JWT** | - | 身份认证 |
| **Redis** | 7+ | 缓存（可选） |
| **MinIO** | - | 文件存储（可选） |

### 前端技术栈

| 技术 | 版本 | 用途 |
|-----|-----|------|
| **Vue 3** | 3.3+ | 前端框架 |
| **TypeScript** | 5.0+ | 类型系统 |
| **Element Plus** | 2.4+ | UI 组件库 |
| **Vite** | 5+ | 构建工具 |
| **Pinia** | - | 状态管理 |
| **Vue Router** | 4 | 路由管理 |

### 移动端技术栈

| 技术 | 用途 |
|-----|------|
| **UniApp (Vue3)** | 跨端框架 |
| **微信小程序** | 小程序支持 |
| **H5** | 移动端网页 |
| **App** | iOS / Android 应用 |

## 📁 项目结构

```
bo-patrol/
├── api/                    # API 定义
├── cmd/                    # 应用入口
│   ├── server/            # 后端主服务
│   └── migrate/           # 数据库迁移工具
├── configs/               # 配置文件
├── internal/              # 核心业务代码
│   ├── domain/           # 领域模型定义
│   ├── repository/       # 数据访问层
│   ├── service/          # 业务逻辑层
│   ├── handler/          # HTTP 接口层
│   └── middleware/       # 中间件
├── pkg/                   # 公共工具包
├── web/                   # Vue3 管理后台
│   ├── src/
│   │   ├── views/        # 页面组件
│   │   ├── layout/       # 布局组件
│   │   ├── api/          # API 调用
│   │   └── utils/        # 工具函数
├── mobile/                # UniApp 移动端
├── deployments/           # Docker 部署配置
├── docs/                  # 项目文档
│   ├── product-spec.md   # 产品需求文档
│   └── tech-spec.md      # 技术架构文档
├── .env.example           # 环境变量示例
├── Makefile               # 构建脚本
└── go.mod                 # Go 依赖管理
```

## 🚀 快速开始

### 环境准备

- Go 1.21+
- Node.js 18+
- PostgreSQL 15+

### 1. 克隆项目

```bash
git clone https://github.com/dongm2ez/bo-patrol.git
cd bo-patrol
```

### 2. 配置环境变量

```bash
cp .env.example .env
# 编辑 .env 文件，配置数据库连接等信息
```

### 3. 启动后端服务

```bash
# 安装 Go 依赖
go mod download

# 执行数据库迁移（初始化表结构）
go run cmd/migrate/main.go

# 启动服务
go run cmd/server/main.go
# 或使用 Makefile
make run
```

后端服务将在 `http://localhost:8080` 启动

### 4. 启动 Web 管理端

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

# H5 开发模式
npm run dev:h5

# 微信小程序开发模式
npm run dev:mp-weixin
```

> 也可以使用 HBuilderX 打开 mobile 目录进行开发和发行

## 🔐 测试账号

| 角色 | 用户名 | 密码 | 说明 |
|-----|-------|-----|------|
| 管理员 | admin | admin123 | 系统最高权限 |
| 巡检员 | inspector | 123456 | 执行巡检任务 |

## 🎬 功能预览

### 🖥️ Web 管理端

| 页面 | 功能 | 预览 |
|-----|------|------|
| 登录页 | 品牌登录入口 | ✅ |
| 数据概览 | 巡检统计仪表盘 | ✅ |
| 科室管理 | 组织架构与权限 | ✅ |
| 设备台账 | 全生命周期管理 | ✅ |
| 空间管理 | 博物馆空间结构 | ✅ |
| 地图编辑 | 可视化点位布局 | ✅ |
| 路线管理 | 巡检路线配置 | ✅ |
| 任务管理 | 任务创建与分配 | ✅ |
| 巡检记录 | 打卡记录查看 | ✅ |
| 工单管理 | 异常工单处理 | ✅ |

### 📱 移动端

| 页面 | 功能 |
|-----|------|
| 登录页 | 移动端认证 |
| 首页 | 待办任务与快捷操作 |
| 任务 | 扫码执行巡检 |
| 记录 | 个人巡检记录 |
| 上报 | 异常问题上报 |

## 📚 文档资源

- [产品需求文档](docs/product-spec.md) - 详细的业务功能说明
- [技术架构文档](docs/tech-spec.md) - 技术实现与架构设计
- [API 接口文档](docs/api.md) - RESTful API 说明
- [部署指南](docs/deploy.md) - 生产环境部署说明
- [开发指南](docs/dev-guide.md) - 开发者入门指南

## 🤝 贡献指南

我们欢迎任何形式的贡献！无论你是：

- 🐛 提交 Bug 报告
- 💡 提出新功能建议
- 📝 改进文档
- 💻 提交代码 PR

### 参与开发

1. Fork 本仓库
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 打开 Pull Request

### 代码规范

- 后端：遵循 Go 官方编码规范
- 前端：使用 TypeScript，遵循 Vue3 Composition API 规范
- 提交信息：使用 [Conventional Commits](https://www.conventionalcommits.org/) 规范

## 🌟 星标历史

[![Star History Chart](https://api.star-history.com/svg?repos=dongm2ez/bo-patrol&type=Date)](https://star-history.com/#dongm2ez/bo-patrol&Date)

## 📄 许可证

本项目采用 **MIT 开源许可证 + 商业许可证** 双许可模式。

### 开源 MIT 许可证

✅ 个人学习、研究使用
✅ 非商业用途部署
✅ 可自由修改、分发代码
✅ 需保留原始版权声明

详见 [LICENSE](LICENSE) 文件

### 商业许可证

适用于企业/机构商业部署、SaaS服务等场景。可提供：

- 🎯 专业技术支持
- 📦 定制化开发服务
- 📱 移动端品牌定制
- 🚀 部署运维指导
- 📚 管理员培训

**商业合作请联系**: dongm2ez@163.com（注明"博巡系统-商业合作"）

## 💬 社区交流

- GitHub Issues: [提交问题](https://github.com/dongm2ez/bo-patrol/issues)
- Email: dongm2ez@163.com
- 微信: (待定)

## 👥 贡献者

感谢所有为博巡项目做出贡献的人！

<a href="https://github.com/dongm2ez/bo-patrol/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=dongm2ez/bo-patrol" />
</a>

## 🔔 更新日志

### v1.0.0 (2024-xx-xx)

- ✨ 初版发布
- ✨ 完整的巡检核心功能
- ✨ 博物馆科室与设备管理
- ✨ 空间地图与可视化编辑
- ✨ Web 管理端与移动端

## ⭐ 支持项目

如果这个项目对你有帮助，欢迎：

- 🌟 给项目点个 Star
- 📢 分享给更多需要的人
- 💬 提出建议帮助我们做得更好
- 🤝 参与代码贡献

---

<div align="center">
  <strong>🏛️ 博巡 - 博物馆智能巡检，开源守护文化</strong>
  <br>
  Made with ❤️ by the BoPatrol Team
</div>
