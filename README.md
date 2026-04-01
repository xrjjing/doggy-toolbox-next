# doggy-toolbox-next

> 使用 **Go + Wails + Vue 3 + TypeScript** 重写 doggy-toolbox 的新项目白板。

## 项目定位

这个目录是一个**重写白板项目**，目标不是立刻实现全部功能，而是先把：

- 技术路线定清楚
- 目录结构定清楚
- 启动方式定清楚
- 分阶段迁移顺序定清楚
- 文档和协作边界定清楚

当前状态：**仅建立项目骨架与文档，不包含正式业务实现。**

## 为什么单独新建项目

- 不动当前 `doggy-toolbox` 的 Python + PyWebView 项目
- 给 Go / Vue 3 / TypeScript / Wails 学习留出安全空间
- 让后续重写过程可逐步推进，不和现有可用版本互相干扰
- 未来如果要推广或给别人使用，可以从一开始就按正式工程方式组织

## 目标技术栈

- **后端 / 桌面逻辑**：Go
- **桌面框架**：Wails
- **前端框架**：Vue 3
- **前端语言**：TypeScript
- **前端构建**：Vite
- **状态管理**：Pinia（计划）
- **路由**：Vue Router（按实际页面复杂度决定是否启用）

## 当前目录说明

```text
.
├── README.md
├── LICENSE
├── .gitignore
├── app/
├── build/
├── docs/
├── frontend/
├── internal/
└── scripts/
```

> 这些目录现在主要用于**占位和设计**，并非已经完成 Wails 官方脚手架初始化。

## 阅读顺序

1. `docs/00-START-HERE.md`
2. `docs/01-分阶段重写路线.md`
3. `docs/02-目录结构与配置说明.md`
4. `docs/03-启动与验证流程.md`
5. `docs/README.md`

## 当前分支规划

建议本白板项目在远端仓库使用独立分支维护，例如：

- `doggy-toolbox-next`

这样可以：

- 不影响现有 `main`
- 保留完整重写过程
- 后续需要时再决定是否单独拆仓库

## 近期最小目标

### 第 1 步：跑通技术栈

- 安装 Wails CLI
- 明确 Go / Node / npm / PATH
- 用官方模板生成一份 `vue-ts` 脚手架作为参考
- 确认 `wails dev` 能启动桌面窗口

### 第 2 步：搭基础壳

- 左侧导航
- 顶部标题区 / 工具区
- 一个设置页
- 一个最简单工具页（建议 Base64 或 JSON 格式化）

### 第 3 步：建立迁移节奏

- 先迁无状态小工具
- 再迁 CRUD 页面
- 最后迁 AI 聊天、节点转换等复杂模块

## License

本项目暂定沿用 MIT，详见 `LICENSE`。
